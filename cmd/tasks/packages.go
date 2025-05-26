package tasks

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
