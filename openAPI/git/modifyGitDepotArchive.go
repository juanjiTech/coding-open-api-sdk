package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyGitDepotArchiveReq struct {
	Action  string `json:"Action"`  // ModifyGitDepotArchive
	DepotID int64  `json:"DepotId"` // 2
}

type ModifyGitDepotArchiveResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 0f569eca-7575-e6fe-d02a-8d67d8c8b6be
	} `json:"Response"`
}

// ModifyGitDepotArchive 仓库归档
func (g *Git) ModifyGitDepotArchive(req *ModifyGitDepotArchiveReq) (resp ModifyGitDepotArchiveResp, err error) {
	err = poster.Post(g.c, "ModifyGitDepotArchive", req, &resp)
	return
}
