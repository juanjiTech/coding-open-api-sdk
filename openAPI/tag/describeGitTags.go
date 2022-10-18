package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitTagsReq struct {
	Action  string `json:"Action"`  // DescribeGitTags
	DepotID int64  `json:"DepotId"` // 1001
}

type DescribeGitTagsResp struct {
	Response struct {
		GitTags []struct {
			Message string `json:"Message"` // this is tag-demo1
			TagName string `json:"TagName"` // tag-demo1
		} `json:"GitTags"`
		RequestID string `json:"RequestId"` // 05d4f18d-8602-ad9c-f6a2-8ad6f8f73127
	} `json:"Response"`
}

// DescribeGitTags 查询仓库所有标签
func (t *Tag) DescribeGitTags(req *DescribeGitTagsReq) (resp DescribeGitTagsResp, err error) {
	err = poster.Post(t.c, "DescribeGitTags", req, &resp)
	return
}
