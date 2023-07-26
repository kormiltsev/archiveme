package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	// Call the New() function to create a new Coder instance
	coder, err := New()
	if err == nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if coder != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestSetDefault(t *testing.T) {
	assert := assert.New(t)

	cod := Coder{
		Password:   "test",
		SourceName: "test",
		Result:     "./test",
		FileType:   ".test",
	}

	cod.setDefault()
	assert.Equal(cod.Password, "0000", "Default Password in Coder was changed in app/default.go")
	assert.Equal(cod.Result, ".", "Default Result in Coder was changed in app/default.go")
	assert.Equal(cod.FileType, ".archiveme", "Default FileType in Coder was changed in app/default.go")
}
