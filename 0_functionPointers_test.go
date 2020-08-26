package apperror

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fmtSprintfExpected             int
	fmtSprintfCalled               int
	fmtErrorfExpected              int
	fmtErrorfCalled                int
	stringsJoinExpected            int
	stringsJoinCalled              int
	jsonMarshalExpected            int
	jsonMarshalCalled              int
	cleanupInnerErrorsFuncExpected int
	cleanupInnerErrorsFuncCalled   int
	wrapErrorFuncExpected          int
	wrapErrorFuncCalled            int
	wrapSimpleErrorFuncExpected    int
	wrapSimpleErrorFuncCalled      int
)

func createMock(t *testing.T) {
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return ""
	}
	jsonMarshalExpected = 0
	jsonMarshalCalled = 0
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		return nil, nil
	}
	cleanupInnerErrorsFuncExpected = 0
	cleanupInnerErrorsFuncCalled = 0
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		return nil
	}
	wrapErrorFuncExpected = 0
	wrapErrorFuncCalled = 0
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		return nil
	}
	wrapSimpleErrorFuncExpected = 0
	wrapSimpleErrorFuncCalled = 0
	wrapSimpleErrorFunc = func(innerErrors []error, messageFormat string, parameters ...interface{}) AppError {
		wrapSimpleErrorFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	jsonMarshal = json.Marshal
	assert.Equal(t, jsonMarshalExpected, jsonMarshalCalled, "Unexpected number of calls to jsonMarshal")
	cleanupInnerErrorsFunc = cleanupInnerErrors
	assert.Equal(t, cleanupInnerErrorsFuncExpected, cleanupInnerErrorsFuncCalled, "Unexpected number of calls to cleanupInnerErrorsFunc")
	wrapErrorFunc = WrapError
	assert.Equal(t, wrapErrorFuncExpected, wrapErrorFuncCalled, "Unexpected number of calls to wrapErrorFunc")
	wrapSimpleErrorFunc = WrapSimpleError
	assert.Equal(t, wrapSimpleErrorFuncExpected, wrapSimpleErrorFuncCalled, "Unexpected number of calls to wrapSimpleErrorFunc")
}
