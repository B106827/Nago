package presenters

import (
    "NagoBackend/domain/model"
)

type userPresenter struct {
}

func NewMemberPresenter() MemberPresenter {
    return &userPresenter{}
}

type MemberPresenter interface {
    ResponseMembers(users []*model.Member) []*model.Member
}

func (userPresenter *userPresenter) ResponseMembers(users []*model.Member) []*model.Member {
    for _, u := range users {
        u.Name = u.Name + "hoge"
    }
    return users
}
