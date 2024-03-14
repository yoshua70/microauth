package store

import (
	"github.com/yoshua70/microauth/models"
)

type MemStore struct {
	Store map[string]models.User
}
