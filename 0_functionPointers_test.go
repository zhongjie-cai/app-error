package apperror

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fmtSprintExpected              int
	fmtSprintCalled                int
	fmtSprintfExpected             int
	fmtSprintfCalled               int
	fmtErrorfExpected              int
	fmtErrorfCalled                int
	stringsJoinExpected            int
	stringsJoinCalled              int
	formatExtraDataFuncExpected    int
	formatExtraDataFuncCalled      int
	printBaseAppErrorFuncExpected  int
	printBaseAppErrorFuncCalled    int
	getErrorMessageFuncExpected    int
	getErrorMessageFuncCalled      int
	printInnerErrorsFuncExpected   int
	printInnerErrorsFuncCalled     int
	errorsIsExpected               int
	errorsIsCalled                 int
	equalsErrorFuncExpected        int
	equalsErrorFuncCalled          int
	appErrorContainsFuncExpected   int
	appErrorContainsFuncCalled     int
	innerErrorContainsFuncExpected int
	innerErrorContainsFuncCalled   int
	cleanupInnerErrorsFuncExpected int
	cleanupInnerErrorsFuncCalled   int
	newBaseAppErrorFuncExpected    int
	newBaseAppErrorFuncCalled      int
)

func createMock(t *testing.T) {
	fmtSprintExpected = 0
	fmtSprintCalled = 0
	fmtSprint = func(a ...interface{}) string {
		fmtSprintCalled++
		return ""
	}
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
	formatExtraDataFuncExpected = 0
	formatExtraDataFuncCalled = 0
	formatExtraDataFunc = func(extraData map[string]interface{}) string {
		formatExtraDataFuncCalled++
		return ""
	}
	printBaseAppErrorFuncExpected = 0
	printBaseAppErrorFuncCalled = 0
	printBaseAppErrorFunc = func(baseAppError *BaseAppError) string {
		printBaseAppErrorFuncCalled++
		return ""
	}
	getErrorMessageFuncExpected = 0
	getErrorMessageFuncCalled = 0
	getErrorMessageFunc = func(err error) string {
		getErrorMessageFuncCalled++
		return ""
	}
	printInnerErrorsFuncExpected = 0
	printInnerErrorsFuncCalled = 0
	printInnerErrorsFunc = func(innerErrors []error) string {
		printInnerErrorsFuncCalled++
		return ""
	}
	errorsIsExpected = 0
	errorsIsCalled = 0
	errorsIs = func(err, target error) bool {
		errorsIsCalled++
		return false
	}
	equalsErrorFuncExpected = 0
	equalsErrorFuncCalled = 0
	equalsErrorFunc = func(err, target error) bool {
		equalsErrorFuncCalled++
		return false
	}
	appErrorContainsFuncExpected = 0
	appErrorContainsFuncCalled = 0
	appErrorContainsFunc = func(appError AppError, err error) bool {
		appErrorContainsFuncCalled++
		return false
	}
	innerErrorContainsFuncExpected = 0
	innerErrorContainsFuncCalled = 0
	innerErrorContainsFunc = func(innerErrors []error, err error) bool {
		innerErrorContainsFuncCalled++
		return false
	}
	cleanupInnerErrorsFuncExpected = 0
	cleanupInnerErrorsFuncCalled = 0
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		return nil
	}
	newBaseAppErrorFuncExpected = 0
	newBaseAppErrorFuncCalled = 0
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	fmtSprint = fmt.Sprint
	assert.Equal(t, fmtSprintExpected, fmtSprintCalled, "Unexpected number of calls to fmtSprint")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	formatExtraDataFunc = formatExtraData
	assert.Equal(t, formatExtraDataFuncExpected, formatExtraDataFuncCalled, "Unexpected number of calls to formatExtraDataFunc")
	printBaseAppErrorFunc = printBaseAppError
	assert.Equal(t, printBaseAppErrorFuncExpected, printBaseAppErrorFuncCalled, "Unexpected number of calls to printBaseAppErrorFunc")
	getErrorMessageFunc = getErrorMessage
	assert.Equal(t, getErrorMessageFuncExpected, getErrorMessageFuncCalled, "Unexpected number of calls to getErrorMessageFunc")
	printInnerErrorsFunc = printInnerErrors
	assert.Equal(t, printInnerErrorsFuncExpected, printInnerErrorsFuncCalled, "Unexpected number of calls to printInnerErrorsFunc")
	errorsIs = errors.Is
	assert.Equal(t, errorsIsExpected, errorsIsCalled, "Unexpected number of calls to errorsIs")
	equalsErrorFunc = equalsError
	assert.Equal(t, equalsErrorFuncExpected, equalsErrorFuncCalled, "Unexpected number of calls to equalsErrorFunc")
	appErrorContainsFunc = appErrorContains
	assert.Equal(t, appErrorContainsFuncExpected, appErrorContainsFuncCalled, "Unexpected number of calls to appErrorContainsFunc")
	innerErrorContainsFunc = innerErrorContains
	assert.Equal(t, innerErrorContainsFuncExpected, innerErrorContainsFuncCalled, "Unexpected number of calls to innerErrorContainsFunc")
	cleanupInnerErrorsFunc = cleanupInnerErrors
	assert.Equal(t, cleanupInnerErrorsFuncExpected, cleanupInnerErrorsFuncCalled, "Unexpected number of calls to cleanupInnerErrorsFunc")
	newBaseAppErrorFunc = NewBaseAppError
	assert.Equal(t, newBaseAppErrorFuncExpected, newBaseAppErrorFuncCalled, "Unexpected number of calls to newBaseAppErrorFunc")
}
