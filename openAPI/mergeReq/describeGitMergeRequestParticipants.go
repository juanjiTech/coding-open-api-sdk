package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitMergeRequestParticipantsReq struct {
	Action  string `json:"Action"`  // DescribeGitMergeRequestParticipants
	DepotID int64  `json:"DepotId"` // 2
	MergeID int64  `json:"MergeId"` // 1
}

type DescribeGitMergeRequestParticipantsResp struct {
	Response struct {
		Participants []struct {
			Avatar    string `json:"Avatar"`    // /static/fruit_avatar/Fruit-20.png
			Email     string `json:"Email"`     // coding@coding.com
			GlobalKey string `json:"GlobalKey"` // coding-coding
			ID        int64  `json:"Id"`        // 2
			Name      string `json:"Name"`      // coding
			Status    string `json:"Status"`    // INACTIVE
			TeamID    int64  `json:"TeamId"`    // 1
		} `json:"Participants"`
		RequestID string `json:"RequestId"` // 33affd68-fb4f-d677-0a75-1c1abf78bd2a
	} `json:"Response"`
}

// DescribeGitMergeRequestParticipants 查询合并请求的参与者
func (m *MergeReq) DescribeGitMergeRequestParticipants(req *DescribeGitMergeRequestParticipantsReq) (resp DescribeGitMergeRequestParticipantsResp, err error) {
	err = poster.Post(m.c, "DescribeGitMergeRequestParticipants", req, &resp)
	return
}
