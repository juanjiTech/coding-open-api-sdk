package cd

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type TriggerCdPipelineReq struct {
	Action             string `json:"Action"`             // TriggerCdPipeline
	Application        string `json:"Application"`        // test
	PipelineNameOrID   string `json:"PipelineNameOrId"`   // test
	TriggerJsonContent string `json:"TriggerJsonContent"` // {"buildNumber":"","type":"manual","dryRun":false,"user":"coding"}
}

type TriggerCdPipelineResp struct {
	Response struct {
		Data struct {
			PipelineExecutionID  string `json:"PipelineExecutionId"`  // 01EXNXJVMK6EVRHB49CH7FG2BH
			PipelineExecutionRef string `json:"PipelineExecutionRef"` // /pipelines/01EXNXJVMK6EVRHB49CH7FG2BH
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 2969c6e0-689c-479c-98f0-dce488795abe
	} `json:"Response"`
}

// TriggerCdPipeline 触发部署流程
func (c *Cd) TriggerCdPipeline(req *TriggerCdPipelineReq) (resp TriggerCdPipelineResp, err error) {
	err = poster.Post(c.c, "TriggerCdPipeline", req, &resp)
	return
}
