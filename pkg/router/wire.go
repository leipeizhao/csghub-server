package router

import (
	"git-devops.opencsg.com/product/community/starhub-server/config"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/accesstoken"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/dataset"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/member"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/model"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/organization"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/sshkey"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/api/controller/user"
	"github.com/google/wire"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideAPIHandler,
	ProvideGitHandler,
	ProvideRouter,
)

func ProvideAPIHandler(
	config *config.Config,
	modelCtrl *model.Controller,
	datasetCtrl *dataset.Controller,
	userCtrl *user.Controller,
	acCtrl *accesstoken.Controller,
	sshKeyCtrl *sshkey.Controller,
	orgCtrl *organization.Controller,
	memberCtrl *member.Controller,
) APIHandler {
	return NewAPIHandler(config, modelCtrl, datasetCtrl, userCtrl, acCtrl, sshKeyCtrl, orgCtrl, memberCtrl)
}

func ProvideGitHandler(
	config *config.Config,
	modelCtrl *model.Controller,
	datasetCtrl *dataset.Controller,
) GitHandler {
	return NewGitHandler(config, modelCtrl, datasetCtrl)
}

func ProvideRouter(api APIHandler, git GitHandler) *Router {
	return NewRouter(api, git)
}
