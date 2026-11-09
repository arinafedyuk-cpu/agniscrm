package storage

import (
	"agnis-crm/internal/storage/memory"
	"context"
)

type Storage struct{ *memory.Store }

func New() *Storage { return &Storage{Store: memory.NewWithSeed()} }

func (s *Storage) Ping(ctx context.Context) error { return nil }
