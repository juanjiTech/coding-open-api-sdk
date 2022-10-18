package cd

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCdApplicationReq struct {
	Action      string `json:"Action"`      // DescribeCdApplication
	Application string `json:"Application"` // test
}

type DescribeCdApplicationResp struct {
	Response struct {
		Data struct {
			ApplicationJsonContent string `json:"ApplicationJsonContent"` // {"cloudProviders":"kubernetes,hostserver","description":"","instancePort":80,"lastModifyNickName":"coding","name":"test","trafficGuards":[],"user":"coding"}
		} `json:"Data"`
		RequestID string `json:"RequestId"` // a95b0fcf-8e96-4a2c-5052-01184fa4caff
	} `json:"Response"`
}

// DescribeCdApplication 获取 CD 应用详情
func (c *Cd) DescribeCdApplication(req *DescribeCdApplicationReq) (resp DescribeCdApplicationResp, err error) {
	err = poster.Post(c.c, "DescribeCdApplication", req, &resp)
	return
}
