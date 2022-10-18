package requirement

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateRequirementDefectRelationReq struct {
	DefectCode      int64  `json:"DefectCode"`      // 2
	ProjectName     string `json:"ProjectName"`     // project-demo
	RequirementCode int64  `json:"RequirementCode"` // 1
}

type CreateRequirementDefectRelationResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateRequirementDefectRelation 需求关联缺陷
func (r *Requirement) CreateRequirementDefectRelation(req *CreateRequirementDefectRelationReq) (resp CreateRequirementDefectRelationResp, err error) {
	err = poster.Post(r.c, "CreateRequirementDefectRelation", req, &resp)
	return
}
