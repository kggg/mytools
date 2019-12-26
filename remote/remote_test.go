package remote

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	client, err := NewClient("10.68.2.30", "steven", "steven", 22, "nocserver", false)
	if err != nil {
		t.Error(err)
	}
	res, err := client.Run("ls -al")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(res))
}
