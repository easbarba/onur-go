package storage_test

import (
	"testing"

	"gitlab.com/easbarba/onur/internal/storage"
)

func TestFiles(t *testing.T) {
	files := storage.Files()

	got := len(files)
	expected := 4

	if got != expected {
		t.Errorf("%d and %d ", got, expected)
	}
}
