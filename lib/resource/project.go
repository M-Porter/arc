package resource

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func CreateProject(options *CreateResourceOptions) {
	project := config.Project{
		Name: options.ResourceName,
	}

	cfg := config.LoadArcConfig()
	if p, _ := cfg.ProjectByName(project.Name); p != nil {
		util.Fatalf("error: cannot create duplicate project %v: %v", project.Name)
	}

	cfg.Projects = append(cfg.Projects, project)

	config.WriteArcConfig(cfg)
}

func ActiveProject() *config.Project {
	return &config.Project{}
}
