package vsslib

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession(op OpHandler) (*session.Session, error) {
	var err error
	var sess *session.Session

	var opAWSUser string = ""
	opAWSUserValue, opAWSUserPresent := os.LookupEnv("OP_AWS_USER")
	if opAWSUserPresent {
		opAWSUser = opAWSUserValue
	} else {
		panic("Missing ENV Variable OP_AWS_USER")
	}

	awsID, err := op.Get(opAWSUser, "username")
	if err != nil {
		return nil, err
	}

	awsKey, err := op.Get(opAWSUser, "password")
	if err != nil {
		return nil, err
	}

	var awsRegion string = "eu-west-1"
	awsRegionValue, awsRegionPresent := os.LookupEnv("AWS_REGION")
	if awsRegionPresent {
		awsRegion = awsRegionValue
	}

	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsID, awsKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}
