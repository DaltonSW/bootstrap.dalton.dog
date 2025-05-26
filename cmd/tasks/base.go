package tasks

// Task represents the minimum things a task should be capable of doing.
type Task interface {
	GetType() string
	GetName() string
	Execute() error
}

// BaseTask is the base implementation of the Task interface.
type BaseTask struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

// GetType returns the type of the task being run.
func (t BaseTask) GetType() string { return t.Type }

// GetName returns the name of the task being run.
// The name is the user's description of the current task.
func (t BaseTask) GetName() string { return t.Name }
