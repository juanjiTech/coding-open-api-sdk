package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteMemberSshKeyReq struct {
	Action       string `json:"Action"`       // DeleteMemberSshKey
	MemberUserID int64  `json:"MemberUserId"` // 2
	SshKeyID     int64  `json:"SshKeyId"`     // 2
}

type DeleteMemberSshKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 13c2cc53-88fe-ad64-6253-865d7ccb69ed
	} `json:"Response"`
}

// DeleteMemberSshKey 删除团队成员的 SSH 公钥
func (g *Git) DeleteMemberSshKey(req *DeleteMemberSshKeyReq) (resp DeleteMemberSshKeyResp, err error) {
	err = poster.Post(g.c, "DeleteMemberSshKey", req, &resp)
	return
}
