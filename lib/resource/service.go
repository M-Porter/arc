package resource

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func CreateService(options *CreateResourceOptions) {
	if options.ProjectName == "" {
		options.ProjectName = ActiveProject().Name
	}

	cfg := config.LoadArcConfig()

	service := config.Service{
		Name:   options.ResourceName,
		Path:   options.ServicePath,
		Branch: options.ServiceBranch,
	}

	if service.Branch == "" {
		service.Branch = options.ProjectName
	}

	project, err := cfg.ProjectByName(options.ProjectName)
	if err != nil {
		util.Fatalf("error: cannot add service to non-existent project")
	}

	project.Services = append(project.Services, service)

	config.WriteArcConfig(cfg)
}
