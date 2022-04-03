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

type UserInteractor interface {
    NewUserHandler() handler.UserHandler
}

func NewUserInteractor(conn *gorm.DB) UserInteractor {
    return &userInteractor{conn}
}

func (i *userInteractor) NewUserHandler() handler.UserHandler {
    return handler.NewUserHandler(i.NewUserController())
}

func (i *userInteractor) NewUserController() controllers.UserController {
    return controllers.NewUserController(i.NewUserService())
}

func (i *userInteractor) NewUserService() service.UserService {
    return service.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *userInteractor) NewUserRepository() repository.UserRepository {
    return datastore.NewUserRepository(i.conn)
}

func (i *userInteractor) NewUserPresenter() presenter.UserPresenter {
    return presenters.NewUserPresenter()
}
