package config

import (
	"github.com/BurntSushi/toml"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/service"
	"github.com/rickshawdriver/somebody/pkg/system"
	"github.com/rickshawdriver/somebody/store"
	"os"
)

var (
	defaultConfigPath = []string{
		"../somebody.toml",
		"somebody.toml",
		"/etc/rickshawdriver/somebody.toml",
	}
)

type FilePath struct {
	OriginalPath    string `toml:"originalPath"`
	PidFileLocation string `toml:"pidFileLocation"`
}

type Config struct {
	FilePath     FilePath            `toml:"filepath"`
	Log          log.Logger          `toml:"log"`
	RouterDegree int                 `toml:"routerDegree"`
	DnsCacheConf system.DnsCacheConf `toml:"dnsCache"`
	Store        store.StoreConf     `toml:"store"`
	HttpConf     service.HttpConf    `toml:"httpConf"`
}

// load my config file
func Load(config *Config) error {
	for _, path := range defaultConfigPath {
		_, err := os.Open(path)
		if err == nil { // success
			config.FilePath.OriginalPath = path
			break
		}
		if os.IsNotExist(err) {
			continue
		}

		return err
	}

	// if file not exist
	if config.FilePath.OriginalPath == "" && nil != createDefaultFile() {
		//create default file
	}

	// analyse config file
	if _, err := toml.DecodeFile(config.FilePath.OriginalPath, &config); err != nil {
		return err
	}

	return nil
}

func createDefaultFile() error {
	f, err := os.Create("somebody.toml")
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(``))

	return nil
}
