package validate

import (
	"fmt"
	"regexp"
	"testing"
)

func TestValid(t *testing.T) {
	flag1 := Email("490@qq.com")
	flag2 := Phone("13265555490")
	fmt.Println(flag1, flag2)
}

func Email(source string) bool {
	flag, err := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", source)
	if err != nil {
		fmt.Println("Email正则表达式编译有误", err.Error())
	}
	return flag
}

func Phone(source string) bool {
	flag, err := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$", source)
	if err != nil {
		fmt.Println("Phone正则表达式编译有误", err.Error())
	}
	return flag
}
