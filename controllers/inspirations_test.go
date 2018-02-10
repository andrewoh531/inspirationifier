package controllers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"spaceship/lib"
)

const ProtocolNotSpecified = "www.nba.com"

func TestWhenProtocolNotSpecifiedUserErrorShouldBeReturned(t *testing.T) {
	_, actualError := CreateInspiration(ProtocolNotSpecified, "")
	expectedError := lib.NewUserError("Please provide 'http' or 'https' protocol.")
	assert.Equal(t, expectedError, actualError)
}