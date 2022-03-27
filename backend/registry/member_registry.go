package registry

import (
    "github.com/jinzhu/gorm"
    "NagoBackend/infrastructure/api/handler"
    "NagoBackend/infrastructure/datastore"
    "NagoBackend/interface/controllers"
    "NagoBackend/interface/presenters"
    "NagoBackend/usecase/presenter"
    "NagoBackend/usecase/repository"
    "NagoBackend/usecase/service"
)

type memberInteractor struct {
    conn *gorm.DB
}

type MemberInteractor interface {
    NewMemberHandler() handler.MemberHandler
}

func NewMemberInteractor(conn *gorm.DB) MemberInteractor {
    return &memberInteractor{conn}
}

func (i *memberInteractor) NewMemberHandler() handler.MemberHandler {
    return handler.NewMemberHandler(i.NewMemberController())
}

func (i *memberInteractor) NewMemberController() controllers.MemberController {
    return controllers.NewMemberController(i.NewMemberService())
}

func (i *memberInteractor) NewMemberService() service.MemberService {
    return service.NewMemberService(i.NewMemberRepository(), i.NewMemberPresenter())
}

func (i *memberInteractor) NewMemberRepository() repository.MemberRepository {
    return datastore.NewMemberRepository(i.conn)
}

func (i *memberInteractor) NewMemberPresenter() presenter.MemberPresenter {
    return presenters.NewMemberPresenter()
}
