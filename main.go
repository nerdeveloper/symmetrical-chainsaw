package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/joho/godotenv"
	"io/ioutil"
	"strings"
)

func main() {
	// Load aws env vars
	godotenv.Load(".env")

	//config := config2.EnvConfig{
	//	Credentials: aws.Credentials{
	//		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
	//		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	//		SessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
	//	},
	//	Region: "us-west-2",
	//}

	client := ssm.NewFromConfig(aws.Config{})
	file, _ := ioutil.ReadFile(".secure.env")
	for _, secret := range strings.Split(string(file), "\n") {
		fmt.Println("Some secret ", secret)
		parameter, err := client.PutParameter(context.TODO(), {})
	}
}
