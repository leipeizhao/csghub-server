package dataset

import (
	"github.com/gin-gonic/gin"
	"opencsg.com/starhub-server/pkg/types"
	"opencsg.com/starhub-server/pkg/utils/common"
)

func (c *Controller) Branches(ctx *gin.Context) (branches []*types.DatasetBranch, err error) {
	namespace, name, err := common.GetNamespaceAndNameFromContext(ctx)
	if err != nil {
		return
	}
	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		return
	}
	branches, err = c.gitServer.GetDatasetBranches(namespace, name, per, page)
	if err != nil {
		return
	}
	return
}
