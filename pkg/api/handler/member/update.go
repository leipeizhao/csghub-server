package member

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"opencsg.com/starhub-server/pkg/api/controller/member"
)

func HandleUpdate(memberCtrl *member.Controller) func(*gin.Context) {
	return func(c *gin.Context) {
		member, err := memberCtrl.Update(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    "401",
				"message": fmt.Sprintf("Updated failed. %v", err),
			})
			return
		}

		respData := gin.H{
			"code":    200,
			"message": "member repository updated.",
			"data":    member,
		}

		c.JSON(http.StatusOK, respData)
	}
}
