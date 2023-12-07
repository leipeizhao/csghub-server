package user

import (
	"errors"

	"git-devops.opencsg.com/product/community/starhub-server/pkg/store/database"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/utils/common"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Datasets(ctx *gin.Context) (datasets []database.Dataset, err error) {
	currentUser := ctx.Query("current_user")
	username := ctx.Param("username")
	_, err = c.userStore.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("User does not exist")
	}

	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		return nil, err
	}

	onlyPublic := currentUser != username

	datasets, err = c.datasetStore.ByUsername(ctx, username, per, page, onlyPublic)
	if err != nil {
		return nil, err
	}

	return
}
