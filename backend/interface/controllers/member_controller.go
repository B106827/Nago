package controllers

import (
    "NagoBackend/domain/model"
    "NagoBackend/usecase/service"
)

type memberController struct {
    memberService service.MemberService
}

type MemberController interface {
    GetMembers() ([]*model.Member, error)
}

func NewMemberController(us service.MemberService) MemberController {
    return &memberController{us}
}

func (memberController *memberController) GetMembers() ([]*model.Member, error) {
    u := []*model.Member{}
    us, err := memberController.memberService.Get(u)

    if err != nil {
        return nil, err
    }
    return us, nil
}
