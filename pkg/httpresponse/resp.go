package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code  int
	Data  interface{}
	Error string
}

func RespDefault(code int, data interface{}, err error) response {
	var msg string
	if err == nil {
		msg = ""
	} else {
		msg = err.Error()
	}

	return response{
		Code:  code,
		Data:  data,
		Error: msg,
	}
}

func StatusDefault(c *gin.Context, code int, data interface{}, err error) {
	resp := RespDefault(code, data, err)
	c.JSON(code, resp)
}

func RespOK(data interface{}) response {
	return RespDefault(0, data, nil)
}
func StatusOK(c *gin.Context, data interface{}) {
	resp := RespOK(data)
	c.JSON(http.StatusOK, resp)
}

func RespBadRequest(err error) response {
	return RespDefault(400, nil, err)
}

func StatusBadRequest(c *gin.Context, err error) {
	resp := RespBadRequest(err)
	c.JSON(http.StatusBadRequest, resp)
}

func RespNotFound(err error) response {
	return RespDefault(404, nil, err)
}

func StatusNotFound(c *gin.Context, err error) {
	resp := RespNotFound(err)
	c.JSON(http.StatusNotFound, resp)
}
