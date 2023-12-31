package vssvar

import (
	"os"

	"github.com/virsas/lib-go-modules/vsslib"
)

type AWSCreds struct {
	ID     string
	Key    string
	Region string
}

func OPAWSCreds(op vsslib.OpHandler, opAWSItem string) (AWSCreds, error) {
	var err error
	var awscreds AWSCreds = AWSCreds{}

	awscreds.ID, err = op.Get(opAWSItem, "username")
	if err != nil {
		return awscreds, err
	}

	awscreds.Key, err = op.Get(opAWSItem, "password")
	if err != nil {
		return awscreds, err
	}

	var awsRegion string = "eu-west-1"
	awsRegionValue, awsRegionPresent := os.LookupEnv("AWS_REGION")
	if awsRegionPresent {
		awsRegion = awsRegionValue
	}
	awscreds.Region = awsRegion

	return awscreds, nil
}
