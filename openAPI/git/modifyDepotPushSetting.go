package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyDepotPushSettingReq struct {
	Action    string `json:"Action"`    // ModifyDepotPushSetting
	DepotPath string `json:"DepotPath"` // codingcorp/project-d/depot-2
	Param     struct {
		CheckCommitAuthor           bool   `json:"CheckCommitAuthor"`           // false
		CommitMessageMustMatchRegex string `json:"CommitMessageMustMatchRegex"` // ^(feat|fix)((.+))?
		DenyForcePush               bool   `json:"DenyForcePush"`               // true
		PushDenyFile                string `json:"PushDenyFile"`                //
		PushFileSize                string `json:"PushFileSize"`                //
	} `json:"Param"`
}

type ModifyDepotPushSettingResp struct {
	Response struct {
		Data struct {
			CheckCommitAuthor           bool   `json:"CheckCommitAuthor"`           // false
			CommitMessageMustMatchRegex string `json:"CommitMessageMustMatchRegex"` // ^(feat|fix)((.+))?
			DenyForcePush               bool   `json:"DenyForcePush"`               // true
			PushDenyFile                string `json:"PushDenyFile"`                //
			PushFileSize                string `json:"PushFileSize"`                //
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 1c06de89-1345-4188-aeb3-8929c826fa58
	} `json:"Response"`
}

// ModifyDepotPushSetting 修改仓库推送设置
func (g *Git) ModifyDepotPushSetting(req *ModifyDepotPushSettingReq) (resp ModifyDepotPushSettingResp, err error) {
	err = poster.Post(g.c, "ModifyDepotPushSetting", req, &resp)
	return
}
