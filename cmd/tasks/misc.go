package tasks

type DownloadFile struct {
	BaseTask `yaml:",inline"`
	Source   string `yaml:"src"`
	Dest     string `yaml:"dest"`
	SHA256   string `yaml:"SHA256,omitempty"`
	Extract  bool   `yaml:"extract,omitempty"`
}

func (t DownloadFile) Execute() error { /* implementation */ return nil }

type RunCommand struct {
	BaseTask `yaml:",inline"`
	Command  string `yaml:"command"`
	Sudo     bool   `yaml:"sudo,omitempty"`
}

func (t RunCommand) Execute() error { /* implementation */ return nil }
