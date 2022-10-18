package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteGitTagReq struct {
	Action  string `json:"Action"`  // DeleteGitTag
	DepotID int64  `json:"DepotId"` // 1001
	TagName string `json:"TagName"` // tag-demo
}

type DeleteGitTagResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // aa9dc0b9-6c8c-ed2c-ac54-191b40a63e81
	} `json:"Response"`
}

// DeleteGitTag 删除标签
func (t *Tag) DeleteGitTag(req *DeleteGitTagReq) (resp DeleteGitTagResp, err error) {
	err = poster.Post(t.c, "DeleteGitTag", req, &resp)
	return
}
