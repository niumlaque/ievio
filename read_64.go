// +build amd64

package ievio

import (
	"encoding/binary"
	"fmt"
)

func makeReadData(buf []byte) (*InputEvent, error) {
	input := InputEvent{}
	input.Time.Sec = int64(binary.LittleEndian.Uint64(buf))
	input.Time.Usec = int64(binary.LittleEndian.Uint64(buf[8:]))
	input.Type = EV_TYPE(binary.LittleEndian.Uint16(buf[16:]))
	if code, err := getCodeFromString(input.Type, fmt.Sprintf("%v", binary.LittleEndian.Uint16(buf[18:]))); err == nil {
		input.Code = code
	} else {
		return nil, err
	}

	input.Value = int32(binary.LittleEndian.Uint32(buf[20:]))
	return &input, nil
}
