package repository

import "github.com/oooiik/test_09.03.2024/internal/database"

type Interface interface {
	connDB(p database.Interface)
}

type repository struct {
	database database.Interface
}

func (r *repository) connDB(d database.Interface) {
	r.database = d
}
