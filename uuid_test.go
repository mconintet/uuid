package uuid

import (
	"fmt"
	"testing"
)

func TestUuidV1(t *testing.T) {
	if uuid, err := NewV1(nil); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(uuid.String())
	}
}
