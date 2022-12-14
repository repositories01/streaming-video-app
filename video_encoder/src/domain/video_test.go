package domain_test

import (
	"encoder/src/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
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
	video.ResourceID = "a"
	video.FilePath = "a"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "a"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
