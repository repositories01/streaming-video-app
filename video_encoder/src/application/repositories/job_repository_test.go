package repositories_test

import (
	"encoder/src/application/repositories"
	"encoder/src/domain"
	"encoder/src/infra/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJobRepositoryDbInsert(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)
	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, j)
	require.Equal(t, job.ID, j.ID)
	require.Equal(t, video.ID, j.Video.ID)
}
func TestJobRepositoryDbUpdate(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)
	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "complete"
	repoJob.Update(job)
	j, err := repoJob.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, j)
	require.Equal(t, job.ID, j.ID)
	require.Equal(t, video.ID, j.Video.ID)
	require.Equal(t, job.Status, j.Status)
}
