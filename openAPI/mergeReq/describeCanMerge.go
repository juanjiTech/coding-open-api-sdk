package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCanMergeReq struct {
	Action  string `json:"Action"`  // DescribeCanMerge
	DepotID int64  `json:"DepotId"` // 809883
	Source  string `json:"Source"`  // dev/branch1
	Target  string `json:"Target"`  // release
}

type DescribeCanMergeResp struct {
	Response struct {
		MergeStatus string `json:"MergeStatus"` // MERGEABLE
		RequestID   string `json:"RequestId"`   // f4135d6c-9bc6-2dd6-97f4-b2a850dd9713
	} `json:"Response"`
}

// DescribeCanMerge 获取分支的合并状态
func (m *MergeReq) DescribeCanMerge(req *DescribeCanMergeReq) (resp DescribeCanMergeResp, err error) {
	err = poster.Post(m.c, "DescribeCanMerge", req, &resp)
	return
}
