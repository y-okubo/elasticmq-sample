package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var svc *sqs.SQS

func main() {

	const (
		AWS_REGION      = "us-west-2"
		MAIN_QUEUE_NAME = "main-queue"
		DEAD_QUEUE_NAME = "dead-queue"
	)

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(AWS_REGION),
		Endpoint: aws.String("http://127.0.0.1:9324"),
	},
	)

	// Create a SQS service client.
	svc = sqs.New(sess)

	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String("http://localhost:9324/queue/MyGroovyQueue"),
		// 一度に取得する最大メッセージ数。最大でも10まで。
		MaxNumberOfMessages: aws.Int64(1),
		// これでキューが空の場合はロングポーリング(20秒間繋ぎっぱなし)になる。
		WaitTimeSeconds: aws.Int64(20),
	}
	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("messages count: %d\n", len(resp.Messages))
	fmt.Println(*resp.Messages[0].Body)
	// if len(resp.Messages) != 0 {
	// 	DeleteMessage(resp.Messages[0])
	// }
}

// メッセージを削除する。
func DeleteMessage(msg *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String("http://localhost:9324/queue/MyGroovyQueue"),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}
	_, err := svc.DeleteMessage(params)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
