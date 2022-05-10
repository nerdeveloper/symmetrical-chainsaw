package main

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {

	// Load aws env vars
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, "")},
	)

	file, err := os.Open(".secure.env")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				// add key/value to ssm
				key := parts[0]
				value := parts[1]
				fmt.Println(key, value)
				svc := ssm.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
				_, err := svc.PutParameter(&ssm.PutParameterInput{
					Name:      aws.String(key),
					Value:     aws.String(value),
					Type:      aws.String("SecureString"),
					Overwrite: aws.Bool(true),
				})
				if err != nil {
					fmt.Println(err)

				}
			}

		}

	}
}
