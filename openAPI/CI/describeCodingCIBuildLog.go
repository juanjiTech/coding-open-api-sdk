package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCodingCIBuildLogReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildLog
	BuildID int64  `json:"BuildId"` // 1
	StageID int64  `json:"StageId"` // 1
	StepID  int64  `json:"StepId"`  // 1
	Start   int64  `json:"start"`   // 1
}

type DescribeCodingCIBuildLogResp struct {
Response struct {
Data struct {
Log string `json:"Log"` // Started by user coding
Running in Durability level: MAX_SURVIVABILITY
[Pipeline] Start of Pipeline
[Pipeline] node
Running on Jenkins in /var/jenkins_home/workspace/cci-******-605077
[Pipeline] {
[Pipeline] stage
[Pipeline] { (检出)
[Pipeline] checkout
using credential 9378a66a-b949-4c4a-bf7e-094d0d7e36c0
Cloning the remote Git repository
Cloning repository git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git
 > git init /var/jenkins_home/workspace/cci-******-605077 # timeout=10
Fetching upstream changes from git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git
 > git --version # timeout=10
using GIT_SSH to set credentials 
 > git fetch --tags --progress -- git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git +refs/heads/*:refs/remotes/origin/*
 > git config remote.origin.url git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git # timeout=10
 > git config --add remote.origin.fetch +refs/heads/*:refs/remotes/origin/* # timeout=10
 > git config remote.origin.url git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git # timeout=10
Fetching upstream changes from git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git
using GIT_SSH to set credentials 
 > git fetch --tags --progress -- git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git +refs/heads/*:refs/remotes/origin/* +refs/merge/*:refs/remotes/origin/merge/*
 > git rev-parse 14c8e6e51ea01fc916dbcaf416f79a68f4******c7634^{commit} # timeout=10
Checking out Revision 14c8e6e51ea01fc916dbcaf416f79a68f4******c7634 (detached)
 > git config core.sparsecheckout # timeout=10
 > git checkout -f 14c8e6e51ea01fc916dbcaf416f79a68f4******c7634
Commit message: "Initial commit"
First time build. Skipping changelog.
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (构建)
[Pipeline] echo
构建中...
[Pipeline] sh
+ docker version
Client: Docker Engine - Community
 Version:           19.03.5
 API version:       1.39 (downgraded from 1.40)
 Go version:        go1.1******.1******
 Git commit:        633a0ea838
 Built:             Wed Nov 13 07:******:05 ******019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          18.09.7
  API version:      1.39 (minimum version 1.1******)
  Go version:       go1.10.8
  Git commit:       ******d0083d
  Built:            Thu Jun ******7 17:******6:******8 ******019
  OS/Arch:          linux/amd64
  Experimental:     false
[Pipeline] echo
构建完成.
[Pipeline] script
[Pipeline] {
[Pipeline] fileExists
[Pipeline] }
[Pipeline] // script
[Pipeline] archiveArtifacts
Archiving artifacts
Recording fingerprints
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (测试)
[Pipeline] echo
单元测试中...
[Pipeline] echo
单元测试完成.
[Pipeline] writeFile
[Pipeline] junit
Recording test results
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (部署)
[Pipeline] echo
部署中...
[Pipeline] echo
部署完成
[Pipeline] }
[Pipeline] // stage
[Pipeline] }
[Pipeline] // node
[Pipeline] End of Pipeline
Finished: SUCCESS

MoreData bool `json:"MoreData"` // false
TextDelivered int64 `json:"TextDelivered"` // 3003
TextSize int64 `json:"TextSize"` // 3003
} `json:"Data"`
RequestID int64 `json:"RequestId,string"` // 1
} `json:"Response"`
}

// DescribeCodingCIBuildLog 获取构建日志
func (c *Ci) DescribeCodingCIBuildLog(req *DescribeCodingCIBuildLogReq) (resp DescribeCodingCIBuildLogResp, err error) {
	err = poster.Post(c.c, "DescribeCodingCIBuildLog", req, &resp)
	return
}
