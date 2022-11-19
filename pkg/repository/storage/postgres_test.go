package storage

import (
	"testing"
	"os"
)

func TestXxx(t *testing.T) {

	os.Setenv("POSTGRES_USER", "dev")
	os.Setenv("POSTGRES_PASSWORD", "dev")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "myapp")

	_, err := New()
	if err != nil {
		panic(err)
	}
}