package presenter

import "NagoBackend/domain/model"

type MemberPresenter interface {
    ResponseMembers(members []*model.Member) []*model.Member
}
