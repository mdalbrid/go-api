package config

import (
	"fmt"
	configUtil "github.com/mdalbrid/utils/config"
	"github.com/mdalbrid/utils/server"
	"time"
)

var ServerConfig server.Config
var DbConfig *configUtil.DatabasePostgresConfig

func init() {
	DbConfig = configUtil.GetDatabasePostgresConfig()
	viper := configUtil.GetViper()
	ServerConfig = server.Config{
		Addr: fmt.Sprintf(`%s:%d`,
			viper.GetString("AWS.OBJECT_GROUP.SERVER_HOST"),
			viper.GetInt("AWS.OBJECT_GROUP.SERVER_PORT"),
		),
		WriteTimeout: time.Duration(viper.GetInt("AWS.OBJECT_GROUP.SERVER_WRITE_TIMEOUT")) * time.Second,
		ReadTimeout:  time.Duration(viper.GetInt("AWS.OBJECT_GROUP.SERVER_READ_TIMEOUT")) * time.Second,
	}
}
