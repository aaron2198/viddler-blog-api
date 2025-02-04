package bucket

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

// BucketStore provides a simple abstraction for S3-compatible object storage
type BucketStore struct {
	client *s3.Client
	bucket string
}

type BucketStoreConfig struct {
	Type      string
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
	Bucket    string
}

type MinioResolver struct {
	endpoint string
}

func New(cfg *BucketStoreConfig) (*BucketStore, error) {
	switch cfg.Type {
	case "minio":
		return NewMinioBucket(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey, cfg.Bucket)
	default:
		return NewAWSBucket(cfg.AccessKey, cfg.SecretKey, cfg.Region, cfg.Bucket)
	}
}

func NewMinioResolver(endpoint string) *MinioResolver {
	return &MinioResolver{endpoint: endpoint}
}

func (m *MinioResolver) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {
	uri, err := url.Parse(m.endpoint)
	if err != nil {
		fmt.Println(err)
		return s3.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
	}
	return smithyendpoints.Endpoint{
		URI: *uri,
	}, nil
}

// NewMinioBucket creates a new BucketStore instance configured for MinIO
func NewMinioBucket(endpoint, accessKey, secretKey, bucket string) (*BucketStore, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load MinIO config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.EndpointResolverV2 = NewMinioResolver(endpoint)
	})
	return newBucketStore(client, bucket), nil
}

// NewAWSBucket creates a new BucketStore instance configured for AWS S3
func NewAWSBucket(accessKey, secretKey, region, bucket string) (*BucketStore, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	return newBucketStore(client, bucket), nil
}

// NewBucketStore creates a new BucketStore instance
func newBucketStore(client *s3.Client, bucket string) *BucketStore {
	return &BucketStore{
		client: client,
		bucket: bucket,
	}
}

// GetObject retrieves an object from the bucket
func (b *BucketStore) GetObject(ctx context.Context, key string) (io.ReadCloser, error) {
	result, err := b.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &b.bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s: %w", key, err)
	}
	return result.Body, nil
}

func (b *BucketStore) GetFileAsString(ctx context.Context, key string) (string, error) {
	body, err := b.GetObject(ctx, key)
	if err != nil {
		return "", err
	}
	defer body.Close()
	content, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (b *BucketStore) ObjectExists(ctx context.Context, key string) (bool, error) {
	_, err := b.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &b.bucket,
		Key:    &key,
	})
	return err == nil, nil
}

func (b *BucketStore) PutObjectFromFs(ctx context.Context, source, key string) error {
	file, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", source, err)
	}
	defer file.Close()
	return b.PutObject(ctx, key, file)
}

// PutObject uploads an object to the bucket
func (b *BucketStore) PutObject(ctx context.Context, key string, body io.Reader) error {
	input := s3.PutObjectInput{
		Bucket: &b.bucket,
		Key:    &key,
		Body:   body,
	}
	_, err := b.client.PutObject(ctx, &input)
	if err != nil {
		return fmt.Errorf("failed to put object %s: %w", key, err)
	}
	return nil
}

func (b *BucketStore) PublicUrl(key string) (string, error) {
	e, err := b.client.Options().EndpointResolverV2.ResolveEndpoint(context.Background(), s3.EndpointParameters{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", e.URI.String(), key), nil
}
