package tasks

type UpdateDistroPackages struct {
	BaseTask `yaml:",inline"`
}

func (t UpdateDistroPackages) Execute() error { /* implementation */ return nil }

type UpgradeDistroPackages struct {
	BaseTask `yaml:",inline"`
}

func (t UpgradeDistroPackages) Execute() error { /* implementation */ return nil }

type InstallBrew struct {
	BaseTask `yaml:",inline"`
}

func (t InstallBrew) Execute() error { /* implementation */ return nil }

type InstallDistroPackages struct {
	BaseTask `yaml:",inline"`
	Packages map[string][]string `yaml:"packages"`
}

func (t InstallDistroPackages) Execute() error { /* implementation */ return nil }

type InstallBrewPackages struct {
	BaseTask `yaml:",inline"`
	Packages []string `yaml:"packages"`
}

func (t InstallBrewPackages) Execute() error { /* implementation */ return nil }
