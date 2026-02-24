package users

import (
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/cache/redis"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/database"
)

type UserRepository struct {
	db    *database.Postres
	cache *redis.Redis
}

type Params struct {
	DB    *database.Postres
	Cache *redis.Redis
}

func NewUserRepository(p Params) *UserRepository {
	return &UserRepository{p.DB, p.Cache}
}

func (u *UserRepository) Create() {

}

func (u *UserRepository) Get() {

}

func (u *UserRepository) Update() {

}
