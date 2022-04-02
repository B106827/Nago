package repository

import "NagoBackend/domain/model"

type MemberRepository interface {
    FindAll(posts []*model.Member) ([]*model.Member, error)
}
