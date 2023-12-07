package organization

import (
	"fmt"
	"net/http"

	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/organization"
	"github.com/gin-gonic/gin"
)

func HandleIndex(orgCtrl *organization.Controller) func(*gin.Context) {
	return func(c *gin.Context) {
		orgs, err := orgCtrl.Index(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": fmt.Sprintf("Get organization list failed. %v", err),
			})
			return
		}

		respData := gin.H{
			"code":    200,
			"message": "Get organization list successfully.",
			"data":    orgs,
		}

		c.JSON(http.StatusOK, respData)
	}
}
