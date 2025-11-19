package db

import (
	"context"
	"log/slog"
	"testing"
)

func TestStorage_CreateHosts_EmptyArray_Success(t *testing.T) {
	storage, err := NewStorage(slog.Default())
	if err != nil {
		t.Fatal(err)
	}

	hosts := []Host{}
	err = storage.CreateHosts(context.Background(), hosts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorage_DeleteHosts_EmptyArray_Success(t *testing.T) {
	storage, err := NewStorage(slog.Default())
	if err != nil {
		t.Fatal(err)
	}

	hosts := []string{}
	err = storage.DeleteHosts(context.Background(), hosts)
	if err != nil {
		t.Fatal(err)
	}
}
