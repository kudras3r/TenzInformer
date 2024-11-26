package grab

import (
	"os"

	"gopkg.in/yaml.v2"
)

type OS struct {
	Family   string `yaml:"family"`
	OSName   string `yaml:"name"`
	Kernel   string `yaml:"kernel"`
	Codename string `yaml:"codename"`
	Type     string `yaml:"type"`
	Platform string `yaml:"platform"`
	Version  string `yaml:"version"`
}

type Info struct {
	OS     OS     `yaml:"os"`
	PCName string `yaml:"name"`
	MAC    string `yaml:"mac"`
}

func PCInfo(confFile string) (Info, error) {
	data, err := os.ReadFile(confFile)
	if err != nil {
		return Info{}, err
	}

	var info Info

	err = yaml.Unmarshal(data, &info)
	if err != nil {
		return Info{}, err
	}

	return info, nil
}
