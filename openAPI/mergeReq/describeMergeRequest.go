package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeMergeRequestReq struct {
	Action  string `json:"Action"`  // DescribeMergeRequest
	DepotID int64  `json:"DepotId"` // 46
	MergeID int64  `json:"MergeId"` // 5
}

type DescribeMergeRequestResp struct {
Response struct {
MergeRequestInfo struct {
Describe string `json:"Describe"` // ## 改动说明
[scheduler]仓库导入同步。开发人员：@洋葱猴`
## 测试要点

无
SourceBranch string `json:"SourceBranch"` // dev
Status string `json:"Status"` // ACCEPTED
TargetBranch string `json:"TargetBranch"` // master
Title string `json:"Title"` // 测试合并
} `json:"MergeRequestInfo"`
RequestID string `json:"RequestId"` // bcc8760c-5abf-ab67-f88c-f29befab2c8e
} `json:"Response"`
}

// DescribeMergeRequest 查询合并请求详情
func (m *MergeReq) DescribeMergeRequest(req *DescribeMergeRequestReq) (resp DescribeMergeRequestResp, err error) {
	err = poster.Post(m.c, "DescribeMergeRequest", req, &resp)
	return
}
