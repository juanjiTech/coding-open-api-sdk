package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitTagReq struct {
	Action  string `json:"Action"`  // DescribeGitTag
	DepotID int64  `json:"DepotId"` // 1001
	TagName string `json:"TagName"` // tag-demo
}

type DescribeGitTagResp struct {
	Response struct {
		GitTag struct {
			Message string `json:"Message"` // this is a tag-demo
			TagName string `json:"TagName"` // tag-demo
		} `json:"GitTag"`
		RequestID string `json:"RequestId"` // a3a852b3-3d5a-d87f-fe9a-7ea8af9e22ed
	} `json:"Response"`
}

// DescribeGitTag 查询指定标签
func (t *Tag) DescribeGitTag(req *DescribeGitTagReq) (resp DescribeGitTagResp, err error) {
	err = poster.Post(t.c, "DescribeGitTag", req, &resp)
	return
}
