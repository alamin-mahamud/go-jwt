package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"local": "local/config.json",
	"dev": "dev/config.json",
	"staging": "staging/config.json",
	"deployment": "deployment/config.json",
	"production": "production/config.json",
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings = Settings{}
var env = "local"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting local environment due to lack of GO_ENV value")
		env = "local"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsLocalEnvironment() bool {
	return env == "local"
}
