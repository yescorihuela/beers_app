package api

import (
	"testing"
)

func TestDescriptionResponse(t *testing.T) {
	var (
		descriptionResponse DescriptionResponse
		genericInterface    interface{}
	)

	descriptionResponse.NewDescriptionResponse("", genericInterface)
	if descriptionResponse.Description != "" {
		t.Error("Invalid value while testing description value must be empty")
	}

	if descriptionResponse.Content != nil {
		t.Error("Invalid value while testing content value must be empty")
	}

	descriptionResponse.NewDescriptionResponse("Any value", struct{ data string }{data: "any data"})

	if descriptionResponse.Description == "" {
		t.Error("Invalid value while testing description value assigment")
	}

	if descriptionResponse.Content == struct{ data string }{} {
		t.Error("Invalid value while testing description value assigment")
	}

}
