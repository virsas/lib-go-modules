package vsslib

import (
	"os"

	"github.com/virsas/lib-go-modules/vsslib"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSess(op vsslib.OpHandler) (*session.Session, error) {
	var err error
	var sess *session.Session

	awsID, err := op.GetItem("acc_aws_iam_user", "username")
	if err != nil {
		return nil, err
	}

	awsKey, err := op.GetItem("acc_aws_iam_user", "password")
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
