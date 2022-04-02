package presenter

import "NagoBackend/domain/model"

type MemberPresenter interface {
    ResponseMembers(users []*model.Member) []*model.Member
}
