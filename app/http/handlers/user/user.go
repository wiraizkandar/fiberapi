package user

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetUser(c *fiber.Ctx) error {
	return c.SendString("This is Get Users 22")
}

func GetUserById(c *fiber.Ctx) error {
	return c.SendString("This user id is " + c.Params("id"))
}

func SendVerification(c *fiber.Ctx) error {

	// Define your AWS credentials. Replace these values with your own.
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_SQS_KEY"), os.Getenv("AWS_SQS_SECRET"), "")),
		config.WithRegion(os.Getenv("AWS_SQS_REGION")),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	svc := sqs.NewFromConfig(cfg)

	// Define the input parameters for sending a message.
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(os.Getenv("AWS_SQS_URL")), // Replace with your SQS queue URL.
		MessageBody: aws.String("Hello, SQS!"),            // Replace with your message body.
	}

	_, err = svc.SendMessage(context.TODO(), input)
	if err != nil {
		return err
	}

	data := make(map[string]string)

	responseData := map[string]interface{}{
		"message": "Verification Code Sent",
		"data":    data,
	}

	return c.Status(fiber.StatusOK).JSON(responseData)
}
