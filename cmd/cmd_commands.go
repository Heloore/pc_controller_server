package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

// CommandModel ...
type CommandModel struct {
	Command string `json:"command"`
}

func lockPC() error {
	cmd := exec.Command("rundll32.exe", "user32.dll", "LockWorkStation")
	_, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("Error executing command 'lock': %+v", err)
	}

	return nil
}

func shutdown(param string) error {
	cmd := exec.Command("shutdown", param)
	_, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("Error executing command 'lock': %+v", err)
	}

	return nil
}

func validateCommand(command string) error {
	switch command {
	case "Lock":
		break
	case "Sleep":
		break
	case "PowerOff":
		break
	case "Restart":
		break
	case "LogOff":
		break
	default:
		return fmt.Errorf("Command requeted is not in the list")
	}
	return nil
}

// HandleCMD ...
func HandleCMD(body io.ReadCloser) error {
	var msg CommandModel

	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading body: %+v", err)
	}

	err = json.Unmarshal(b, &msg)
	if err != nil {
		return fmt.Errorf("Json unmarshal error: %+v", err)
	}

	if err = validateCommand(msg.Command); err != nil {
		return err
	}

	switch msg.Command {
	case "Lock":
		return lockPC()
	case "Sleep":
		return shutdown("/h")
	case "PowerOff":
		return shutdown("/s")
	case "Restart":
		return shutdown("/r")
	case "LogOff":
		return shutdown("/l")
	}

	return nil
}

