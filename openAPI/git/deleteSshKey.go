package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteSshKeyReq struct {
	Action   string `json:"Action"`   // DeleteSshKey
	SshKeyID int64  `json:"SshKeyId"` // 15
}

type DeleteSshKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ef532935-ca07-87c8-3ebd-9d89b7aae330
	} `json:"Response"`
}

// DeleteSshKey 删除当前用户的 SSH 公钥
func (g *Git) DeleteSshKey(req *DeleteSshKeyReq) (resp DeleteSshKeyResp, err error) {
	err = poster.Post(g.c, "DeleteSshKey", req, &resp)
	return
}
