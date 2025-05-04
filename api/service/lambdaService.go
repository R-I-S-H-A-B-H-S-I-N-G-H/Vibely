package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaService struct {
}

var lambdaObj *lambda.Lambda

// NewLambdaService initializes and returns a new LambdaService instance with static credentials.
func (l *LambdaService) GetLambdaClient() (*lambda.Lambda, error) {
	if lambdaObj != nil {
		return lambdaObj, nil
	}

	region := "us-east-1"
	accessKeyID := os.Getenv("LAMBDA_ACCESS_KEY")
	secretAccessKey := os.Getenv("LAMBDA_SECRET_KEY")

	fmt.Println()
	fmt.Println("refion : ", region)
	fmt.Println("accesskey", accessKeyID)
	fmt.Println("secretkey", secretAccessKey)
	fmt.Println()

	// Create a session using static credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	client := lambda.New(sess)
	lambdaObj = client

	return lambdaObj, nil
}

// InvokeLambda triggers a Lambda function with the provided function name and payload.
func (s *LambdaService) InvokeLambda(functionName string, payload any) ([]byte, error) {
	lambdaclient, err := s.GetLambdaClient()
	if err != nil {
		return nil, err
	}
	println()
	println("invoke lambda :: ", functionName, payload)
	println()

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	println()
	println("final lambda payload :: ", string(payloadBytes))
	println()

	// Invoke the Lambda function
	result, err := lambdaclient.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payloadBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to invoke Lambda function: %v", err)
	}

	// Check for errors in the response
	if *result.StatusCode != 200 {
		return nil, fmt.Errorf("error invoking Lambda: %s", string(result.Payload))
	}

	return result.Payload, nil
}
