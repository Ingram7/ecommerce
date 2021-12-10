package app

import (
	"ecommerce/controller"
)

func (s *server) initRouter() {
	//repoUser := repository.NewUser()
	//srvUser := service.NewUser(repoUser)


	ctrUser := new(controller.User)
	user := s.handler.Group("/user")
	{
		user.POST("/login", ctrUser.Login)
	}
}
