package db

import (
	"github.com/nerdzeu/nerdz-core/proto"
)

var (
	dbLang = map[proto.Language]string{
		proto.Language_CROATIAN:   "hr",
		proto.Language_GERMAN:     "de",
		proto.Language_ENGLISH:    "en",
		proto.Language_ITALIAN:    "it",
		proto.Language_PORTUGUESE: "pt",
		proto.Language_ROMANIAN:   "ro",
	}

	protoLang = map[string]proto.Language{
		"de": proto.Language_GERMAN,
		"en": proto.Language_ENGLISH,
		"hr": proto.Language_CROATIAN,
		"it": proto.Language_ITALIAN,
		"pt": proto.Language_PORTUGUESE,
		"ro": proto.Language_ROMANIAN,
	}
)
