package tasks

type GitConfig struct {
	BaseTask
	Settings []GitSetting `mapstructure:"settings"`
}

func (t GitConfig) Execute() error { /* implementation */ return nil }

type GitSetting struct {
	Setting string `mapstructure:"setting"`
	Value   string `mapstructure:"value"`
	Global  bool   `mapstructure:"global,omitempty"`
}

type CloneRepo struct {
	BaseTask
	Source string `mapstructure:"src"`
	Dest   string `mapstructure:"dest"`
}

func (t CloneRepo) Execute() error { /* implementation */ return nil }
