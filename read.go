package ievio

import (
	"os"
)

func Read(dev string, handler func(*InputEvent)) error {
	f, err := os.Open(dev)
	if err != nil {
		return err
	}

	defer f.Close()
	buf := make([]byte, InputEventSize)
	for {
		len, err := f.Read(buf)
		if err != nil {
			return err
		}
		if len == 0 {
			break
		}

		input, err := makeReadData(buf)
		if err != nil {
			return err
		}

		handler(input)
	}

	return nil
}
