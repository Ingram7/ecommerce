package request

type Captcha struct {
	Val string `form:"captcha" valid:"required~验证码不能为空,stringlength(4|4)~验证码4位"`
}

//func (req *Captcha) Validate() error {
//	_, err := validator.ValidateStruct(req)
//	return err
//}