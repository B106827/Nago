package service

import (
    "NagoBackend/domain/model"
    "NagoBackend/usecase/presenter"
    "NagoBackend/usecase/repository"
)

type memberService struct {
    MemberRepository repository.MemberRepository
    MemberPresenter  presenter.MemberPresenter
}

type MemberService interface {
    Get(members []*model.Member) ([]*model.Member, error)
}

func NewMemberService(repo repository.MemberRepository, pre presenter.MemberPresenter) MemberService {
    return &memberService{repo, pre}
}

func (memberService *memberService) Get(members []*model.Member) ([]*model.Member, error) {
    mems, err := memberService.MemberRepository.FindAll(members)
    if err != nil {
        return nil, err
    }
    return memberService.MemberPresenter.ResponseMembers(mems), nil
}
