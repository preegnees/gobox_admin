package memstorage

import (
	"context"
	dto "core/dto"
	"testing"
)

func TestSaveUserMetadata(t *testing.T) {
	r := New(nil)
	ctx := context.TODO()
	metadata := dto.DTOUserMetadata{
		Fingerprint:  "finger",
		RefreshToken: "token",
		Username:     "user",
	}
	if err := (*r).SaveUserMetadata(ctx, &metadata); err != nil {
		panic(err)
	}

	metadataNew, err := (*r).GetUserMetadata(ctx, metadata.Fingerprint)
	if err != nil {
		panic(err)
	}
	if metadata != *metadataNew {
		panic("metadata != *metadataNew")
	}
}
