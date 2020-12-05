package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pervasive-chain/statecode"
	"reflect"
)

func FailResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": statecode.Fail, "msg": statecode.CodeInfo(statecode.Fail)})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": statecode.Success, "msg": statecode.CodeInfo(statecode.Success), "data": data})
}

func ResponseWithCode(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": statecode.CodeInfo(code)})
}


// todo
func WsFailResponse(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": statecode.Fail, "msg": statecode.CodeInfo(statecode.Fail)})
}

func WsSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": statecode.Success, "msg": statecode.CodeInfo(statecode.Success), "data": data})
}

func WsResponseWithCode(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": statecode.CodeInfo(code)})
}

func Response(v *interface{}) {
	reflect.TypeOf(*v)
}
