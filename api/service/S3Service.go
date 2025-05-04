package service

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
}

func (s *S3Service) getS3Client() *s3.S3 {
	// Hardcoded AWS credentials and region
	awsAccessKey := os.Getenv("S3_ACCESS_KEY")
	awsSecretKey := os.Getenv("S3_SECRET_KEY")
	awsRegion := os.Getenv("S3_REGION")
	endpoint := getEndpoint()

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(awsRegion),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)
	return s3.New(newSession)
}

func GetObjectPath(filename string) string {
	return fmt.Sprintf("%s/%s/%s", getEndpoint(), getS3Bucket(), filename)
}

func getS3Bucket() string {
	return os.Getenv("S3_BUCKET")
}

func getEndpoint() string {
	return os.Getenv("S3_ENDPOINT")
}

func (s *S3Service) UploadStrDataToS3(objectKey string, str string) (string, error) {
	data := []byte(str)
	return s.UploadDataToS3(objectKey, data)
}

func (s *S3Service) UploadDataToS3(objectKey string, data []byte) (string, error) {
	fmt.Println("bucket : :", getS3Bucket())
	svc := s.getS3Client()

	// Upload the data
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(getS3Bucket()),
		Key:                aws.String(objectKey),
		Body:               bytes.NewReader(data),
		ContentDisposition: aws.String("inline"), // Set Content-Disposition to inline
	})

	if err != nil {
		return "", err
	}

	// Return the object path
	return GetObjectPath(objectKey), nil
}

// getContentType returns the MIME type based on the file extension.
func getContentType(filename string) string {
	// Map of file extensions to their corresponding MIME types
	contentTypes := map[string]string{
		".txt":  "text/plain",
		".html": "text/html",
		".json": "application/json",
		".xml":  "application/xml",
		".pdf":  "application/pdf",
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".csv":  "text/csv",
		".zip":  "application/zip",
		".mp4":  "video/mp4",
		".mp3":  "audio/mpeg",
		// Add more types as needed
	}

	// Get the file extension
	ext := strings.ToLower(filename[strings.LastIndex(filename, "."):])

	// Return the corresponding content type, or a default if not found
	if contentType, exists := contentTypes[ext]; exists {
		return contentType
	}
	return "text/" + strings.Split(ext, ".")[1] // Default content type for unknown extensions
}

func (s *S3Service) GeneratePresignedGetURL(objectKey string, expiryDuration int64) (string, error) {
	// Create an S3 client
	svc := s.getS3Client()

	// Generate a presigned URL for the object
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(objectKey),
	})

	// Set the expiry duration for the URL
	presignedURL, err := req.Presign(time.Duration(expiryDuration) * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return presignedURL, nil
}

func (s *S3Service) GeneratePresignedPutURL(objectKey string, expiryDuration int64) (string, error) {
	// Create an S3 client
	svc := s.getS3Client()

	// Generate a presigned PUT URL
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(objectKey),
	})

	presignedURL, err := req.Presign(time.Duration(expiryDuration) * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned PUT URL: %w", err)
	}

	return presignedURL, nil
}

func (s *S3Service) ObjectExists(objectKey string) (bool, error) {
	svc := s.getS3Client()

	_, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		// Check for a not found error
		if aerr, ok := err.(awserr.RequestFailure); ok && aerr.StatusCode() == 404 {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if object exists: %w", err)
	}

	return true, nil
}
