package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var bucketConfig BucketConfig

func main() {

	loadConfig()

	Headers := []Header{Header{
		Name:  "filePath",
		Value: "/home/laertida/agenda/architecture.org",
	}}

	paths := Message{
		Headers: Headers,
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketConfig.Bucket.Name),
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, object := range output.Contents {
		log.Printf("key=%s, size=%v\n date=%v", aws.ToString(object.Key), aws.ToString(object.ETag), object.LastModified)
	}

	log.Printf("this is the bucket name from config file: %v\n", bucketConfig.Bucket.Name)

}

func loadConfig() {
	_, err := os.Stat("./my-cloud-box.yaml")
	if err != nil {
		log.Fatal(err)
	}

	configFile, error := os.ReadFile("./my-cloud-box.yaml")

	if error != nil {
		log.Printf("There was an error reading, config file: %v", error)
	}

	yaml.Unmarshal(configFile, &bucketConfig)
}

func putObject(client s3.Client, key string, filePath string) {

	file, err := os.Open(filePath)

	if err != nil {
		log.Printf("there was an error open file %v", file)
	}

	client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketConfig.Bucket.Name),
		Key:    aws.String(key),
		Body:   file,
	})
}

func SyncUp(path string) {

	pathProperties := []Property{Property{
		Key:   "move",
		Value: ".done",
	}}

	Path := Endpoint{Schema: "file", Path: path, Properties: pathProperties}

	S3 := Endpoint{Schema: "s3", Path: path}

	log.Println(Path, S3)
}

type BucketConfig struct {
	Bucket Bucket `yaml:"bucket"`
}

type Bucket struct {
	Name   string `yaml:"name"`
	Region string `yaml:"region"`
}

type Endpoint struct {
	Schema     string
	Path       string
	Properties []Property
	Exchange   Exchange
}

type Message struct {
	Headers     []Header
	Attachments map[string]string
	Body        os.File
}

type Header struct {
	Name  string
	Value string
}

type Property struct {
	Key   string
	Value string
}

type Exchange struct {
	in         Message
	out        Message
	Properties map[string]string
}
