package resource

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func CreateProject(resourceName string, setActive bool) {
	project := config.Project{
		Name: resourceName,
	}

	cfg := config.LoadArcConfig()
	if p, _ := cfg.ProjectByName(project.Name); p != nil {
		util.Fatalf("error: cannot create duplicate project %v: %v", project.Name)
		return
	}

	cfg.Projects = append(cfg.Projects, project)

	if setActive {
		cfg.CurrentProject = project.Name
	}

	config.WriteArcConfig(cfg)
}

func ActiveProject() *config.Project {
	return &config.Project{}
}

func RemoveProject(resourceName string) {
	cfg := config.LoadArcConfig()

	for idx, project := range cfg.Projects {
		if project.Name == resourceName {
			cfg.Projects = append(cfg.Projects[:idx], cfg.Projects[idx+1:]...)

			if cfg.CurrentProject == resourceName {
				cfg.CurrentProject = ""
			}
		}
	}

	config.WriteArcConfig(cfg)

	util.Printlnf("project %s removed", resourceName)
}
