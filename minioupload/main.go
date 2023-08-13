package main

import (
	"flag"
	"github.com/minio/minio-go"
	"log"
)

func FileUploader(endpoint, accessKeyID, secretAccessKey, filePath, objectName, bucketName string, useSSL bool) {

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	//bucketName := "test"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	//objectName := "111.txt"
	//filePath := "C:\\Users\\volin\\GolandProjects\\go-learn\\minioupload\\1111.txt"
	contentType := "application/data"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

func main() {
	//objectName = "12311.txt"
	//filePath = "C:\\Users\\volin\\GolandProjects\\go-learn\\minioupload\\1111.txt"
	//endpoint = "192.168.7.245:9000"
	//accessKeyID = "AKIDxxxxxxxxx"
	//secretAccessKey = "AKEYyyyyyyyyy23123123123"
	//useSSL = false
	//bucketName = "test"

	var endpoint, accessKeyID, secretAccessKey, filePath, objectName, bucketName string
	var useSSL bool

	flag.StringVar(&endpoint, "ep", "localhost:9000", "minio地址")
	flag.StringVar(&accessKeyID, "id", "", "accessKeyID，默认为空")
	flag.StringVar(&secretAccessKey, "key", "", "secretAccessKey，默认为空")
	flag.StringVar(&filePath, "file", "./", "filePath")
	flag.StringVar(&objectName, "name", "", "objectName，默认空")
	flag.StringVar(&bucketName, "bucket", "", "bucketName，默认空")
	flag.BoolVar(&useSSL, "ssl", false, "是否使用ssl,默认为false")
	flag.Parse()
	if objectName == "" {
		objectName = filePath
	}
	FileUploader(endpoint, accessKeyID, secretAccessKey, filePath, objectName, bucketName, useSSL)
}
