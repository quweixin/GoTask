package utils

import (
	"database/sql"
	"strconv"
	"strings"
)

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func ParamToUint(s string) uint {
	s = strings.TrimSpace(s)
	n64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	if n64 == 0 {
		return 0
	}
	return uint(n64)
}
