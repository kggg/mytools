package config

import (
	"fmt"
	"testing"
)

func TestReadconfig(t *testing.T) {
	sec := "nocserver"
	section, err := Readconfig(sec)
	if err != nil {
		t.Error(err)
	}
	for k, v := range section {
		fmt.Println(k, v)
	}

}
