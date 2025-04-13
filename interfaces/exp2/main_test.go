package main

import (
	"bytes"
	"testing"
)

func Test_save(t *testing.T) {

	t.Run("check the save func", func(t *testing.T) {
		var buf bytes.Buffer

		err := save(&buf, "hello")
		if err != nil {
			t.Error(err)
		}

		if buf.String() != "hello" {
			t.Errorf("save func: got %q, want %q", buf.String(), "hello")
		}

	})
}
