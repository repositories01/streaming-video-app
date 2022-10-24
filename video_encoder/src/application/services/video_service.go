package services

import (
	"encoder/src/application/repositories"
	"encoder/src/domain"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

//func (v *VideoService) Download(bucketName string) error {
//
//}
