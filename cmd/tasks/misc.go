package tasks

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/tree"
	"go.dalton.dog/setup/cmd/styles"
)

type DownloadFile struct {
	BaseTask
	Source  string `mapstructure:"src"`
	Dest    string `mapstructure:"dest"`
	SHA256  string `mapstructure:"SHA256,omitempty"`
	Extract bool   `mapstructure:"extract,omitempty"`
}

func (t DownloadFile) Execute() error { /* implementation */ return nil }

func (t DownloadFile) String() string {
	tree := tree.New()
	tree.Root(styles.TaskStyle.Render(t.Name)).
		Child(fmt.Sprintf("SRC: %v", t.Source)).
		Child(fmt.Sprintf("DST: %v", t.Dest))
	if t.SHA256 != "" {

		tree.Child(fmt.Sprintf("SHA: %v", t.SHA256))
	}
	return fmt.Sprint(tree)
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
