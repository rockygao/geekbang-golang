package ut

import (
	"bytes"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

// LowerFirst 首字母转小写
func LowerFirst(str string) string {
	if len(strings.TrimSpace(str)) == 0 {
		return str
	}

	r := []rune(str)
	r[0] = unicode.ToLower(r[0])

	// for i, v := range str {
	// 	return string(unicode.ToLower(v)) + str[i+1:]
	// }
	return string(r)
}

// CamelToSnake 驼峰转换为蛇形
func CamelToSnake(s string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Infoln("err:", err)
		}
	}()

	buf := bytes.Buffer{}
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				buf.WriteByte('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func IsURL(s string) bool {
	s = strings.TrimSpace(s)
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
