package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteTeamLevelDepotSpecReq struct {
	Action        string `json:"Action"`        // DeleteTeamLevelDepotSpec
	DepotSpecName string `json:"DepotSpecName"` // depot-spec-name
}

type DeleteTeamLevelDepotSpecResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 00b6cddb-bd7d-964e-6885-45bbc387bd64
	} `json:"Response"`
}

// DeleteTeamLevelDepotSpec 删除团队级别的分支规范
func (g *Git) DeleteTeamLevelDepotSpec(req *DeleteTeamLevelDepotSpecReq) (resp DeleteTeamLevelDepotSpecResp, err error) {
	err = poster.Post(g.c, "DeleteTeamLevelDepotSpec", req, &resp)
	return
}
