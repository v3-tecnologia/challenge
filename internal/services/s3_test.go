package services

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type s3ClientInterface interface{}

type s3UploaderInterface interface {
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

type MockS3Uploader struct {
	mock.Mock
}

func (m *MockS3Uploader) Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*manager.UploadOutput), args.Error(1)
}

type TestS3Uploader struct {
	client   s3ClientInterface
	uploader s3UploaderInterface
}

func (up *TestS3Uploader) UploadAsync(ctx context.Context, file io.Reader, key string, ch chan<- string, errCh chan<- error) {
	defer close(ch)
	defer close(errCh)

	uploadedFile, err := up.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:         aws.String(os.Getenv("AWS_BUCKET_KEY_PREFIX") + key),
		Body:        file,
		ContentType: aws.String("image/png"),
		ACL:         "public-read",
	})

	if err != nil {
		errCh <- err
		return
	}

	ch <- uploadedFile.Location
}

func TestS3Uploader_UploadAsync_Success(t *testing.T) {
	t.Parallel()

	mockUploader := new(MockS3Uploader)

	uploader := &TestS3Uploader{
		client:   &s3.Client{},
		uploader: mockUploader,
	}

	ctx := context.Background()
	testFile := strings.NewReader("test file content")
	testKey := "test-key.png"
	testLocation := "https://test-bucket.s3.amazonaws.com/prefix/test-key.png"

	resultCh := make(chan string)
	errCh := make(chan error)

	expectedOutput := &manager.UploadOutput{
		Location: testLocation,
	}

	mockUploader.On("Upload", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
		return aws.ToString(input.Bucket) == "test-bucket" &&
			aws.ToString(input.Key) == "prefix/test-key.png" &&
			input.ContentType != nil && aws.ToString(input.ContentType) == "image/png" &&
			input.ACL == "public-read"
	})).Return(expectedOutput, nil)

	go uploader.UploadAsync(ctx, testFile, testKey, resultCh, errCh)

	select {
	case result := <-resultCh:
		assert.Equal(t, testLocation, result)
	case err := <-errCh:
		assert.Fail(t, "Expected success but got error", err)
	case <-time.After(time.Second):
		assert.Fail(t, "Test timed out")
	}

	mockUploader.AssertExpectations(t)
}

func TestS3Uploader_UploadAsync_Error(t *testing.T) {
	t.Parallel()

	mockUploader := new(MockS3Uploader)

	uploader := &TestS3Uploader{
		client:   &s3.Client{},
		uploader: mockUploader,
	}

	ctx := context.Background()
	testFile := strings.NewReader("test file content")
	testKey := "test-key.png"
	expectedError := errors.New("upload failed")

	resultCh := make(chan string)
	errCh := make(chan error)

	mockUploader.On("Upload", mock.Anything, mock.Anything).Return((*manager.UploadOutput)(nil), expectedError)

	go uploader.UploadAsync(ctx, testFile, testKey, resultCh, errCh)

	select {
	case <-resultCh:
		assert.Fail(t, "Expected error but got success")
	case err := <-errCh:
		assert.Equal(t, expectedError, err)
	case <-time.After(time.Second):
		assert.Fail(t, "Test timed out")
	}

	mockUploader.AssertExpectations(t)
}
