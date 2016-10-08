package translator

import (
	"io/ioutil"

	"github.com/bitly/go-simplejson"
)

// Lang language and defaults
type Lang struct {
	JSON    *simplejson.Json
	Default string
}

// Language main var
var trans = New("language.json")

// New lanauge init
func New(filename string) *Lang {
	var l Lang
	l.Read(filename)
	return &l
}

func (u *Lang) Read(filename string) (*Lang, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return u, err // could not read file
	}
	u.JSON, err = simplejson.NewJson(file)
	if err != nil {
		return u, err // file not json
	}
	defaultLanguage, err := u.JSON.Get("default").String()
	if err == nil {
		u.Default = defaultLanguage
	}
	return u, err
}

// Translate strings from the language file
func (u *Lang) Translate(text, language string) string {
	if u.JSON == nil {
		return "No language file loaded"
	}
	// Translate to set language
	newtext, err := u.JSON.Get(language).Get(text).String()
	if err == nil {
		return newtext
	}
	// Translate to default language
	newtext, err = u.JSON.Get(u.Default).Get(text).String()
	if err == nil {
		return newtext
	}
	// No translation possible
	return text
}

// SetDefault sets the default language
func (u *Lang) SetDefault(lang string) {
	u.Default = lang
}

// Translate translates a string
func Translate(text, language string) string {
	return trans.Translate(text, language)
}

// SetDefault sets the default language
func SetDefault(lang string) {
	trans.SetDefault(lang)
}

// Read a new language file
func Read(file string) (*Lang, error) {
	return trans.Read(file)
}
