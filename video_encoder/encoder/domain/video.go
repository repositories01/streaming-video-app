package domain
import "time"

type Video struct {
  ID        string
  ResorceID string
  FilePath  string
  CreatedAt time.Time
}
