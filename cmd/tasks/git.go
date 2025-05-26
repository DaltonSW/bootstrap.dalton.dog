package tasks

type GitConfig struct {
	BaseTask `yaml:",inline"`
	Settings []GitSetting `yaml:"settings"`
}

func (t GitConfig) Execute() error { /* implementation */ return nil }

type GitSetting struct {
	Setting string `yaml:"setting"`
	Value   string `yaml:"value"`
	Global  bool   `yaml:"global,omitempty"`
}

type CloneRepo struct {
	BaseTask `yaml:",inline"`
	Source   string `yaml:"src"`
	Dest     string `yaml:"dest"`
}

func (t CloneRepo) Execute() error { /* implementation */ return nil }
