package release

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteGitReleaseReq struct {
	Action  string `json:"Action"`  // DeleteGitRelease
	DepotID int64  `json:"DepotId"` // 809883
	TagName string `json:"TagName"` // TagName
}

type DeleteGitReleaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 542449d7-c975-9fbe-9389-7d009e742d0f
	} `json:"Response"`
}

// DeleteGitRelease 删除仓库的版本息
func (r *Release) DeleteGitRelease(req *DeleteGitReleaseReq) (resp DeleteGitReleaseResp, err error) {
	err = poster.Post(r.c, "DeleteGitRelease", req, &resp)
	return
}
