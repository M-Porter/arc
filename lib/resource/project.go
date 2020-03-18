package resource

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func CreateProject(resourceName string) {
	project := config.Project{
		Name: resourceName,
	}

	cfg := config.LoadArcConfig()
	if p, _ := cfg.ProjectByName(project.Name); p != nil {
		util.Fatalf("error: cannot create duplicate project %v: %v", project.Name)
		return
	}

	cfg.Projects = append(cfg.Projects, project)

	config.WriteArcConfig(cfg)
}

func ActiveProject() *config.Project {
	return &config.Project{}
}
