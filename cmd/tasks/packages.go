package tasks

import "fmt"

type UpdateDistroPackages struct {
	BaseTask
}

func (t UpdateDistroPackages) Execute() error { /* implementation */ return nil }

func (t UpdateDistroPackages) String() string { return "Updating distro package manager sources" }

type UpgradeDistroPackages struct {
	BaseTask
}

func (t UpgradeDistroPackages) Execute() error { /* implementation */ return nil }

func (t UpgradeDistroPackages) String() string {
	return "Upgrading distro package manager managed packages"
}

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

func (t InstallDistroPackages) String() string { return "Updating distro package manager sources" }

type InstallBrewPackages struct {
	BaseTask
	Packages []string `mapstructure:"packages"`
}

func (t InstallBrewPackages) Execute() error { /* implementation */ return nil }

func (t InstallBrewPackages) String() string {
	return fmt.Sprintf("%v\n\t%v", t.Name, t.Packages)
}
