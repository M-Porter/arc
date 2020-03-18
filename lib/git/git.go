package git

import (
	"fmt"
	"github.com/m-porter/arc/lib/util"
	"strings"
)

func BranchAtPath(path string) (string, error) {
	out, err := util.RunCmdAt("git rev-parse --abbrev-ref HEAD", path)
	if err != nil {
		return "", fmt.Errorf("git error: %v", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func RepoIsDirty(path string) (bool, error) {
	out, err := util.RunCmdAt("git diff --stat", path)
	if err != nil {
		return false, fmt.Errorf("git error: %v", err)
	}
	return strings.TrimSpace(string(out)) != "", nil
}

func CheckoutAndPull(path string, branch string) error {
	_, err := util.RunCmdAt(fmt.Sprintf("git checkout %s && git pull origin %s", branch, branch), path)
	if err != nil {
		return fmt.Errorf("git error: %v", err)
	}
	return nil
}
