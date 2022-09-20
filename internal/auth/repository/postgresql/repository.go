package postgresql

import (
	"github.com/juicyluv/online-store/internal/auth"
	"github.com/juicyluv/online-store/internal/infrastructure/database"
)

type repository struct {
	db database.Database
}

func NewAuthRepository(db database.Database) auth.Repository {
	return &repository{db: db}
}
