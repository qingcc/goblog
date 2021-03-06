package databases

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/go-ini/ini"
	"strconv"
)

//
func SessionStore() sessions.Store {
	//读取配置文件
	config, err := ini.Load("config/redis.ini")
	if err != nil {
		fmt.Println(err.Error())
	}
	key, _ := config.Section("session").GetKey("host")
	host = key.Value()
	key, _ = config.Section("session").GetKey("port")
	port = key.Value()
	key, _ = config.Section("session").GetKey("password")
	password = key.Value()
	key, _ = config.Section("session").GetKey("dbname")
	dbname, _ = strconv.Atoi(key.Value())
	key, _ = config.Section("session").GetKey("maxidle")
	maxidle, _ = strconv.Atoi(key.Value())
	key, _ = config.Section("session").GetKey("maxactive")
	maxactive, _ = strconv.Atoi(key.Value())
	store, _ := sessions.NewRedisStore(dbname, "tcp", host+":"+port, password, []byte("secret"))
	option := sessions.Options{
		MaxAge: 60 * 60 * 2, //2h
		Path:   "/",
	}
	store.Options(option)
	return store
}
