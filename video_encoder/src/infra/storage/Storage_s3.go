package storage

type StorageS3 struct {
	BucketName string
}

func (s *StorageS3) NewStorageS3(bucketName string) *StorageS3 {
	return &StorageS3{BucketName: bucketName}
}

func (s *StorageS3) DownloadFile(filePath string) {
}
