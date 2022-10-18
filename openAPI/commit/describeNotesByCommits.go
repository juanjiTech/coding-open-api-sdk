package commit

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeNotesByCommitsReq struct {
	Action  string   `json:"Action"`  // DescribeNotesByCommits
	Commits []string `json:"Commits"` // commitSha1
	DepotID int64    `json:"DepotId"` // 7323
	NoteRef string   `json:"NoteRef"` // refs/notes/commits
}

type DescribeNotesByCommitsResp struct {
	Response struct {
		Notes []struct {
			CommitSha   string `json:"CommitSha"`   // commitSha1
			NoteContent string `json:"NoteContent"` // note 1
		} `json:"Notes"`
		RequestID string `json:"RequestId"` // c1c83461-dd49-3092-6ccd-4eecbabbd8aa
	} `json:"Response"`
}

// DescribeNotesByCommits 获取提交的note信息
func (c *Commit) DescribeNotesByCommits(req *DescribeNotesByCommitsReq) (resp DescribeNotesByCommitsResp, err error) {
	err = poster.Post(c.c, "DescribeNotesByCommits", req, &resp)
	return
}
