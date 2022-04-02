package controllers

import (
    "NagoBackend/domain/model"
    "NagoBackend/usecase/service"
)

type userController struct {
    userService service.MemberService
}

type MemberController interface {
    GetMembers() ([]*model.Member, error)
}

func NewMemberController(us service.MemberService) MemberController {
    return &userController{us}
}

func (userController *userController) GetMembers() ([]*model.Member, error) {
    u := []*model.Member{}
    us, err := userController.userService.Get(u)

    if err != nil {
        return nil, err
    }
    return us, nil
}
