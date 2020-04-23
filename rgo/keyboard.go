package rgo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-vgo/robotgo"
)

// KeyboardModel ...
type KeyboardModel struct {
	Key string `json:"key"`
}

// HandleKeyboardClick ...
func HandleKeyboardClick(body io.ReadCloser) error {
	var msg KeyboardModel

	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading body: %+v", err)
	}

	err = json.Unmarshal(b, &msg)
	if err != nil {
		return fmt.Errorf("Json unmarshal error: %+v", err)
	}

	if len(msg.Key) > 1 {
		return fmt.Errorf("Input more than one symbol")
	}

	res := robotgo.KeyTap(msg.Key)
	if res != "" {
		return fmt.Errorf("Keyboard tap error %+v", res)
	}

	return nil
}
