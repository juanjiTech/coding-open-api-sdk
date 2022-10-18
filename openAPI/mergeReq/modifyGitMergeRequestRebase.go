package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyGitMergeRequestRebaseReq struct {
	Action  string `json:"Action"`  // ModifyGitMergeRequestRebase
	DepotID int64  `json:"DepotId"` // 2
	MergeID int64  `json:"MergeId"` // 2
}

type ModifyGitMergeRequestRebaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // fd409d06-9765-c2d5-28e7-bee9cfca19c2
	} `json:"Response"`
}

// ModifyGitMergeRequestRebase 合并请求中的源分支进行rebase操作
func (m *MergeReq) ModifyGitMergeRequestRebase(req *ModifyGitMergeRequestRebaseReq) (resp ModifyGitMergeRequestRebaseResp, err error) {
	err = poster.Post(m.c, "ModifyGitMergeRequestRebase", req, &resp)
	return
}
