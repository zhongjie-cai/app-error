package apperror

import (
	"encoding/json"
	"fmt"
	"strings"
)

// func pointers for injection / testing: apperror.go
var (
	fmtSprintf             = fmt.Sprintf
	fmtErrorf              = fmt.Errorf
	stringsJoin            = strings.Join
	jsonMarshal            = json.Marshal
	cleanupInnerErrorsFunc = cleanupInnerErrors
	wrapErrorFunc          = WrapError
	wrapSimpleErrorFunc    = WrapSimpleError
)
