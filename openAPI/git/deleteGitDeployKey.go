package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteGitDeployKeyReq struct {
	Action      string `json:"Action"`      // DeleteGitDeployKey
	DeployKeyID int64  `json:"DeployKeyId"` // 1
	DepotID     int64  `json:"DepotId"`     // 1
}

type DeleteGitDeployKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 5649c350-f27a-a84a-3cff-f76c36f03280
	} `json:"Response"`
}

// DeleteGitDeployKey 删除部署公钥
func (g *Git) DeleteGitDeployKey(req *DeleteGitDeployKeyReq) (resp DeleteGitDeployKeyResp, err error) {
	err = poster.Post(g.c, "DeleteGitDeployKey", req, &resp)
	return
}
