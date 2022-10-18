package branchProtection

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeBranchProtectionsReq struct {
	Action  string `json:"Action"`  // DescribeBranchProtections
	DepotID int64  `json:"DepotId"` // 809883
}

type DescribeBranchProtectionsResp struct {
	Response struct {
		BranchProtections []struct {
			BranchProtectionID     int64  `json:"BranchProtectionId"`     // 28233
			DenyForcePush          bool   `json:"DenyForcePush"`          // true
			ForceSquash            bool   `json:"ForceSquash"`            // false
			MatcherCount           int64  `json:"MatcherCount"`           // 0
			RequiredCodeOwnerGrant bool   `json:"RequiredCodeOwnerGrant"` // false
			RequiredGrantCount     int64  `json:"RequiredGrantCount"`     // 1
			RequiredStatusCheck    bool   `json:"RequiredStatusCheck"`    // false
			Rule                   string `json:"Rule"`                   // master
			SrcMustMergedTo        string `json:"SrcMustMergedTo"`        //
		} `json:"BranchProtections"`
		RequestID string `json:"RequestId"` // 9fd71e2c-c7d1-5dfe-9667-cff5340a7941
	} `json:"Response"`
}

// DescribeBranchProtections 查询保护分支规则列表
func (b *BranchProtection) DescribeBranchProtections(req *DescribeBranchProtectionsReq) (resp DescribeBranchProtectionsResp, err error) {
	err = poster.Post(b.c, "DescribeBranchProtections", req, &resp)
	return
}
