package config

type Basic struct {
	// 二选一
	Bearer string // OAuth2
	Token  string // 个人令牌
}
