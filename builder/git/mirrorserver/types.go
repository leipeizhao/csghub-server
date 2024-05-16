package mirrorserver

import (
	"opencsg.com/csghub-server/common/types"
)

// enumerate all the kinds of task status
const (
	TaskStatusQueued   TaskStatus = iota // 0 task is queued
	TaskStatusRunning                    // 1 task is running
	TaskStatusStopped                    // 2 task is stopped (never used)
	TaskStatusFailed                     // 3 task is failed
	TaskStatusFinished                   // 4 task is finished
)

type CreateMirrorRepoReq struct {
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	CloneUrl    string `json:"clone_url"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	Private     bool   `json:"private"`
	Description string `json:"description"`
	Interval    string `json:"interval"`
}

type CreatePushMirrorReq struct {
	Name        string `json:"name"`
	PushUrl     string `json:"push_url"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	Interval    string `json:"interval"`
}

type MirrorSyncReq struct {
	Namespace string               `json:"namespace"`
	Name      string               `json:"name"`
	RepoType  types.RepositoryType `json:"repo_type"`
}

type MirrorTaskInfo struct {
	Status    TaskStatus `json:"status"`
	Message   string     `json:"message"`
	RepoID    int64      `json:"repo_id"`
	RepoName  string     `json:"repo_name"`
	StartedAt int64      `json:"start"`
	EndedAt   int64      `json:"end"`
}

type TaskStatus int
