package main

import (
	"os"
	"testing"
)

const ConfigPath  ="./.git/config"
const InvalidConfigPath  ="./.git/invalid_config"
const SshUrl  ="git@github.com:mo7amed-3bdalla7/grurl.git"
const HttpUrl  ="https://github.com/mo7amed-3bdalla7/grurl.git"

func NewConfig() *RemoteConfig {
	return &RemoteConfig{data: map[string]string{}}
}

func TestRemoteConfig_ParseFile_validConfigPath(t *testing.T) {
	config := NewConfig()

	config.ParseFile(ConfigPath)

	if len(config.data) == 0 {
		t.Error("Empty git config file")
	}
}

func TestRemoteConfig_ParseFile_invalidConfigPath(t *testing.T) {
	config := NewConfig()

	err := config.ParseFile(InvalidConfigPath)

	if err == nil {
		t.Error("Valid git config path")
	}
}

func TestRemoteConfig_Parse(t *testing.T) {

	file, err := os.Open(ConfigPath)

	if err != nil {
		t.Error("Wrong git config path")
	}

	defer file.Close()

	config := NewConfig()

	config.configPath = ConfigPath

	config.Parse(file)

	if ConfigPath !=config.configPath {
		t.Error("Invalid config file path storing")
	}

	if len(config.data) == 0 {
		t.Error("Empty git config file")
	}

}

func TestSshUrlToHttpUrl(t *testing.T) {

	url:=SshUrlToHttpUrl(SshUrl)

	if url !=HttpUrl {
		t.Error("Invalid ssh url")
	}
}