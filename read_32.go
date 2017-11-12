// +build 386 arm

package ievio

import (
	"encoding/binary"
	"fmt"
)

func makeReadData(buf []byte) (*InputEvent, error) {
	input := InputEvent{}
	input.Time.Sec = int32(binary.LittleEndian.Uint32(buf))
	input.Time.Usec = int32(binary.LittleEndian.Uint32(buf[4:]))
	input.Type = EV_TYPE(binary.LittleEndian.Uint16(buf[8:]))
	if code, err := getCodeFromString(input.Type, fmt.Sprintf("%v", binary.LittleEndian.Uint16(buf[10:]))); err == nil {
		input.Code = code
	} else {
		return nil, err
	}

	input.Value = int32(binary.LittleEndian.Uint32(buf[12:]))
	return &input, nil
}
