package storage

import "bytes"

type IStorage interface {
	DownloadFile(filePath string) []bytes.Buffer
}
