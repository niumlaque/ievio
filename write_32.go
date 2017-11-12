// +build 386 arm

package ievio

import (
	"encoding/binary"
	"syscall"
)

func makeWriteData(eventType, code, value string) ([]byte, *InputEvent, error) {
	buf := make([]byte, InputEventSize)
	input := InputEvent{}
	if et, err := getEventTypeFromString(eventType); err == nil {
		input.Type = et
	} else {
		return nil, nil, err
	}

	if cd, err := getCodeFromString(input.Type, code); err == nil {
		input.Code = cd
	} else {
		return nil, nil, err
	}

	if v, err := stringToUint64(value, 32); err == nil {
		input.Value = int32(v)
	} else {
		return nil, nil, err
	}

	syscall.Gettimeofday(&input.Time)
	binary.LittleEndian.PutUint32(buf[0:], uint32(input.Time.Sec))
	binary.LittleEndian.PutUint32(buf[4:], uint32(input.Time.Usec))
	binary.LittleEndian.PutUint16(buf[8:], uint16(input.Type))
	binary.LittleEndian.PutUint16(buf[10:], input.Code.ValueUint16())
	binary.LittleEndian.PutUint32(buf[12:], uint32(input.Value))

	return buf, &input, nil
}
