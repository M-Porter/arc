package sync

import (
	"fmt"
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/git"
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
		util.Printf("syncing service %s => ", svc.Name)
		err := Service(&svc, force)
		if err != nil {
			numErrors += 1
			util.Printlnf("error: %v", err)
		} else {
			util.Printlnf("done")
		}
	}

	if numErrors > 0 {
		return fmt.Errorf("%d services could not be sync'd", numErrors)
	}

	return nil
}

func Service(service *config.Service, force bool) error {
	currentBranch, err := git.BranchAtPath(service.Path)
	if err != nil {
		return err
	}

	if currentBranch != service.Branch {
		return fmt.Errorf("service %s branch %s not equal to current checked out branch %s... skipping", service.Name, service.Branch, currentBranch)
	}

	isDirty, err := git.RepoIsDirty(service.Path)
	if err != nil {
		return err
	} else if isDirty && !force {
		return fmt.Errorf("service %s is dirty... skipping", service.Name)
	}

	err = git.CheckoutAndPull(service.Path, service.Branch)
	if err != nil {
		return err
	}

	return nil
}
