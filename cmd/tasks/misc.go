package tasks

import "fmt"

type DownloadFile struct {
	BaseTask
	Source  string `mapstructure:"src"`
	Dest    string `mapstructure:"dest"`
	SHA256  string `mapstructure:"SHA256,omitempty"`
	Extract bool   `mapstructure:"extract,omitempty"`
}

func (t DownloadFile) Execute() error { /* implementation */ return nil }

func (t DownloadFile) String() string {
	return fmt.Sprintf("%v\n\t%v -> %v\n\tSHA: %v", t.Name, t.Source, t.Dest, t.SHA256)
}

type RunCommand struct {
	BaseTask
	Command string `mapstructure:"command"`
	Sudo    bool   `mapstructure:"sudo,omitempty"`
}

func (t RunCommand) Execute() error { /* implementation */ return nil }

func (t RunCommand) String() string {
	return fmt.Sprintf("%v\n\t%v", t.Name, t.Command)
}
