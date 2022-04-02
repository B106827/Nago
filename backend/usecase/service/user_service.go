package service

import (
    "NagoBackend/domain/model"
    "NagoBackend/usecase/presenter"
    "NagoBackend/usecase/repository"
)

type userService struct {
    MemberRepository repository.MemberRepository
    MemberPresenter  presenter.MemberPresenter
}

type MemberService interface {
    Get(users []*model.Member) ([]*model.Member, error)
}

func NewMemberService(repo repository.MemberRepository, pre presenter.MemberPresenter) MemberService {
    return &userService{repo, pre}
}

func (userService *userService) Get(users []*model.Member) ([]*model.Member, error) {
    mems, err := userService.MemberRepository.FindAll(users)
    if err != nil {
        return nil, err
    }
    return userService.MemberPresenter.ResponseMembers(mems), nil
}
