package controllers

import (
    "NagoBackend/domain/model"
    "NagoBackend/usecase/service"
)

type userController struct {
    userService service.UserService
}

type UserController interface {
    GetUsers() ([]*model.User, error)
}

func NewUserController(us service.UserService) UserController {
    return &userController{us}
}

func (userController *userController) GetUsers() ([]*model.User, error) {
    u := []*model.User{}
    us, err := userController.userService.Get(u)

    if err != nil {
        return nil, err
    }
    return us, nil
}
