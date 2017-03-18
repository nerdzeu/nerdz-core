package db

import (
	"fmt"

	"github.com/nerdzeu/nerdz-core/utils"
)

var (
	// Languages contains the languages supported by the current NERDZ implementation.
	Languages = utils.StringSet{ // hash set
		"de": struct{}{},
		"en": struct{}{},
		"hr": struct{}{},
		"it": struct{}{},
		"pt": struct{}{},
		"ro": struct{}{},
	}
)

func sanitiseLanguage(lang, defaultLang string) (string, error) {
	if lang == "" {
		return defaultLang, nil
	}

	if Languages.Contains(lang) {
		return lang, nil
	}

	return "", fmt.Errorf("language '%s' is not a valid or supported language", lang)
}
