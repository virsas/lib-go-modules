package vsslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession(awsID string, awsKey string, awsRegion string) (*session.Session, error) {
	var err error
	var sess *session.Session

	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsID, awsKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}
