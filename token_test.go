package gtranslate

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func TestGetSM(t *testing.T) {
	ttk, err := sM(otto.FalseValue())
	if err != nil {
		t.Error(err)
	}
	if ttk.IsNull() {
		t.Error("ttk is null")
	}
}
