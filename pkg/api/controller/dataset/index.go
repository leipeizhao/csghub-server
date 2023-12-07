package dataset

import (
	"git-devops.opencsg.com/product/community/starhub-server/pkg/store/database"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/utils/common"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Index(ctx *gin.Context) (datasets []database.Dataset, total int, err error) {
	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		return
	}
	datasets, err = c.datasetStore.Public(ctx, per, page)
	if err != nil {
		return
	}
	total, err = c.datasetStore.PublicCount(ctx)
	if err != nil {
		return
	}
	return
}
