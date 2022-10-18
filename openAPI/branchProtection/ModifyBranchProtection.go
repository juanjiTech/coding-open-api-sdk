package branchProtection

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyBranchProtectionReq struct {
	Action           string `json:"Action"` // ModifyBranchProtection
	BranchProtection struct {
		BranchProtectionID     int64  `json:"BranchProtectionId"`     // 28250
		DenyForcePush          bool   `json:"DenyForcePush"`          // true
		ForceSquash            bool   `json:"ForceSquash"`            // false
		MatcherCount           int64  `json:"MatcherCount"`           // 0
		RequiredCodeOwnerGrant bool   `json:"RequiredCodeOwnerGrant"` // false
		RequiredGrantCount     int64  `json:"RequiredGrantCount"`     // 1
		RequiredStatusCheck    bool   `json:"RequiredStatusCheck"`    // false
		Rule                   string `json:"Rule"`                   // dev/*
		SrcMustMergedTo        string `json:"SrcMustMergedTo"`        //
	} `json:"BranchProtection"`
	DepotID int64 `json:"DepotId"` // 809883
}

type ModifyBranchProtectionResp struct {
	Response struct {
		BranchProtection struct {
			BranchProtectionID     int64  `json:"BranchProtectionId"`     // 28250
			DenyForcePush          bool   `json:"DenyForcePush"`          // true
			ForceSquash            bool   `json:"ForceSquash"`            // false
			MatcherCount           int64  `json:"MatcherCount"`           // 0
			RequiredCodeOwnerGrant bool   `json:"RequiredCodeOwnerGrant"` // false
			RequiredGrantCount     int64  `json:"RequiredGrantCount"`     // 1
			RequiredStatusCheck    bool   `json:"RequiredStatusCheck"`    // false
			Rule                   string `json:"Rule"`                   // dev/*
			SrcMustMergedTo        string `json:"SrcMustMergedTo"`        //
		} `json:"BranchProtection"`
		RequestID string `json:"RequestId"` // 93bd76b6-68aa-af54-2607-b637e53edb74
	} `json:"Response"`
}

// ModifyBranchProtection 修改保护分支规则
func (b *BranchProtection) ModifyBranchProtection(req *ModifyBranchProtectionReq) (resp ModifyBranchProtectionResp, err error) {
	err = poster.Post(b.c, "ModifyBranchProtection", req, &resp)
	return
}
