package nicetrace

import (
	"testing"
)

func TestTrace(t *testing.T) {
	defer Print()
	
	panic("ow")
}