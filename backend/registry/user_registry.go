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

type userInteractor struct {
    conn *gorm.DB
}

type MemberInteractor interface {
    NewMemberHandler() handler.MemberHandler
}

func NewMemberInteractor(conn *gorm.DB) MemberInteractor {
    return &userInteractor{conn}
}

func (i *userInteractor) NewMemberHandler() handler.MemberHandler {
    return handler.NewMemberHandler(i.NewMemberController())
}

func (i *userInteractor) NewMemberController() controllers.MemberController {
    return controllers.NewMemberController(i.NewMemberService())
}

func (i *userInteractor) NewMemberService() service.MemberService {
    return service.NewMemberService(i.NewMemberRepository(), i.NewMemberPresenter())
}

func (i *userInteractor) NewMemberRepository() repository.MemberRepository {
    return datastore.NewMemberRepository(i.conn)
}

func (i *userInteractor) NewMemberPresenter() presenter.MemberPresenter {
    return presenters.NewMemberPresenter()
}
