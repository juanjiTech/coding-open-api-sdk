package coding_open_api_sdk

import (
	"github.com/juanjiTech/coding-open-api-sdk/config"
	"github.com/juanjiTech/coding-open-api-sdk/openAPI"
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/team"
	"testing"
)

func TestFullInit(t *testing.T) {
	api := openAPI.New(&config.Basic{
		Token: "abc",
	})
	_, _ = api.Team.DescribeCodingProjects(1, 10)           // 未提供可选项
	_, _ = api.Team.DescribeCodingProjects(1, 10, "juanji") // 提供可选项
}

func TestTeamInit(t *testing.T) {
	teamAPI := team.New(&config.Basic{
		Token: "abc",
	})
	_, _ = teamAPI.DescribeCodingProjects(1, 10) // 也可以单独初始化部分
}
