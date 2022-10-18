package branch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitRefReq struct {
	Action   string `json:"Action"`   // DescribeGitRef
	DepotID  int64  `json:"DepotId"`  // 809883
	Revision string `json:"Revision"` // tag_one
}

type DescribeGitRefResp struct {
	Response struct {
		GitRef struct {
			AnnotatedTag bool   `json:"AnnotatedTag"` // true
			DisplayName  string `json:"DisplayName"`  // refs/tags/tag_one
			FullMessage  string `json:"FullMessage"`  //
			Name         string `json:"Name"`         // refs/tags/tag_one
			ObjectID     string `json:"ObjectId"`     // 8121d556dc3b30767d77a8e3fa814826976d2003
			RefObjectID  string `json:"RefObjectId"`  // 6efc794f0d694fed6cd033ddc984026887af258a
			ShortMessage string `json:"ShortMessage"` //
		} `json:"GitRef"`
		RequestID string `json:"RequestId"` // 772d0468-ae45-bd32-377a-8e2de57cf685
	} `json:"Response"`
}

// DescribeGitRef 查询分支或标签的信息
func (b *Branch) DescribeGitRef(req *DescribeGitRefReq) (resp DescribeGitRefResp, err error) {
	err = poster.Post(b.c, "DescribeGitRef", req, &resp)
	return
}
