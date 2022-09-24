package domain_test

import (
	"encoder/domain"
	"github.com/stretchr/testify/require"
  uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestValidadeIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "not_uuid"
	video.ResorceID = "a"
	video.FilePath = "a"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4.String()
  video.ResorceID = "a"
	video.FilePath = "a"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
