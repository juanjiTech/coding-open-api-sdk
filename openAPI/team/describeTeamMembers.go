package team

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeTeamMembersReq struct {
	Action     string `json:"Action"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
}

type DescribeTeamMembersResp struct {
	Response struct {
		RequestId string `json:"RequestId"`
		Data      struct {
			PageNumber  int `json:"PageNumber"`
			PageSize    int `json:"PageSize"`
			TotalCount  int `json:"TotalCount"`
			TeamMembers []struct {
				Id              int    `json:"Id"`
				TeamId          int    `json:"TeamId"`
				Name            string `json:"Name"`
				NamePinYin      string `json:"NamePinYin"`
				Avatar          string `json:"Avatar"`
				Email           string `json:"Email"`
				Phone           string `json:"Phone"`
				EmailValidation int    `json:"EmailValidation"`
				PhoneValidation int    `json:"PhoneValidation"`
				Status          int    `json:"Status"`
			} `json:"TeamMembers"`
		} `json:"Data"`
	} `json:"Response"`
}

func (t *Team) DescribeTeamMembers(req *DescribeTeamMembersReq) (resp DescribeCodingProjectsResp, err error) {
	err = poster.Post(t.c, req.Action, req, &resp)
	return
}
