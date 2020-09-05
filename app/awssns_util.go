package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// SnsHandler is in charge that send message through AWS SNS.
type SnsHandler struct {
	cfg         aws.Config
	snsTopicArn string
}

var snsTopicArn string = getEnv("SNS_TOPIC_ARN", "")

func (s *SnsHandler) publishAlert(message string) {
	svc := sns.New(s.cfg)
	publishInput := sns.PublishInput{
		Message:  &message,
		TopicArn: &s.snsTopicArn,
	}

	req := svc.PublishRequest(&publishInput)
	resp, err := req.Send(context.TODO())
	if err == nil {
		fmt.Println("following message was sent successfully ")
		fmt.Println(message)
		fmt.Println(resp)
	}
}

func listSNSTopics(cfg aws.Config) {
	svc := sns.New(cfg)
	s := ""
	listTopicsInput := sns.ListTopicsInput{NextToken: &s}
	req := svc.ListTopicsRequest(&listTopicsInput)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	if err == nil {
		fmt.Println(resp)
	}

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func (s *SnsHandler) init() error {
	cfg, err := external.LoadDefaultAWSConfig()

	if err != nil {
		exitErrorf("failed to load config, %v", err)
		return err
	}
	s.cfg = cfg
	// TODO validate ARN
	if snsTopicArn != "" {
		s.snsTopicArn = snsTopicArn
		fmt.Println("SNS_TOPIC_ARN is ", s.snsTopicArn)
	} else {
		exitErrorf("SNS_TOPIC_ARN is not set", err)
	}
	fmt.Println("completed to load aws config")
	return nil
}

// Init initialise SnsHandler
func Init() (SnsHandler, error) {
	snsHandler := SnsHandler{}
	err := snsHandler.init()
	if err != nil {
		fmt.Println(err)
		return SnsHandler{}, err
	}
	return snsHandler, nil
}
