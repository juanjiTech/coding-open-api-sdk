package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyIssueCommentReq struct {
	Action      string `json:"Action"`      // CreateIssueComment
	CommentID   int64  `json:"CommentId"`   // 1
	Content     string `json:"Content"`     // 这个要配合测试
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // TestProjectDemo
}

// ModifyIssueComment 修改事项评论
func (i *Issue) ModifyIssueComment(req *ModifyIssueCommentReq) error {
	return poster.Post(i.c, "ModifyIssueComment", req, nil)
}
