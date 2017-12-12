package crawler

import (
	"testing"
	"time"
)

type Industry struct {
	ID        uint
	Title     string
	CreatedAt time.Time
}

func Test_createModule(t *testing.T) {
	CreateModule("industry", &Industry{})
}
