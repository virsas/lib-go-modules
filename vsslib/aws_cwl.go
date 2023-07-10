package vsslib

import (
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type CWLHandler interface {
	Write(msg string) error
	Get(limit int64) (*cloudwatchlogs.GetLogEventsOutput, error)
}

type cwl struct {
	session *cloudwatchlogs.CloudWatchLogs
	mutex   *sync.Mutex
	token   *string
	group   string
	stream  string
}

func NewCloudwatchSess(sess *session.Session) (CWLHandler, error) {
	var err error

	var cwGroup string = ""
	cwGroupValue, cwGroupPresent := os.LookupEnv("AWS_CW_GROUP")
	if cwGroupPresent {
		cwGroup = cwGroupValue
	} else {
		panic("Missing ENV Variable AWS_CW_GROUP")
	}

	var cwStream string = ""
	cwStreamValue, cwStreamPresent := os.LookupEnv("AWS_CW_STREAM")
	if cwStreamPresent {
		cwStream = cwStreamValue
	} else {
		panic("Missing ENV Variable AWS_CW_STREAM")
	}

	cwsess := cloudwatchlogs.New(sess)
	cwh := &cwl{session: cwsess, mutex: &sync.Mutex{}, token: nil, group: cwGroup, stream: cwStream}

	err = createCWLGroup(cwsess, cwGroup)
	if err != nil {
		return nil, err
	}

	err = createCWLStream(cwsess, cwGroup, cwStream)
	if err != nil {
		return nil, err
	}

	return cwh, nil
}

func createCWLGroup(sess *cloudwatchlogs.CloudWatchLogs, group string) error {
	_, err := sess.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &group,
	})
	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok {
			if aerr.Code() == "ResourceAlreadyExistsException" {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func createCWLStream(sess *cloudwatchlogs.CloudWatchLogs, group string, stream string) error {
	_, err := sess.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &group,
		LogStreamName: &stream,
	})
	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok {
			if aerr.Code() == "ResourceAlreadyExistsException" {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func getNextToken(sess *cloudwatchlogs.CloudWatchLogs, group string, stream string) (*string, error) {
	resp, err := sess.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &group,
		LogStreamNamePrefix: &stream,
	})
	if err != nil {
		return nil, err
	}

	return resp.LogStreams[0].UploadSequenceToken, nil
}

func (cwh *cwl) Write(msg string) error {
	var err error
	var logEvent []*cloudwatchlogs.InputLogEvent

	if cwh.token == nil {
		resp, err := getNextToken(cwh.session, cwh.group, cwh.stream)
		if err != nil {
			return err
		}
		cwh.token = resp
	}

	logEvent = append(logEvent, &cloudwatchlogs.InputLogEvent{
		Message:   &msg,
		Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
	})

	cwh.mutex.Lock()
	resp, err := cwh.session.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  &cwh.group,
		LogStreamName: &cwh.stream,
		LogEvents:     logEvent,
		SequenceToken: cwh.token,
	})

	if err != nil {
		var err2 error
		cwh.token, err2 = getNextToken(cwh.session, cwh.group, cwh.stream)
		cwh.mutex.Unlock()
		if err2 != nil {
			return err2
		}
		return err
	}

	cwh.token = resp.NextSequenceToken
	cwh.mutex.Unlock()
	return nil
}

func (cwh *cwl) Get(limit int64) (*cloudwatchlogs.GetLogEventsOutput, error) {
	resp, err := cwh.session.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         &limit,
		LogGroupName:  &cwh.group,
		LogStreamName: &cwh.stream,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
