package presenter

import "NagoBackend/domain/model"

type UserPresenter interface {
    ResponseUsers(users []*model.User) []*model.User
}
