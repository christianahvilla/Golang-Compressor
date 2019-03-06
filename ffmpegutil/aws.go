package ffmpegutil

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//Const to set session for aws server
const (
	AwsBucket = "vis3dev"
	AwsRegion = "us-east-1"
)

//AwsSet is a function where server session set
func AwsSet() (*session.Session, error) {
	svc, err := session.NewSession(&aws.Config{Region: aws.String(AwsRegion)})
	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		WriteLog(Info, "Session opened for response: "+IDResponse)
	}
	return svc, err
}

//AddFileToS3 is a function to save files in Amazon
func AddFileToS3(svc *session.Session) error {
	// Open the file for use
	file, err := os.Open("videos/" + NameFile + ".mp4")

	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		// Config settings: this is where you choose the bucket, filename, content-type etc.
		// of the file you're uploading.

		fileInfo, _ := file.Stat()
		var size = fileInfo.Size()
		buffer := make([]byte, size)
		file.Read(buffer)
		fileBytes := bytes.NewReader(buffer)
		fileType := http.DetectContentType(buffer)

		acl := "public-read"

		// Config settings: this is where you choose the bucket, filename, content-type etc.
		// of the file you're uploading.
		_, err = s3.New(svc).PutObject(&s3.PutObjectInput{
			Bucket:        aws.String(AwsBucket),
			Key:           aws.String(NameFile + ".mp4"),
			Body:          fileBytes,
			ContentLength: aws.Int64(size),
			ContentType:   aws.String(fileType),
			ACL:           aws.String(acl),
		})

		if err != nil {
			WriteLog(Error, err.Error())
		} else {
			WriteLog(Info, "File uploaded for response: "+IDResponse)
		}

		file.Close()
	}
	return err
}

//GetFileLink is a function to get file URL
func GetFileLink() {
	URLVideo = "https://vis3dev.s3.amazonaws.com/" + NameFile + ".mp4"
	WriteLog(Info, "URL "+URLVideo+" for response: "+IDResponse)
}
