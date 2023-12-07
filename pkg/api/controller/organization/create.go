package organization

import (
	"errors"

	"git-devops.opencsg.com/product/community/starhub-server/pkg/store/database"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/types"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Create(ctx *gin.Context) (org *database.Organization, err error) {
	var req types.CreateOrgReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	// Check if username exists
	user, err := c.userStore.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("User does not exists")
	}

	// Check if namespace exists
	exists, err := c.namespaceStore.Exists(ctx, req.Name)
	if exists {
		return nil, errors.New("Namespace has already existed")
	}

	req.User = user
	org, err = c.gitServer.CreateOrganization(&req)
	if err == nil {
		namespace := &database.Namespace{
			Path:   org.Path,
			UserID: user.ID,
		}
		c.orgStore.Create(ctx, org, namespace)
	}
	return
}
