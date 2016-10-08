package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var language Lang

func TestLanguage_Initialization(t *testing.T) {
	// we didn't load a language file yet
	assert.Equal(t, language.Translate("nl", "login with"), "No language file loaded")
}

func TestLanguage_ReadFile(t *testing.T) {
	// File errors
	_, err := language.Read("nonexistingfile")
	assert.NotNil(t, err)
	_, err = language.Read("translator.go")
	assert.NotNil(t, err)
	// Successful read
	_, err = language.Read("language.json.example")
	assert.Nil(t, err)
}

func TestLanguage_Translate(t *testing.T) {
	// Translate existing language
	assert.Equal(t, "aanmelden met", language.Translate("login with", "nl"))
	// Translate non-existing language
	assert.Equal(t, "login with", language.Translate("login with", "non_existing_country_code"))

	// Translate non-existing language with new default
	language.SetDefault("nl")
	assert.Equal(t, "aanmelden met", language.Translate("login with", "non_existing_country_code_new_default"))
}
