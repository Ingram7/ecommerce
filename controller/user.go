package controller

import (
	"ecommerce/model/request"
	"ecommerce/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	svs service.User
}

//@Summary 用户登陆 code=401 重新登陆
//@Accept application/json
//@Produce application/json
//@Param   mobile query string true "手机号"
//@Param   captcha query string true "验证码"
//@Success 200 {object} domain.User
//@Router /user/login [post]
func (ctr *User) Login(c *gin.Context) {

	userLogin := new(request.UserLogin)
	if err := c.Bind(userLogin); err != nil {
		fmt.Println(err)
	}
	fmt.Println(userLogin)
}