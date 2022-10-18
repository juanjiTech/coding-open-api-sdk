package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateAttachmentPrepareSignUrlReq struct {
	Action      string `json:"Action"`      // CreateAttachmentPrepareSignUrl
	FileName    string `json:"FileName"`    // xx
	ProjectName string `json:"ProjectName"` // project name
}

type CreateAttachmentPrepareSignUrlResp struct {
	Response struct {
		Data struct {
			AttachmentPrepare struct {
				AttachmentID   int64  `json:"AttachmentId"`   // 1750
				PrepareSignURL string `json:"PrepareSignUrl"` // http://www.coding.net/xxxx
			} `json:"AttachmentPrepare"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// CreateAttachmentPrepareSignUrl 生成附件预上传信息
func (t *Test) CreateAttachmentPrepareSignUrl(req *CreateAttachmentPrepareSignUrlReq) (resp CreateAttachmentPrepareSignUrlResp, err error) {
	err = poster.Post(t.c, "CreateAttachmentPrepareSignUrl", req, &resp)
	return
}
