package utils

import (
	"fmt"

	"github.com/BurntSushi/toml"
	logs "github.com/sirupsen/logrus"
)

type appconf struct {
	App struct {
		Mode   string `toml:"mode"`
		Addr   string `toml:"addr"`
		Srv    string `toml:"srv"`
		JwtKey string `toml:"jwt_key"`
		JwtExp int    `toml:"jwt_exp"`
	}
	Database struct {
		Host   string `toml:"host"`
		Port   int    `toml:"port"`
		User   string `toml:"user"`
		Passwd string `toml:"passwd"`
		Dbname string `toml:"dbname"`
		Params string `toml:"params"`
	} `toml:"database"`
	Xorm struct {
		Idle        int  `toml:"idle"`
		Open        int  `toml:"open"`
		Show        bool `toml:"show"`
		Sync        bool `toml:"sync"`
		CacheEnable bool `toml:"cache_enable"`
		CacheCount  int  `toml:"cache_count"`
	} `toml:"xorm"`
	Info
}

type Info struct {
	Website struct {
		Title   string `toml:"title"`
		Explain string `toml:"explain"`
		Email   string `toml:"email"`
	} `toml:"website"`
	Wechat struct {
		Appid  string `toml:"appid"`
		Secret string `toml:"secret"`
	} `toml:"wechat"`
	Weibo struct {
		Username   string `toml:"username"`
		ProfileURL string `toml:"profile_url"`
	}
}

func (conf *appconf) IsProd() bool {
	return conf.App.Mode == "prod"
}
func (conf *appconf) IsDev() bool {
	return conf.App.Mode == "dev"
}

const _dsn = "%s:%s@tcp(%s:%d)/%s?%s"

func (conf *appconf) Dsn() string {
	return fmt.Sprintf(_dsn, conf.Database.User, conf.Database.Passwd, conf.Database.Host, conf.Database.Port, conf.Database.Dbname, conf.Database.Params)
}

var (
	Conf          *appconf
	appConfig     = "./conf/app.toml"
	profileConfig = "./conf/profile.toml"
)

func Init() {
	var err error
	Conf, err = initConf()
	if err != nil {
		logs.Fatal("config init error : ", err.Error())
	}
	logs.Debug("conf init")
}

func initConf() (*appconf, error) {
	app := &appconf{}
	_, err := toml.DecodeFile(appConfig, app)
	if err != nil {
		return nil, err
	}

	profile := &Info{}
	_, err = toml.DecodeFile(profileConfig, profile)
	if err != nil {
		return nil, err
	}
	app.Info = *profile
	return app, nil
}
