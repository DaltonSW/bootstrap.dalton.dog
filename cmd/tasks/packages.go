package tasks

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/tree"
)

type UpdateDistroPackages struct {
	BaseTask
}

func (t UpdateDistroPackages) Execute() error { /* implementation */ return nil }

func (t UpdateDistroPackages) String() string { return t.Name }

type UpgradeDistroPackages struct {
	BaseTask
}

func (t UpgradeDistroPackages) Execute() error { /* implementation */ return nil }

func (t UpgradeDistroPackages) String() string { return t.Name }

type InstallBrew struct {
	BaseTask
}

func (t InstallBrew) Execute() error { /* implementation */ return nil }

func (t InstallBrew) String() string { return "Installing Linuxbrew" }

type InstallDistroPackages struct {
	BaseTask
	Packages map[string][]string `mapstructure:"packages"`
}

func (t InstallDistroPackages) Execute() error { /* implementation */ return nil }

func (t InstallDistroPackages) String() string { return t.Name }

type InstallBrewPackages struct {
	BaseTask
	Packages []string `mapstructure:"packages"`
}

func (t InstallBrewPackages) Execute() error { /* implementation */ return nil }

func (t InstallBrewPackages) String() string {
	brewTree := tree.New()
	brewTree.Root(t.Name)
	for _, pkg := range t.Packages {
		brewTree.Child(pkg)
	}
	return fmt.Sprint(brewTree)
}
