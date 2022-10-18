package project

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteOneProjectReq struct {
	Action    string `json:"Action"`    // DeleteOneProject
	ProjectID int64  `json:"ProjectId"` // 2
}

type DeleteOneProjectResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// DeleteOneProject 删除项目
func (p *Project) DeleteOneProject(req *DeleteOneProjectReq) (resp DeleteOneProjectResp, err error) {
	err = poster.Post(p.c, "DeleteOneProject", req, &resp)
	return
}
