package wiki

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyWikiOrderReq struct {
	Action        string `json:"Action"`        // ModifyWikiOrder
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	Iid           int64  `json:"Iid"`           // 1
	ProjectName   string `json:"ProjectName"`   // demo-project
}

type ModifyWikiOrderResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyWikiOrder 修改父级顺序
func (w *Wiki) ModifyWikiOrder(req *ModifyWikiOrderReq) (resp ModifyWikiOrderResp, err error) {
	err = poster.Post(w.c, "ModifyWikiOrder", req, &resp)
	return
}
