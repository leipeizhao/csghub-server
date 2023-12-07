package gitserver

import (
	"errors"

	"git-devops.opencsg.com/product/community/starhub-server/config"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/gitserver/gitea"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/store/database"
	"git-devops.opencsg.com/product/community/starhub-server/pkg/types"
)

type GitServer interface {
	CreateUser(*types.CreateUserRequest) (*database.User, error)
	UpdateUser(*types.UpdateUserRequest) (int, *database.User, error)
	CreateUserToken(*types.CreateUserTokenRequest) (*database.AccessToken, error)
	DeleteUserToken(*types.DeleteUserTokenRequest) error

	CreateModelRepo(*types.CreateModelReq) (*database.Model, *database.Repository, error)
	UpdateModelRepo(string, string, *database.Model, *database.Repository, *types.UpdateModelReq) error
	DeleteModelRepo(string, string) error
	GetModelBranches(string, string, int, int) ([]*types.ModelBranch, error)
	GetModelCommits(string, string, string, int, int) ([]*types.Commit, error)
	GetModelLastCommit(string, string, string) (*types.Commit, error)
	GetModelDetail(namespace, name string) (*types.ModelDetail, error)
	GetModelFileRaw(string, string, string, string) (string, error)
	GetModelTags(string, string, int, int) ([]*types.ModelTag, error)
	GetModelFileTree(string, string, string, string) ([]*types.File, error)

	CreateDatasetRepo(*types.CreateDatasetReq) (*database.Dataset, *database.Repository, error)
	UpdateDatasetRepo(string, string, *database.Dataset, *database.Repository, *types.UpdateDatasetReq) error
	DeleteDatasetRepo(string, string) error
	GetDatasetBranches(string, string, int, int) ([]*types.DatasetBranch, error)
	GetDatasetCommits(string, string, string, int, int) ([]*types.Commit, error)
	GetDatasetLastCommit(string, string, string) (*types.Commit, error)
	GetDatasetDetail(namespace, name string) (*types.DatasetDetail, error)
	GetDatasetFileRaw(string, string, string, string) (string, error)
	GetDatasetTags(string, string, int, int) ([]*types.DatasetTag, error)
	GetDatasetFileTree(string, string, string, string) ([]*types.File, error)

	CreateSSHKey(*types.CreateSSHKeyRequest) (*database.SSHKey, error)
	// ListSSHKeys(string, int, int) ([]*database.SSHKey, error)
	DeleteSSHKey(int) error

	CreateOrganization(*types.CreateOrgReq) (*database.Organization, error)
	DeleteOrganization(string) error
	UpdateOrganization(string, *types.EditOrgReq, *database.Organization) (*database.Organization, error)
}

func NewGitServer(config *config.Config) (GitServer, error) {
	if config.GitServer.Type == "gitea" {
		gitServer, err := gitea.NewClient(config)
		return gitServer, err
	}

	return nil, errors.New("Undefined git server type.")
}
