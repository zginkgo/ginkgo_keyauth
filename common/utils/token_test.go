package utils

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	v := MakeBearer(24)
	fmt.Println(v)
	t.Log(v)
}
