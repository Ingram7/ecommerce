package request

import (
	"errors"
	"regexp"
)

type Mobile struct {
	Val string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,stringlength(11|11)~手机号码长度错误"`
}

func (req *Mobile) Validate() error {
	if !req.valid() {
		return errors.New("手机号码格式错误")
	}

	return nil
}

func (req *Mobile) valid() bool {
	b, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, req.Val)
	return b
}