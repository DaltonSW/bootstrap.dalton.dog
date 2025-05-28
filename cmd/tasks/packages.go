package tasks

import (
	"fmt"

	"go.dalton.dog/setup/cmd/styles"

	"github.com/charmbracelet/lipgloss/tree"
)

type UpdateDistroPackages struct {
	BaseTask
}

func (t UpdateDistroPackages) Execute() error { /* implementation */ return nil }

type UpgradeDistroPackages struct {
	BaseTask
}

func (t UpgradeDistroPackages) Execute() error { /* implementation */ return nil }

type InstallBrew struct {
	BaseTask
}

func (t InstallBrew) Execute() error { /* implementation */ return nil }

type InstallDistroPackages struct {
	BaseTask
	Packages map[string][]string `mapstructure:"packages"`
}

func (t InstallDistroPackages) Execute() error { /* implementation */ return nil }

type InstallBrewPackages struct {
	BaseTask
	Packages []string `mapstructure:"packages"`
}

func (t InstallBrewPackages) Execute() error { /* implementation */ return nil }

func (t InstallBrewPackages) String() string {
	brewTree := tree.New()
	brewTree.Root(styles.TaskStyle.Render(t.Name))
	for _, pkg := range t.Packages {
		brewTree.Child(styles.CommandStyle.Render(pkg))
	}
	return fmt.Sprint(brewTree)
}
