package utils

import (
	"testing"
)

func TestConf(t *testing.T) {
	conf, err := initConf()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(conf.Weibo.Username)
		t.Logf("%v", conf)
		t.Log(conf.Dsn())
	}
}
