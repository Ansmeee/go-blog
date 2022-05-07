package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

var INI_FILE *ini.File

func Init() error {
	args := os.Args

	if len(args) == 2 {
		config, err := ini.Load(args[1])
		INI_FILE = config
		return err
	}


	config, err := ini.Load("config.ini")
	INI_FILE = config
	return err
}

func Get(name string) string {
	cfgSlice := strings.Split(name, ".")
	if cfgSlice[0] == "" {
		log.Println("invalid config section")
		return ""
	}

	if cfgSlice[1] == "" {
		log.Println("invalid config name")
		return ""
	}

	return INI_FILE.Section(cfgSlice[0]).Key(cfgSlice[1]).String()
}
