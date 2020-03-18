package config

import (
	"fmt"
)

var arcDirectory = ".arc"
var arcConfig = "arc.yaml"

type ArcConfig struct {
	CurrentProject string    `yaml:"currentProject"`
	Projects       []Project `yaml:"projects"`
}

type Project struct {
	Name     string    `yaml:"name"`
	Services []Service `yaml:"services"`
}

type Service struct {
	Name   string `yaml:"name"`
	Path   string `yaml:"path"`
	Branch string `yaml:"branch"`
}

func (cfg *ArcConfig) ProjectByName(name string) (*Project, error) {
	for _, project := range cfg.Projects {
		if project.Name == name {
			return &project, nil
		}
	}
	return nil, fmt.Errorf("error: project %s not defined", name)
}

func (svc *Project) ServiceByName(name string) (*Service, error) {
	for _, service := range svc.Services {
		if service.Name == name {
			return &service, nil
		}
	}
	return nil, fmt.Errorf("error: service %s not defined", name)
}
