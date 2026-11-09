package usecase

import (
	"context"

	"agnis-crm/internal/storage"
)

type Usecase struct {
	storage  *storage.Storage
	Contacts *Contacts
}

func New(st *storage.Storage) *Usecase {
	return &Usecase{storage: st, Contacts: NewContacts(st)}
}

// Health проверяет, что все нижележащие слои живы.
func (u *Usecase) Health(ctx context.Context) error {
	return u.storage.Ping(ctx)
}
