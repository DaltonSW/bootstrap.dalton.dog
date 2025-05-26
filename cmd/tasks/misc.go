package tasks

type DownloadFile struct {
	BaseTask
	Source  string `mapstructure:"src"`
	Dest    string `mapstructure:"dest"`
	SHA256  string `mapstructure:"SHA256,omitempty"`
	Extract bool   `mapstructure:"extract,omitempty"`
}

func (t DownloadFile) Execute() error { /* implementation */ return nil }

type RunCommand struct {
	BaseTask
	Command string `mapstructure:"command"`
	Sudo    bool   `mapstructure:"sudo,omitempty"`
}

func (t RunCommand) Execute() error { /* implementation */ return nil }
