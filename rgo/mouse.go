package rgo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-vgo/robotgo"
)

// MouseWheelModel ...
type MouseWheelModel struct {
	Direction string `json:"direction"`
	Offset    int    `json:"offset"`
}

// MouseMoveModel ...
type MouseMoveModel struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func validateMouseWheelModel(model *MouseWheelModel) error {
	if model.Direction != "up" && model.Direction != "down" {
		return fmt.Errorf("Error validating mouse wheel model: wheel direction must be either up or down")
	}
	return nil
}

// HandleMouseWheel ...
func HandleMouseWheel(body io.ReadCloser) error {
	var msg MouseWheelModel

	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading body: %+v", err)
	}

	err = json.Unmarshal(b, &msg)
	if err != nil {
		return fmt.Errorf("Json unmarshal error: %+v", err)
	}

	err = validateMouseWheelModel(&msg)
	if err != nil {
		return err
	}

	robotgo.ScrollMouse(msg.Offset, msg.Direction)
	return nil
}

// HandleMouseMovement ...
func HandleMouseMovement(body io.ReadCloser) error {
	var msg MouseMoveModel

	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading body: %+v", err)
	}

	err = json.Unmarshal(b, &msg)
	if err != nil {
		return fmt.Errorf("Json unmarshal error: %+v", err)
	}

	xPos, yPos := robotgo.GetMousePos()

	xPos += msg.X
	yPos += msg.Y

	robotgo.MoveMouseSmooth(xPos, yPos)
	return nil
}
