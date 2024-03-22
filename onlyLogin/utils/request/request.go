package request

import "github.com/gin-gonic/gin"

func GetParam(r *gin.Context, key string) (string, bool) {
	val := r.GetHeader(key)
	if val != "" {
		return val, true
	}
	val, err := r.Cookie(key)
	if err != nil {
		return "", false
	}
	return val, true
}
