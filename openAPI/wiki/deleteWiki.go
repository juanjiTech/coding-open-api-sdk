package wiki

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteWikiReq struct {
	Action        string `json:"Action"`        // DeleteWiki
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	Iid           int64  `json:"Iid"`           // 1
	ProjectName   string `json:"ProjectName"`   // demo-project
}

type DeleteWikiResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteWiki 删除 Wiki
func (w *Wiki) DeleteWiki(req *DeleteWikiReq) (resp DeleteWikiResp, err error) {
	err = poster.Post(w.c, "DeleteWiki", req, &resp)
	return
}
