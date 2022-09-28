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
	fmt.Printf("Error CreateFile 1")
	output, err := uploadFile(content, sess)
	fmt.Printf("Error CreateFile 2")
	if err != nil {
		fmt.Printf("error generating file: %v", err)
		return nil, err
	}

	content.Path = output.Location
	return content, nil
}

func uploadFile(content *File, sess *session.Session) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(sess)
	fmt.Printf("Error uploadFile 1")
	input := &s3manager.UploadInput{
		Bucket:      aws.String(content.Bucket),
		Key:         aws.String(content.FileName),
		Body:        bytes.NewReader([]byte(content.Content)),
		ContentType: aws.String("text/plain"),
	}
	fmt.Printf("Error uploadFile 2")
	output, err := uploader.UploadWithContext(context.Background(), input)
	
	fmt.Printf("Error uploadFile 3")
	if err != nil {
		fmt.Printf("Error uploading file: %v", err)
		fmt.Printf(err)
		fmt.Printf(output)
		return nil, err
	}

	return output, nil
}
