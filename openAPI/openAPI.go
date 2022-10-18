package openAPI

import (
	"github.com/juanjiTech/coding-open-api-sdk/config"
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/artifact"         // 制品仓库
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/branch"           // 分支
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/branchProtection" // 保护分支规则
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/cd"               // 持续部署
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/ci"               // 持续集成
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/commit"           // 提交
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/git"              // 代码托管
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/issue"            // 事项
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/iteration"        // 迭代
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/mergeReq"         // 合并请求
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/project"          // 项目
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/projectSetting"   // 项目设置
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/protectedBranch"  // 保护分支
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/release"          // 版本
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/requirement"      // 需求
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/tag"              // 标签
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/team"             // 团队
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/test"             // 测试管理
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/wiki"             // 文档管理
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/work"             // 工时
)

type OpenAPI struct {
	c                *config.Basic
	Team             team.Team                         // 团队
	Project          project.Project                   // 项目
	ProjectSetting   projectSetting.ProjectSetting     // 项目设置
	Requirement      requirement.Requirement           // 需求
	Issue            issue.Issue                       // 事项
	Iteration        iteration.Iteration               // 迭代
	Work             work.Work                         // 工时
	Git              git.Git                           // 代码托管
	Commit           commit.Commit                     // 提交
	Tag              tag.Tag                           // 标签
	Release          release.Release                   // 版本
	Branch           branch.Branch                     // 分支
	ProtectedBranch  protectedBranch.ProtectedBranch   // 保护分支
	BranchProtection branchProtection.BranchProtection // 保护分支规则
	MergeReq         mergeReq.MergeReq                 // 合并请求
	Ci               ci.Ci                             // 持续集成
	Cd               cd.Cd                             // 持续部署
	Artifact         artifact.Artifact                 // 制品仓库
	Test             test.Test                         // 测试管理
	Wiki             wiki.Wiki                         // 文档管理
}

func New(c *config.Basic) OpenAPI {
	return OpenAPI{c,
		team.New(c),
		project.New(c),
		projectSetting.New(c),
		requirement.New(c),
		issue.New(c),
		iteration.New(c),
		work.New(c),
		git.New(c),
		commit.New(c),
		tag.New(c),
		release.New(c),
		branch.New(c),
		protectedBranch.New(c),
		branchProtection.New(c),
		mergeReq.New(c),
		ci.New(c),
		cd.New(c),
		artifact.New(c),
		test.New(c),
		wiki.New(c),
	}
}
