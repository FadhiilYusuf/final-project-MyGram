package helpers

import "github.com/gin-gonic/gin"

type RequestHeader struct {
	ContentType   string
	Authorization string
}

func GetRequestHeaders(c *gin.Context) RequestHeader {
	return RequestHeader{
		ContentType:   c.Request.Header.Get("Content-Type"),
		Authorization: c.Request.Header.Get("Authorization"),
	}
}
