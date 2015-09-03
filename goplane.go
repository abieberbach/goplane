package goplane

import (
	"crypto/rand"
	"fmt"
)


func IdGenerator() string {
	buffer := make([]byte, 16)
	rand.Read(buffer)
	return fmt.Sprintf("%X-%X-%X-%X-%X", buffer[0:4], buffer[4:6], buffer[6:8], buffer[8:10], buffer[10:])
}

func FromBoolToInt(value bool) int {
	iValue := 0
	if value {
		iValue=1
	}
	return iValue
}