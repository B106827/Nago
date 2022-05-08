package repository

import "NagoBackend/domain/model"

type UserRepository interface {
    FindAll(posts []*model.User) ([]*model.User, error)
}
