package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitTagsByShaReq struct {
	Action  string `json:"Action"`  // DescribeGitTagsBySha
	DepotID int64  `json:"DepotId"` // 2
	Sha     string `json:"Sha"`     // 575489c9aed1698d11f1ed398e3c4cb3d45d2ca9
}

type DescribeGitTagsByShaResp struct {
	Response struct {
		GitTags []struct {
			Message string `json:"Message"` //
			TagName string `json:"TagName"` // 2021.12.31
		} `json:"GitTags"`
		RequestID string `json:"RequestId"` // 8b048fe5-eb93-e8aa-a239-b864d9abd660
	} `json:"Response"`
}

// DescribeGitTagsBySha 查询含有某次提交的标签列表
func (t *Tag) DescribeGitTagsBySha(req *DescribeGitTagsByShaReq) (resp DescribeGitTagsByShaResp, err error) {
	err = poster.Post(t.c, "DescribeGitTagsBySha", req, &resp)
	return
}
