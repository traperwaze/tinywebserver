package config

import (
	"fmt"
	"reflect"

	"github.com/tkanos/gonfig"
)

type configuration struct {
	HOST           string
	MYSQL_HOST     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
}

var Config *configuration

func Init() {
	Config = &configuration{}
	err := gonfig.GetConf("config.json", Config)
	if err != nil {
		panic("Config file not found")
	}
}

func Get(key string) string {
	if hasField(Config, key) {
		rc := reflect.ValueOf(Config)

		if rc.Kind() == reflect.Ptr {
			rc = rc.Elem()
		}

		return rc.FieldByName(key).String()
	}

	return ""
}

func GetMysqlConnectionString() string {
	f := fmt.Sprintf("%s:%s@(%s)/%s", Config.MYSQL_USER, Config.MYSQL_PASSWORD, Config.MYSQL_HOST, Config.MYSQL_DATABASE)

	return f
}

func hasField(c *configuration, key string) bool {
	rc := reflect.ValueOf(c)

	if rc.Kind() == reflect.Ptr {
		rc = rc.Elem()
	}

	if rc.Kind() != reflect.Struct {
		return false
	}

	ri := reflect.Indirect(rc)

	return ri.FieldByName(key).IsValid()
}
