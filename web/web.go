package web

import (
	"fmt"
	"net/http"
	"zwh/ghost/stringx"
	"zwh/process"

	"github.com/gin-gonic/gin"
)

var Root string

func Dirs(c *gin.Context) {
	path := c.Query("path")
	path = stringx.TrimPrefixSlash(fmt.Sprintf("%s/%s", Root, path))
	r, err := process.GetAllFile(path)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("walk dir failed"))
	}
	c.JSON(http.StatusOK, r)
}
