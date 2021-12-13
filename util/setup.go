package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	u "github.com/Truth1984/awadau-go"
)

var config = defaultConfig()

// configPath default to ""
// configMap : {"loglevel": 30, "server": false, "port": 3000}
func Setup(configPath string, configMap map[string]interface{}) {

	confFromFile, errFromFile := parseConfigFromFile(configPath)
	confFromEnv := parseConfigFromEnv()

	config = u.MapMerge(defaultConfig(), configMap, confFromFile, confFromEnv)

	setupLogger((config["logging"]).(ConfigLogger))

	if errFromFile != nil {
		Debug(errFromFile)
	}

	Debug("config loaded")
}

func defaultConfig() map[string]interface{} {
	return map[string]interface{}{
		"logging": ConfigLogger{Level: 30},
		"server":  false,
		"port":    3000,
	}
}

func parseConfigFromFile(path string) (map[string]interface{}, error) {
	if path == "" {
		return u.Map(), errors.New("config file not found, using empty config")
	}

	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(path)

			confStr, _ := u.JsonToString(config)
			err := ioutil.WriteFile(path, []byte(confStr), 0644)
			if err != nil {
				panic(err)
			}
			return u.Map(), errors.New("config file created, on path: " + path)
		} else {
			panic(err)
		}
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	result, err := u.StringToJson(string(body))
	if err != nil {
		panic(err)
	}
	return result, nil
}

func parseConfigFromEnv() map[string]interface{} {
	keys := u.MapKeys(defaultConfig())
	return EnvGetExist(keys...)
}

func ConfigGet(key string) interface{} {
	return config[key]
}

func ConfigSet(key string, value interface{}) {
	config[key] = value
}

func setupLogger(conf ConfigLogger) {
	Logger = SetLogger(conf)
}
