package cd

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteCdPipelineReq struct {
	Action       string `json:"Action"`       // DeleteCdPipeline
	Application  string `json:"Application"`  // test
	PipelineName string `json:"PipelineName"` // test
}

type DeleteCdPipelineResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 01d43d0b-ed71-ed55-4704-900a0802ad1d
	} `json:"Response"`
}

// DeleteCdPipeline 删除 CD 部署流程
func (c *Cd) DeleteCdPipeline(req *DeleteCdPipelineReq) (resp DeleteCdPipelineResp, err error) {
	err = poster.Post(c.c, "DeleteCdPipeline", req, &resp)
	return
}
