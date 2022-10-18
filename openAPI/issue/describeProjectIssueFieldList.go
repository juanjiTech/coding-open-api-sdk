package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeProjectIssueFieldListReq struct {
	Action      string `json:"Action"`      // DescribeProjectIssueFieldList
	IssueType   string `json:"IssueType"`   // REQUIREMENT
	ProjectName string `json:"ProjectName"` // demo-project
}

type DescribeProjectIssueFieldListResp struct {
	Response struct {
		ProjectIssueFieldList []struct {
			CreatedAt  int64 `json:"CreatedAt"` // 1.5972834e+12
			IssueField struct {
				ComponentType string        `json:"ComponentType"` // SELECT_MEMBER_SINGLE
				CreatedAt     int64         `json:"CreatedAt"`     // 1.597283395e+12
				CreatedBy     int64         `json:"CreatedBy"`     // 0
				Deletable     bool          `json:"Deletable"`     // false
				Description   string        `json:"Description"`   //
				Editable      bool          `json:"Editable"`      // false
				IconURL       string        `json:"IconUrl"`       //
				ID            int64         `json:"Id"`            // 1
				Name          string        `json:"Name"`          // 处理人
				Options       []interface{} `json:"Options"`       // <nil>
				Required      bool          `json:"Required"`      // false
				Selectable    bool          `json:"Selectable"`    // false
				Sortable      bool          `json:"Sortable"`      // true
				TeamID        int64         `json:"TeamId"`        // 1
				Type          string        `json:"Type"`          // ASSIGNEE
				Unit          string        `json:"Unit"`          //
				UpdatedAt     int64         `json:"UpdatedAt"`     // 1.597283395e+12
			} `json:"IssueField"`
			IssueFieldID int64  `json:"IssueFieldId"` // 1
			IssueType    string `json:"IssueType"`    // REQUIREMENT
			NeedDefault  bool   `json:"NeedDefault"`  // false
			Required     bool   `json:"Required"`     // false
			UpdatedAt    int64  `json:"UpdatedAt"`    // 1.5972834e+12
			ValueString  string `json:"ValueString"`  //
		} `json:"ProjectIssueFieldList"`
		RequestID string `json:"RequestId"` // b2f77ca4-a34c-efde-29aa-389940fc8673
	} `json:"Response"`
}

// DescribeProjectIssueFieldList 查询属性设置
func (i *Issue) DescribeProjectIssueFieldList(req *DescribeProjectIssueFieldListReq) (resp DescribeProjectIssueFieldListResp, err error) {
	err = poster.Post(i.c, "DescribeProjectIssueFieldList", req, &resp)
	return
}
