package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateGitTagReq struct {
	Action     string `json:"Action"`     // CreateGitTag
	DepotID    int64  `json:"DepotId"`    // 1001
	StartPoint string `json:"StartPoint"` // master
	TagName    string `json:"TagName"`    // tag-demo
	Message    string `json:"message"`    // this is my create tag demo
}

type CreateGitTagResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b6445913-3b7c-9907-c6b3-a2d63abfd5c6
	} `json:"Response"`
}

// CreateGitTag 创建标签
func (t *Tag) CreateGitTag(req *CreateGitTagReq) (resp CreateGitTagResp, err error) {
	err = poster.Post(t.c, "CreateGitTag", req, &resp)
	return
}
