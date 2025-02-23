package repo

import (
	"github.com/joshqu1985/lego/store/mysql"
	"github.com/joshqu1985/lego/store/redis"
)

type Repository struct {
}

func New(db *mysql.Store, c *redis.Store) *Repository {
	return &Repository{
	}
}
