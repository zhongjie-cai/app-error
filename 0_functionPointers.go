package apperror

import (
	"errors"
	"fmt"
	"strings"
)

// func pointers for injection / testing: apperror.go
var (
	fmtSprint              = fmt.Sprint
	fmtSprintf             = fmt.Sprintf
	fmtErrorf              = fmt.Errorf
	stringsJoin            = strings.Join
	formatExtraDataFunc    = formatExtraData
	printBaseAppErrorFunc  = printBaseAppError
	getErrorMessageFunc    = getErrorMessage
	printInnerErrorsFunc   = printInnerErrors
	errorsIs               = errors.Is
	equalsErrorFunc        = equalsError
	appErrorContainsFunc   = appErrorContains
	innerErrorContainsFunc = innerErrorContains
	cleanupInnerErrorsFunc = cleanupInnerErrors
	newBaseAppErrorFunc    = NewBaseAppError
)
