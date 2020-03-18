package config

import (
	"fmt"
	"github.com/m-porter/arc/lib/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

// ArcDirectory returns the arc directory.
func ArcDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		util.Fatalf("could not get user home dir: %v", err)
	}
	return path.Join(homeDir, arcDirectory)
}

// ArcConfigPath returns the arc config file path.
func ArcConfigPath() string {
	return path.Join(ArcDirectory(), arcConfig)
}

// Prints the config as yaml
func Println(cfg interface{}) {
	yamlBytes, err := yaml.Marshal(&cfg)
	if err != nil {
		util.Fatalf("error marshalling config: %v", err)
	}
	fmt.Printf("%s", string(yamlBytes))
}

// EnsureArcConfig verifies that the arc config exists. If it does not it is created.
func EnsureArcConfig() {
	arcPath := ArcDirectory()
	arcPathExists, err := util.PathExists(arcPath)
	if !arcPathExists {
		if err := os.MkdirAll(arcPath, os.ModePerm); err != nil {
			util.Fatalf("error creating arc directory: %v", err)
		}
	} else if err != nil {
		util.Fatalf("error checking if directory exists: %v", err)
	}

	arcConfigPath := ArcConfigPath()
	arcConfigPathExists, err := util.PathExists(arcConfigPath)
	if !arcConfigPathExists {
		f, err := os.Create(arcConfigPath)
		if err != nil {
			util.Fatalf("error created arc config: %v", err)
		}
		err = f.Close()
		if err != nil {
			util.Fatalf("error closing arc config: %v", err)
		}
	} else if err != nil {
		util.Fatalf("error checking if config exists: %v", err)
	}
}

func WriteArcConfig(arcConfig *ArcConfig) {
	arcConfigPath := ArcConfigPath()

	_ = os.Remove(arcConfigPath)

	file, err := os.OpenFile(arcConfigPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		util.Fatalf("error opening arc config: %v", err)
	}

	yamlConf, err := yaml.Marshal(arcConfig)
	if err != nil {
		util.Fatalf("error marshalling arc config: %v", err)
	}

	_, err = file.Write(yamlConf)
	if err != nil {
		util.Fatalf("error writing arc config: %v", err)
	}
}

func LoadArcConfig() *ArcConfig {
	arcConfigPath := ArcConfigPath()

	f, err := os.OpenFile(arcConfigPath, os.O_RDWR, 0644)
	if err != nil {
		util.Fatalf("error opening arc config: %v", err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		util.Fatalf("error opening arc config: %v", err)
	}

	arcConfig := &ArcConfig{}

	if fileInfo.Size() == 0 {
		return arcConfig
	}

	fBytes, err := ioutil.ReadFile(arcConfigPath)
	if err != nil {
		util.Fatalf("error reading arc config: %v", err)
	}

	err = yaml.Unmarshal(fBytes, arcConfig)
	if err != nil {
		util.Fatalf("error unmarshalling arc config: %v", err)
	}

	return arcConfig
}
