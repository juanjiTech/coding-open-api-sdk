package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyGitDepotUnarchiveReq struct {
	Action  string `json:"Action"`  // ModifyGitDepotUnarchive
	DepotID int64  `json:"DepotId"` // 2
}

type ModifyGitDepotUnarchiveResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 87b9311d-23fc-9f49-e16d-fa890f613ebb
	} `json:"Response"`
}

// ModifyGitDepotUnarchive 仓库解除归档
func (g *Git) ModifyGitDepotUnarchive(req *ModifyGitDepotUnarchiveReq) (resp ModifyGitDepotUnarchiveResp, err error) {
	err = poster.Post(g.c, "ModifyGitDepotUnarchive", req, &resp)
	return
}
