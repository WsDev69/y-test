package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"io"
)

type Connection interface {
	Open() ObjectStorage
}

type connStr struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
}

type ObjectStorage interface {
	PutObject(ctx context.Context, opt *ObjectOptions) (string, error)
}

func (c *connStr) Open() ObjectStorage {
	log := logrus.WithFields(logrus.Fields{
		"minio-endpoint": c.endpoint,
	})
	minioClient, err := minio.New(c.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.accessKeyID, c.secretAccessKey, ""),
		Secure: c.useSSL,
	})
	if err != nil {
		log.Panicln("Couldn't connect to min.io: ", err)
	}

	return &StorageMinio{c: minioClient, endpoint: c.endpoint}
}

func InitConnection(endpoint, accessKeyID, secretAccessKey string, useSSL bool) Connection {
	return &connStr{endpoint, accessKeyID, secretAccessKey, useSSL}
}

type StorageMinio struct {
	endpoint string
	c        *minio.Client
}

type ObjectOptions struct {
	File   io.Reader
	Path   string
	Bucket string
	Name   string
	Size   int
}

func (s *StorageMinio) PutObject(ctx context.Context, opt *ObjectOptions) (string, error) {
	if opt == nil {
		err := fmt.Errorf("object options must be not nil")
		return "", err
	}

	log := logrus.WithContext(ctx)

	isExist, err := s.c.BucketExists(ctx, opt.Bucket)
	if err != nil || !isExist {
		log.WithField("bucket", opt.Bucket).Debug("Doesn't exist, err=", err)
		if err := s.c.MakeBucket(ctx, opt.Bucket, minio.MakeBucketOptions{
			ObjectLocking: false,
		}); err != nil {
			log.WithField("bucket", opt.Bucket).Debug("couldn't create, err=", err)
			return "", err
		}

	}

	uploadInfo, err := s.c.PutObject(context.Background(), opt.Bucket, opt.Name, opt.File, int64(opt.Size), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Error("couldn't upload object :", err)
		return "", err
	}

	log.Trace("successfully uploaded bytes: ", uploadInfo)
	return fmt.Sprintf("http://%s/%s/%s", s.endpoint, opt.Bucket, opt.Name), nil

}
