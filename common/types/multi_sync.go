package types

import (
	"fmt"
	"strings"
	"time"
)

type SyncVersionResponse struct {
	Message string `json:"msg"`
	Data    struct {
		Versions []SyncVersion `json:"versions"`
		HasMore  bool          `json:"has_more"`
	}
}

type SyncVersionData struct {
	Versions []SyncVersion `json:"versions"`
	HasMore  bool          `json:"has_more"`
}

type SyncVersion struct {
	Version int64 `json:"version"`
	//0: opencsg, 1: HF
	SourceID       int64          `json:"source_id"`
	RepoPath       string         `json:"repo_path"`
	RepoType       RepositoryType `json:"repo_type"`
	LastModifyTime time.Time      `json:"last_modify_time"`
	ChangeLog      string         `json:"change_log"`
}

type ModelResponse struct {
	Message string `json:"msg"`
	Data    Model  `json:"data"`
}

type DatasetResponse struct {
	Message string  `json:"msg"`
	Data    Dataset `json:"data"`
}

const (
	SyncVersionSourceOpenCSG = iota
	SyncVersionSourceHF
)

func (s *SyncVersion) Prefix() string {
	var prefix string
	if s.SourceID == SyncVersionSourceOpenCSG {
		prefix = OpenCSGPrefix
	} else if s.SourceID == SyncVersionSourceHF {
		prefix = HuggingfacePrefix
	}
	return prefix
}

func (s *SyncVersion) BuildCloneURL(url, repoType, namespace, name string) string {
	namespace, _ = strings.CutPrefix(namespace, s.Prefix())
	return fmt.Sprintf("%s/%ss/%s/%s.git", url, repoType, namespace, name)
}

func (s *SyncVersion) BuildLocalVaule(originVaule string) string {
	return fmt.Sprintf("%s_%s", s.Prefix(), originVaule)
}
