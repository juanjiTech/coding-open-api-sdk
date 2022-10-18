package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateGitCommitNoteReq struct {
	Action        string `json:"Action"`        // CreateGitCommitNote
	CommitMessage string `json:"CommitMessage"` // coding
	CommitSha     string `json:"CommitSha"`     // 5178d68bd3ee2a93733354884f0888584d671d7b
	DepotID       int64  `json:"DepotId"`       // 1383
	Note          string `json:"Note"`          // coding
	NotesRef      string `json:"NotesRef"`      //
}

type CreateGitCommitNoteResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // d7b00b5b-00b6-fb21-d6f5-92c146f73e0f
	} `json:"Response"`
}

// CreateGitCommitNote 创建提交注释
func (g *Git) CreateGitCommitNote(req *CreateGitCommitNoteReq) (resp CreateGitCommitNoteResp, err error) {
	err = poster.Post(g.c, "CreateGitCommitNote", req, &resp)
	return
}
