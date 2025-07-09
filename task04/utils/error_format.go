package utils

import "strings"

func ErrorFormat(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		//field := "User.Name"
		//index := strings.Index(field, ".") // 返回 4
		//result := field[5:]
		//所以这一步就是把 "User.Name" 变成 "Name"。
		//将处理后的字段名作为 key，原错误信息作为 value 存入新的 map。
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
