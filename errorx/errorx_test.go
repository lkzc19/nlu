package errorx

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	var err error
	Check(err)

	fmt.Println("=== NLU ===")

	err = errors.New("42")
	Check(err)
}
