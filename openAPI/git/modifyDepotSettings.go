package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyDepotSettingsReq struct {
	Action                   string `json:"Action"`                   // ModifyDepotSettings
	DepotPath                string `json:"DepotPath"`                // codingcorp/k8s/etcd
	DepotStatusCheckRequired bool   `json:"DepotStatusCheckRequired"` // true
}

type ModifyDepotSettingsResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // bce62f13-74f1-678d-b09b-24b1e6d8c867
	} `json:"Response"`
}

// ModifyDepotSettings 修改仓库设置
func (g *Git) ModifyDepotSettings(req *ModifyDepotSettingsReq) (resp ModifyDepotSettingsResp, err error) {
	err = poster.Post(g.c, "ModifyDepotSettings", req, &resp)
	return
}
