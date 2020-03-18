package sync

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
	"strings"
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
	repo, err := git.PlainOpen(service.Path)
	if err != nil {
		return fmt.Errorf("git error: %v", err)
	}

	currentBranch, err := currentBranchForService(repo)
	if err != nil {
		return fmt.Errorf("git error: %v", err)
	}

	if currentBranch != service.Branch {
		return fmt.Errorf("service branch %s not equal to current checked out branch %s... skipping", service.Branch, currentBranch)
	}

	util.Printlnf("syncing service %s", service.Name)

	return nil
}

func currentBranchForService(repo *git.Repository) (string, error) {
	branchRefs, err := repo.Branches()
	if err != nil {
		return "", err
	}

	headRef, err := repo.Head()
	if err != nil {
		return "", err
	}

	var currentBranchName string
	err = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchName = branchRef.Name().String()
			return nil
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	currentBranchName = strings.Replace(currentBranchName, "refs/heads/", "", 1)

	return currentBranchName, nil
}
