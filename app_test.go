package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	except := "Zoo Application"
	actual := AppName()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}
