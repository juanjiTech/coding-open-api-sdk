package requirement

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyDefectRelatedRequirementReq struct {
	DefectCode      int64  `json:"DefectCode"`      // 2
	ProjectName     string `json:"ProjectName"`     // project-demo
	RequirementCode int64  `json:"RequirementCode"` // 3
}

type ModifyDefectRelatedRequirementResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyDefectRelatedRequirement 修改缺陷所属的需求
func (r *Requirement) ModifyDefectRelatedRequirement(req *ModifyDefectRelatedRequirementReq) (resp ModifyDefectRelatedRequirementResp, err error) {
	err = poster.Post(r.c, "ModifyDefectRelatedRequirement", req, &resp)
	return
}
