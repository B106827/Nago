package service

import (
    "NagoBackend/domain/model"
    "NagoBackend/usecase/presenter"
    "NagoBackend/usecase/repository"
)

type userService struct {
    UserRepository repository.UserRepository
    UserPresenter  presenter.UserPresenter
}

type UserService interface {
    Get(users []*model.User) ([]*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
    return &userService{repo, pre}
}

func (userService *userService) Get(users []*model.User) ([]*model.User, error) {
    mems, err := userService.UserRepository.FindAll(users)
    if err != nil {
        return nil, err
    }
    return userService.UserPresenter.ResponseUsers(mems), nil
}
