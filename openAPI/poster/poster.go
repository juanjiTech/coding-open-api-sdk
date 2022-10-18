package poster

import (
	"errors"
	"github.com/guonaihong/gout"
	"github.com/juanjiTech/coding-open-api-sdk/config"
)

func Post(config *config.Basic, action string, req interface{}, resp interface{}) error {
	var auth string
	if config.Token != "" {
		auth = "token " + config.Token
	} else if config.Bearer != "" {
		auth = "Bearer " + config.Bearer
	} else {
		return errors.New("token and bearer mustn't both be empty")
	}
	return gout.POST("https://e.coding.net/open-api?Action=" + action).
		SetHeader(gout.H{"Authorization": auth}).
		SetJSON(req).
		BindJSON(resp).
		Do()
}
