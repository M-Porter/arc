package sync

import (
	"fmt"
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
)

func ProjectByName(name string, force bool) error {
	cfg := config.LoadArcConfig()

	project, err := cfg.ProjectByName(name)
	if err != nil {
		util.Fatalf("%v", err)
	}

	return Project(project, force)
}

func Project(project *config.Project, force bool) error {
	util.Printlnf("syncing project %s", project.Name)

	var numErrors int8 = 0

	for _, svc := range project.Services {
		err := Service(&svc, force)
		if err != nil {
			numErrors += 1
			util.Printlnf("%v", err)
		}
	}

	if numErrors > 0 {
		return fmt.Errorf("%d services could not be sync'd", numErrors)
	}

	return nil
}

func Service(service *config.Service, force bool) error {
	util.Printlnf("syncing service %s", service.Name)
	return nil
}
