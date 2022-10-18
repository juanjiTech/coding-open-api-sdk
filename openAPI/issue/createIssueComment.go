package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateIssueCommentReq struct {
	Action      string `json:"Action"`      // CreateIssueComment
	Content     string `json:"Content"`     // 这个要配合测试
	IssueCode   int64  `json:"IssueCode"`   // 1
	ParentID    int64  `json:"ParentId"`    // 0
	ProjectName string `json:"ProjectName"` // TestProjectDemo
}

// CreateIssueComment 创建事项评论
func (i *Issue) CreateIssueComment(req *CreateIssueCommentReq) error {
	return poster.Post(i.c, "CreateIssueComment", req, nil)
}
