package ievio

import (
	"os"
)

func Write(dev, eventType, code, value string) (*InputEvent, error) {
	f, err := os.OpenFile(dev, os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModeCharDevice)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	buf, input, err := makeWriteData(eventType, code, value)
	if err != nil {
		return nil, err
	}

	if _, err = f.Write(buf); err != nil {
		return nil, err
	}

	f.Sync()

	return input, nil
}
