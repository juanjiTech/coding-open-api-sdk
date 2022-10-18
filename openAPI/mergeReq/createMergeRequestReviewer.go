package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateMergeRequestReviewerReq struct {
	Action            string `json:"Action"`            // CreateMergeRequestReviewer
	DepotID           int64  `json:"DepotId"`           // 1
	MergeID           int64  `json:"MergeId"`           // 1
	ReviewerGlobalKey string `json:"ReviewerGlobalKey"` // coding-coding
}

type CreateMergeRequestReviewerResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 29ec9529-e03a-4a89-b71d-d988676e608e
		Reviewers []struct {
			Avatar    string `json:"Avatar"`    // /static/fruit_avatar/Fruit-20.png
			Email     string `json:"Email"`     //
			GlobalKey string `json:"GlobalKey"` // coding-coding
			ID        int64  `json:"Id"`        // 2
			Name      string `json:"Name"`      // coding
			Status    string `json:"Status"`    // INACTIVE
			TeamID    int64  `json:"TeamId"`    // 0
		} `json:"Reviewers"`
	} `json:"Response"`
}

// CreateMergeRequestReviewer 新增合并请求评论者
func (m *MergeReq) CreateMergeRequestReviewer(req *CreateMergeRequestReviewerReq) (resp CreateMergeRequestReviewerResp, err error) {
	err = poster.Post(m.c, "CreateMergeRequestReviewer", req, &resp)
	return
}
