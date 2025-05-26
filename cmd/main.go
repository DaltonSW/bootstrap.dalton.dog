// main is the bootstrap entry point
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-viper/mapstructure/v2"
	"github.com/goccy/go-yaml"
	"go.dalton.dog/setup/cmd/tasks"
)

// RawTaskList represents a sequence of tasks to be executed
type RawTaskList struct {
	Tasks []map[string]any `yaml:"tasks"`
}

func (r RawTaskList) ConvertToTasks() (*tasks.TaskList, error) {
	outList := &tasks.TaskList{}

	outList.Tasks = make([]tasks.Task, 0)

	for _, task := range r.Tasks {
		var newTask tasks.Task
		taskType := task["type"].(string)
		switch taskType {
		case "update_distro_packages":
			newTask = tasks.UpdateDistroPackages{}
		case "upgrade_distro_packages":
			newTask = tasks.UpgradeDistroPackages{}
		case "install_distro_packages":
			newTask = tasks.InstallDistroPackages{}
		case "install_brew":
			newTask = tasks.InstallBrew{}
		case "install_brew_packages":
			newTask = tasks.InstallBrewPackages{}
		case "git_config":
			newTask = tasks.GitConfig{}
		case "clone_repo":
			newTask = tasks.CloneRepo{}
		case "download_file":
			newTask = tasks.DownloadFile{}
		case "run_command":
			newTask = tasks.RunCommand{}
		}

		decodeConfig := mapstructure.DecoderConfig{
			ErrorUnused: true,
			ErrorUnset:  true,
			ZeroFields:  true,
			Squash:      true,
			Result:      &newTask,
		}
		decoder, err := mapstructure.NewDecoder(&decodeConfig)

		if err != nil {
			return nil, err
		}

		err = decoder.Decode(&task)

		if err != nil {
			return nil, err
		}
		outList.Tasks = append(outList.Tasks, newTask)
	}

	return outList, nil
}

func Run() {
	var configFilePath string

	if len(os.Args) != 2 {
		configFilePath = "config.yaml"
		// fmt.Println("Usage: setup <config_file_path>")
		// os.Exit(1)
	} else {
		configFilePath = os.Args[1]
	}

	configBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var rawList RawTaskList

	if err = yaml.Unmarshal(configBytes, &rawList); err != nil {
		log.Fatal(err)
	}

	taskList, err := rawList.ConvertToTasks()

	if err != nil {
		log.Fatal(err)
	}

	log.Info(taskList)

	return
}
