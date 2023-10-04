package repo

import (
	"BWG/entity"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db    *sqlx.DB
	cache *RedisRepository
}

func NewUserPostgres(db *sqlx.DB, cache *RedisRepository) *UserPostgres {
	return &UserPostgres{db: db, cache: cache}
}

func (r *UserPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, surname)"+
		" values ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Surname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) GetUser(id int) (entity.User, error) {
	var user entity.User

	cacheKey := fmt.Sprintf("user_%d", id)
	cachedData, err := r.cache.Get(cacheKey)
	if err == nil {
		if err := json.Unmarshal([]byte(cachedData), &user); err == nil {
			return user, nil
		}
	}

	query := fmt.Sprintf("SELECT id, username, surname FROM %s WHERE id=$1", userTable)
	err = r.db.Get(&user, query, id)

	if err == nil {
		userData, _ := json.Marshal(user)
		r.cache.Set(cacheKey, string(userData))
	}

	return user, err
}

func (r *UserPostgres) GetUsers(query string) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Select(&users, query)

	if err == nil {
		for _, user := range users {
			cacheKey := fmt.Sprintf("user_%d", user.Id)
			userData, _ := json.Marshal(user)
			r.cache.Set(cacheKey, string(userData))
		}
	}

	return users, err
}
