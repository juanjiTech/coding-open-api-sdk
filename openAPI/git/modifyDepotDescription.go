package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyDepotDescriptionReq struct {
	Action      string `json:"Action"`      // ModifyDepotDescription
	DepotID     int64  `json:"DepotId"`     // 791507
	Description string `json:"Description"` // It is a test depot
	UserID      int64  `json:"UserId"`      // 1
}

type ModifyDepotDescriptionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 4d7694ed-5818-3ea5-4834-a9d74cb2a594
	} `json:"Response"`
}

// ModifyDepotDescription 修改仓库描述
func (g *Git) ModifyDepotDescription(req *ModifyDepotDescriptionReq) (resp ModifyDepotDescriptionResp, err error) {
	err = poster.Post(g.c, "ModifyDepotDescription", req, &resp)
	return
}
