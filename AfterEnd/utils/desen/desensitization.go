package desen

import "strings"

// DesensitizationEmail 对邮箱进行信息脱敏
func DesensitizationEmail(email string) string {
	// 256655@qq.com  2****@qq.com
	// yaheb7479@yaho.com  y****@yaho.com
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

// DesensitizationTel 对电话号码进行信息脱敏
func DesensitizationTel(tel string) string {
	// 18387064785
	// 183****4785
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
