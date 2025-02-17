package controllers

import (
	// "example.com/mod/models"
	// "encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMarkets(t *testing.T) {
	writer := makeRequest("GET", "/api/markets", nil)

	assert.Equal(t, http.StatusOK, writer.Code)
}