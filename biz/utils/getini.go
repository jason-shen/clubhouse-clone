package utils

import (
	"gopkg.in/ini.v1"
	"os"
)

func GetIni(section, key, defaultValue string) string  {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		Errorf("Fail to read file: ", err)
		os.Exit(1)
	}

	if value := cfg.Section(section).Key(key).String(); value !="" {
		return value
	}
	return defaultValue
}
