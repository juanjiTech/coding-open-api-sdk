package commit

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyGitCommitStatusReq struct {
	Action      string `json:"Action"`      // ModifyGitCommitStatus
	CommitSha   string `json:"CommitSha"`   // da3e8c8f93516bc8e1609e5a91df98439c6c7b49
	Context     string `json:"Context"`     // 示例任务 #1
	DepotID     int64  `json:"DepotId"`     // 0
	DepotPath   string `json:"DepotPath"`   // codingcorp/project-d/depot-d
	Description string `json:"Description"` // 构建成功
	State       string `json:"State"`       // COMMIT_STATE_ERROR
	TargetURL   string `json:"TargetURL"`   // http://test.coding.net/p/project-name/ci/job/20684/build/1/pipeline
}

type ModifyGitCommitStatusResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 2cc95576-edd9-4b05-a133-886a712588a9
	} `json:"Response"`
}

// ModifyGitCommitStatus 修改提交对应的流水线状态
func (c *Commit) ModifyGitCommitStatus(req *ModifyGitCommitStatusReq) (resp ModifyGitCommitStatusResp, err error) {
	err = poster.Post(c.c, "ModifyGitCommitStatus", req, &resp)
	return
}
