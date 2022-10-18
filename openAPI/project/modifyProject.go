package project

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyProjectReq struct {
	Action      string `json:"Action"`      // ModifyProject
	Description string `json:"Description"` // project
	DisplayName string `json:"DisplayName"` // project
	EndDate     string `json:"EndDate"`     // 2020-01-02
	Name        string `json:"Name"`        // project
	ProjectID   int64  `json:"ProjectId"`   // 1
	StartDate   string `json:"StartDate"`   // 2020-01-01
}

type ModifyProjectResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// ModifyProject 编辑项目
func (p *Project) ModifyProject(req *ModifyProjectReq) (resp ModifyProjectResp, err error) {
	err = poster.Post(p.c, "ModifyProject", req, &resp)
	return
}
