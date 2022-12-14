package common

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateFile(content *File, sess *session.Session) (*File, error) {
	
	output, err := uploadFile(content, sess)
	
	if err != nil {
		fmt.Printf("error generating file: %v", err)
		return nil, err
	}

	content.Path = output.Location
	fmt.Println(content.Path)
	return content, nil
}

func uploadFile(content *File, sess *session.Session) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(sess)
	
	input := &s3manager.UploadInput{
		Bucket:      aws.String(content.Bucket),
		Key:         aws.String(content.FileName),
		Body:        bytes.NewReader([]byte(content.Content)),
		ContentType: aws.String("text/plain"),
	}
	fmt.Println("------")
	fmt.Println(content.Bucket)
	fmt.Println(content.FileName)
	output, err := uploader.UploadWithContext(context.Background(), input)
	
	fmt.Printf("*-*-*-*-")
	if err != nil {
		fmt.Printf("Error uploading file: %v", err)
		return nil, err
	}

	return output, nil
}
