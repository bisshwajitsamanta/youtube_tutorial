package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ListFiles() (newList []string) {
	// Listing out all the Zip files present
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Printf("Unable to read data from the Directory:%w", err.Error())
	}
	newList = []string{}
	for _, f := range files {
		if strings.Contains(f.Name(), "zip") == true {
			newList = append(newList, f.Name())
		}
	}
	return newList
}

func UploadZip(sess *session.Session) {
	newList := ListFiles()
	for _, v := range newList {
		file, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
		}
		uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
			u.PartSize = 10 * 1024 * 1024 // Default File Size is 5MB
			u.Concurrency = 2
			u.MaxUploadParts = 10
		})
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("myawesomes3bucket-7sq3e2v97bo2"),
			ACL:    aws.String("public-read"),
			Body:   file,
			Key:    aws.String(v),
		})
		// In case it fails to upload
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("File Uploaded :", result.Location)
	}

}

func main() {
	config := aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("/Users/bisshwajitsamanta/.aws/credentials", "default"),
	}
	sess, err := session.NewSession(&config)
	if err != nil {
		fmt.Println(err)
	}
	UploadZip(sess)

}
