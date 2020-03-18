package resource

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func CreateProject(options *CreateResourceOptions) {
	project := config.Project{
		Name: options.ProjectName,
	}

	cfg := config.LoadArcConfig()
	if _, err := cfg.ProjectByName(project.Name); err != nil {
		util.Fatalf("error: cannot create duplicate project %s", project.Name)
	}

	cfg.Projects = append(cfg.Projects, project)

	config.WriteArcConfig(cfg)
}

func ActiveProject() *config.Project {
	return &config.Project{}
}
