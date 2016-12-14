package Response


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"wordsmap/common/Error"
)

// JSON output for response
func JSON(c *gin.Context, result interface{}, err error) {
	if err != nil {
		jsonError(c, err)
	} else {
		jsonSuccess(c, result)
	}
}

// jsonError for error output
func jsonError(c *gin.Context, e error) {
	m := gin.H{}
	switch t := e.(type) {
	case Error.Error:
		// need add response message
		m["code"] = t.Code
		m["message"] = t.Message
		c.JSON(t.Status, m)
	default:
		m["code"] = Error.UnexpectedException.Code
		m["message"] = Error.UnexpectedException.Message
		c.JSON(Error.UnexpectedException.Status, m)
	}
	c.Abort()
}

// JSONSuccess output for response
func jsonSuccess(c *gin.Context, result interface{}) {
	m := gin.H{
		"code":    "common.ok",
		"message": "OK.",
	}
	if result != nil {
		m["result"] = result
	}
	c.JSON(http.StatusOK, m)

	c.Abort()
}
