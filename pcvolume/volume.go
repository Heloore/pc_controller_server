package pcvolume

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/itchyny/volume-go"
)

// SetVolumeRequest ...
type SetVolumeRequest struct {
	Volume int `json:"volume"`
}

// GetVolumeRespose ...
type GetVolumeRespose struct {
	Volume int  `json:"volume"`
	Muted  bool `json:"muted"`
}

// GetCurrentVolume ...
func GetCurrentVolume() ([]byte, error) {
	currentVolume, err := volume.GetVolume()
	if err != nil {
		return nil, fmt.Errorf("Get volume failed: %+v", err)
	}

	muted, err := volume.GetMuted()
	if err != nil {
		return nil, fmt.Errorf("Get muted failed: %+v", err)
	}
	resp, err := json.Marshal(GetVolumeRespose{Volume: currentVolume, Muted: muted})
	if err != nil {
		return nil, fmt.Errorf("Json marshal error %+v", err)
	}
	return resp, nil
}

// SetCurrentVolume ...
func SetCurrentVolume(body io.ReadCloser) error {
	var msg SetVolumeRequest

	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading body: %+v", err)
	}

	err = json.Unmarshal(b, &msg)
	if err != nil {
		return fmt.Errorf("Json unmarshal error: %+v", err)
	}

	err = volume.SetVolume(msg.Volume)
	if err != nil {
		return fmt.Errorf("Error setting volume: %+v", err)
	}
	return nil
}

// MuteVolume ...
func MuteVolume() error {
	err := volume.Mute()
	if err != nil {
		return fmt.Errorf("Error muting volume: %+v", err)
	}

	return nil
}

// UnmuteVolume ...
func UnmuteVolume() error {
	err := volume.Unmute()
	if err != nil {
		return fmt.Errorf("Json unmarshal error: %+v", err)
	}
	return nil
}
