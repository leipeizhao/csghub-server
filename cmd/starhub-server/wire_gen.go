// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"opencsg.com/starhub-server/cmd/starhub-server/cmd/common"
	"opencsg.com/starhub-server/pkg/api/controller/accesstoken"
	"opencsg.com/starhub-server/pkg/api/controller/dataset"
	"opencsg.com/starhub-server/pkg/api/controller/member"
	model2 "opencsg.com/starhub-server/pkg/api/controller/model"
	"opencsg.com/starhub-server/pkg/api/controller/organization"
	"opencsg.com/starhub-server/pkg/api/controller/sshkey"
	"opencsg.com/starhub-server/pkg/api/controller/user"
	"opencsg.com/starhub-server/pkg/apiserver"
	"opencsg.com/starhub-server/pkg/gitserver"
	"opencsg.com/starhub-server/pkg/httpbase"
	"opencsg.com/starhub-server/pkg/model"
	"opencsg.com/starhub-server/pkg/router"
	"opencsg.com/starhub-server/pkg/store/cache"
	"opencsg.com/starhub-server/pkg/store/database"
)

// Injectors from wire.go:

func initAPIServer(ctx context.Context) (*httpbase.GracefulServer, error) {
	config, err := common.ProvideConfig()
	if err != nil {
		return nil, err
	}
	logger := apiserver.ProvideServerLogger()
	dbConfig := model.ProvideDBConfig(config)
	db, err := model.ProvideDatabse(ctx, dbConfig)
	if err != nil {
		return nil, err
	}
	modelStore := database.ProvideModelStore(db)
	redisConfig := cache.ProvideRedisConfig(config)
	cacheCache, err := cache.ProvideCache(ctx, redisConfig)
	if err != nil {
		return nil, err
	}
	modelCache := cache.ProvideModelCache(cacheCache)
	userStore := database.ProvideUserStore(db)
	userCache := cache.ProvideUserCache(cacheCache)
	orgStore := database.ProvideOrgStore(db)
	orgCache := cache.ProvideOrgCache(cacheCache)
	namespaceStore := database.ProvideNamespaceStore(db)
	namespaceCache := cache.ProvideNamespaceCache(cacheCache)
	repoStore := database.ProvideRepoStore(db)
	repoCache := cache.ProvideRepoCache(cacheCache)
	gitServer, err := gitserver.ProvideGitServer(config)
	if err != nil {
		return nil, err
	}
	controller := model2.ProvideController(modelStore, modelCache, userStore, userCache, orgStore, orgCache, namespaceStore, namespaceCache, repoStore, repoCache, gitServer)
	datasetStore := database.ProvideDatasetStore(db)
	datasetCache := cache.ProvideDatasetCache(cacheCache)
	datasetController := dataset.ProvideController(datasetStore, datasetCache, userStore, userCache, orgStore, orgCache, namespaceStore, namespaceCache, repoStore, repoCache, gitServer)
	userController := user.ProvideController(userStore, userCache, modelStore, modelCache, datasetStore, datasetCache, namespaceStore, namespaceCache, gitServer)
	accessTokenStore := database.ProvideAccessTokenStore(db)
	accessTokenCache := cache.ProvideAccessTokenCache(cacheCache)
	accesstokenController := accesstoken.ProvideController(userStore, userCache, accessTokenStore, accessTokenCache, gitServer)
	sshKeyStore := database.ProvideSSHKeyStore(db)
	sshKeyCache := cache.ProvideSSHKeyCache(cacheCache)
	sshkeyController := sshkey.ProvideController(sshKeyStore, sshKeyCache, userStore, gitServer)
	memberStore := database.ProvideMemberStore(db)
	memberCache := cache.ProvideMemberCache(cacheCache)
	organizationController := organization.ProvideController(memberStore, memberCache, orgStore, orgCache, userStore, userCache, namespaceStore, namespaceCache, gitServer)
	memberController := member.ProvideController(memberStore, memberCache, orgStore, orgCache, gitServer)
	apiHandler := router.ProvideAPIHandler(config, controller, datasetController, userController, accesstokenController, sshkeyController, organizationController, memberController)
	gitHandler := router.ProvideGitHandler(config, controller, datasetController)
	routerRouter := router.ProvideRouter(apiHandler, gitHandler)
	gracefulServer := apiserver.ProvideGracefulServer(config, logger, routerRouter)
	return gracefulServer, nil
}
