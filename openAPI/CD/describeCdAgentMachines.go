package cd

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCdAgentMachinesReq struct {
	Action     string `json:"Action"`     // DescribeCdAgentMachines
	Keyword    string `json:"Keyword"`    //
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
}

type DescribeCdAgentMachinesResp struct {
	Response struct {
		Data struct {
			AgentMachines []struct {
				ID     int64  `json:"Id"`     // 10
				Name   string `json:"Name"`   // master
				Status string `json:"Status"` // SUCCESS
			} `json:"AgentMachines"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalPage  int64 `json:"TotalPage"`  // 1
			TotalRow   int64 `json:"TotalRow"`   // 1
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 5a3d0128-36a3-cce3-2938-7f7e2f68b5c3
	} `json:"Response"`
}

// DescribeCdAgentMachines 获取 CD 堡垒机列表
func (c *Cd) DescribeCdAgentMachines(req *DescribeCdAgentMachinesReq) (resp DescribeCdAgentMachinesResp, err error) {
	err = poster.Post(c.c, "DescribeCdAgentMachines", req, &resp)
	return
}
