package postgresql

import (
	"github.com/juicyluv/online-store/internal/infrastructure/database"
	"github.com/juicyluv/online-store/internal/user"
)

type repository struct {
	db database.Database
}

func NewUserRepository(db database.Database) user.Repository {
	return &repository{db: db}
}
