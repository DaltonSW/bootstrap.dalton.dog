package tasks

import "fmt"

type GitConfig struct {
	BaseTask
	Settings []GitSetting `mapstructure:"settings"`
}

func (t GitConfig) Execute() error { /* implementation */ return nil }

func (t GitConfig) String() string {
	return fmt.Sprintf("%v\n\tRequest to set %v settings", t.Name, len(t.Settings))
}

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

func (t CloneRepo) String() string {
	return fmt.Sprintf("%v\n\t%v -> %v", t.Name, t.Source, t.Dest)
}
