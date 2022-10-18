package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitBlobRawReq struct {
	Action  string `json:"Action"`  // DescribeGitBlobRaw
	BlobSha string `json:"BlobSha"` // a7954ee405100719a59c1e03f4ac98ce92a2a257
	DepotID int64  `json:"DepotId"` // 2
}

type DescribeGitBlobRawResp struct {
Response struct {
Content string `json:"Content"` // name: extended_text
description: Extended official text to build special text like inline image or @somebody quickly,it also support custom background,custom over flow and custom selection toolbar and handles.
version: 5.0.4
homepage: https://github.com/fluttercandies/extended_text

environment:
  sdk: '>=2.12.0 <3.0.0'
  flutter: ">=2.0.0"

dependencies:
  extended_text_library: ^5.0.0
  flutter:
    sdk: flutter

dev_dependencies:
  flutter_test:
    sdk: flutter
RequestID string `json:"RequestId"` // 5649c350-f27a-a84a-3cff-f76c36f03280
} `json:"Response"`
}

// DescribeGitBlobRaw 查询 Git Blob raw 信息
func (g *Git) DescribeGitBlobRaw(req *DescribeGitBlobRawReq) (resp DescribeGitBlobRawResp, err error) {
	err = poster.Post(g.c, "DescribeGitBlobRaw", req, &resp)
	return
}
