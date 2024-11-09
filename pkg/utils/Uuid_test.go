package utils

import "testing"

func TestGenerateUUID(t *testing.T) {
	s := GenerateUUID()
	println(s)
}
