package resource

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func CreateService(resourceName string, projectName string, servicePath string, serviceBranch string) {
	if projectName == "" {
		projectName = ActiveProject().Name
	}

	cfg := config.LoadArcConfig()

	service := config.Service{
		Name:   resourceName,
		Path:   servicePath,
		Branch: serviceBranch,
	}

	if service.Branch == "" {
		service.Branch = projectName
	}

	project, err := cfg.ProjectByName(projectName)
	if err != nil {
		util.Fatalf("error: cannot add service to non-existent project")
		return
	}

	util.Printlnf("creating service %v for project %v", service.Name, projectName)

	project.Services = append(project.Services, service)

	cfg.UpdateProject(*project)

	config.WriteArcConfig(cfg)
}

func RemoveService(resourceName string, projectName string) {
	cfg := config.LoadArcConfig()

	project, err := cfg.ProjectByName(projectName)
	if err != nil {
		util.Fatalf("error: cannot remove service from non-existent project")
		return
	}

	for idx, svc := range project.Services {
		if svc.Name == resourceName {
			project.Services = append(project.Services[:idx], project.Services[idx+1:]...)
		}
	}

	cfg.UpdateProject(*project)

	config.WriteArcConfig(cfg)

	util.Printlnf("service %s remove from project %s", resourceName, projectName)
}
