package commit

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateGitFilesReq struct {
	Action   string `json:"Action"`  // CreateGitFiles
	DepotID  int64  `json:"DepotId"` // 809883
	GitFiles []struct {
		Content string `json:"Content"` // line1
		line2
		line3
		line4
		line5
		Path string `json:"Path"` // src/main/hello.java
	} `json:"GitFiles"`
	LastCommitSha string `json:"LastCommitSha"` // cd3e7461b4061346a4803d1494623fc7ffa7f96a
	Message       string `json:"Message"`       // Create a new file named hello111.java
	NewRef        string `json:"NewRef"`        //
	Ref           string `json:"Ref"`           // master
}

type CreateGitFilesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b974da85-10bf-1411-0322-f84ccde0008c
	} `json:"Response"`
}

// CreateGitFiles 新建文件并提交
func (c *Commit) CreateGitFiles(req *CreateGitFilesReq) (resp CreateGitFilesResp, err error) {
	err = poster.Post(c.c, "CreateGitFiles", req, &resp)
	return
}
