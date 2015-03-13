package config
import (
    "github.com/BurntSushi/toml"
    "log"
)

type Config struct {
    User string
    Userpwd string
    Admin string
    Adminpwd string
    AuthSecret string
    SessionExpirationTimeMinutes int64
    SessionCookieName string

}

var Values Config
const confpath = "config/config.toml"

func InitConfig() {


    if _, err := toml.DecodeFile(confpath, &Values); err != nil {
        log.Fatal("Failed to load config file '" + confpath + "' with error: ", err)
    }

    log.Println("Successfully loaded config file '" + confpath + "'")
}