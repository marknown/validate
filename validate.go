package validate

import "regexp"

// IsMobile 是否是手机号
func IsMobile(s string) bool {
	re := regexp.MustCompile(`^1(3|4|5|6|7|8|9)[0-9]{1}\d{8}$`)
	return re.MatchString(s)
}

// IsEmail 是否是邮箱地址
func IsEmail(s string) bool {
	re := regexp.MustCompile(`^([a-z0-9\+_\-]+)(\.[a-z0-9\+_\-]+)*@([a-z0-9\-]+\.)+[a-z]{2,6}$`)
	return re.MatchString(s)
}
