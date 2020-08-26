package apperror

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAppErrorGetCode(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var code = baseAppError.Code()

	// assert
	assert.Equal(t, expectedCode.String(), code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetHTTPStatusCode(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var code = baseAppError.HTTPStatusCode()

	// assert
	assert.Equal(t, expectedCode.HTTPStatusCode(), code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetError_NoInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var expectedMessage = "(GeneralFailure) {dummy error}"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		nil,
		expectedExtraData,
	}

	// act
	var message = baseAppError.Error()

	// assert
	assert.Equal(t, expectedMessage, message)

	// verify
	verifyAll(t)
}

func TestAppErrorGetError_WithInner(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var expectedMessage = "(GeneralFailure) {dummy error} -> [ [ dummy inner error 1 ] | [ dummy inner error 2 ] | [ dummy inner error 3 ] ]"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var message = baseAppError.Error()

	// assert
	assert.Equal(t, expectedMessage, message)

	// verify
	verifyAll(t)
}

func TestAppErrorGetInnerErrors_NoInner(t *testing.T) {
	// arrange
	var expectedMessage = "dummy error"
	var expectedError = errors.New(expectedMessage)
	var expectedCode = CodeGeneralFailure
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		nil,
		expectedExtraData,
	}

	// act
	var innerErrors = baseAppError.InnerErrors()

	// assert
	assert.Equal(t, 0, len(innerErrors))

	// verify
	verifyAll(t)
}

func TestAppErrorGetInnerErrors_WithInner(t *testing.T) {
	// arrange
	var expectedMessage = "dummy error"
	var expectedError = errors.New(expectedMessage)
	var expectedCode = CodeGeneralFailure
	var expectedInnerError1 = errors.New("some inner error 1")
	var expectedInnerError2 = errors.New("some inner error 2")
	var expectedInnerError3 = errors.New("some inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var innerErrors = baseAppError.InnerErrors()

	// assert
	assert.Equal(t, expectedInnerErrors, innerErrors)

	// verify
	verifyAll(t)
}

func TestAppErrorGetMessages_NoInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var expectedMessage = "(GeneralFailure) {dummy error}"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		nil,
		expectedExtraData,
	}

	// act
	var messages = baseAppError.Messages()

	// assert
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, expectedMessage, messages[0])

	// verify
	verifyAll(t)
}

func TestAppErrorGetMessages_WithNormalInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var expectedMessage = "(GeneralFailure) {dummy error}"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var expectedInnerMessage1 = "  " + expectedInnerError1.Error()
	var expectedInnerMessage2 = "  " + expectedInnerError2.Error()
	var expectedInnerMessage3 = "  " + expectedInnerError3.Error()

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var messages = baseAppError.Messages()

	// assert
	assert.Equal(t, 4, len(messages))
	assert.Equal(t, expectedMessage, messages[0])
	assert.Equal(t, expectedInnerMessage1, messages[1])
	assert.Equal(t, expectedInnerMessage2, messages[2])
	assert.Equal(t, expectedInnerMessage3, messages[3])

	// verify
	verifyAll(t)
}

func TestAppErrorGetMessages_WithAppErrorInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var expectedMessage = "(GeneralFailure) {dummy error}"
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var expectedInnerMessage1 = "  " + expectedInnerError1.Error()
	var expectedInnerMessage2 = "  (GeneralFailure) {" + dummyInnerErrorMessage + "}"
	var expectedInnerMostMessage = "    " + dummyInnerMostErrorMessage
	var expectedInnerMessage3 = "  " + expectedInnerError3.Error()

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 2
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var messages = baseAppError.Messages()

	// assert
	assert.Equal(t, 5, len(messages))
	assert.Equal(t, expectedMessage, messages[0])
	assert.Equal(t, expectedInnerMessage1, messages[1])
	assert.Equal(t, expectedInnerMessage2, messages[2])
	assert.Equal(t, expectedInnerMostMessage, messages[3])
	assert.Equal(t, expectedInnerMessage3, messages[4])

	// verify
	verifyAll(t)
}

func TestAppErrorGetExtraData_EmptyExtraData(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData map[string]interface{}

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var result = baseAppError.ExtraData()

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestAppErrorGetExtraData_WithExtraData(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var dummyData1 = "some data 1"
	var dummyData2 = "some data 2"
	var dummyData3 = "some data 3"

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 3
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		if expectedExtraData["foo"] == v {
			return []byte(dummyData1), errors.New("some error 1")
		} else if expectedExtraData["test"] == v {
			return []byte(dummyData2), errors.New("some error 2")
		} else if expectedExtraData["error"] == v {
			return []byte(dummyData3), errors.New("some error 3")
		}
		return nil, nil
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	var result = baseAppError.ExtraData()

	// assert
	assert.Equal(t, 3, len(result))
	assert.Equal(t, dummyData1, result["foo"])
	assert.Equal(t, dummyData2, result["test"])
	assert.Equal(t, dummyData3, result["error"])

	// verify
	verifyAll(t)
}

func TestAppErrorAppend_NoInnerError(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var dummyInnerErrors = []error{
		nil,
		nil,
		nil,
	}
	var cleanedInnerErrors = []error{}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	baseAppError.Append(
		dummyInnerErrors...,
	)

	// assert
	assert.Equal(t, expectedInnerErrors, baseAppError.innerErrors)

	// verify
	verifyAll(t)
}

func TestAppErrorAppend_HasInnerError(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		nil,
		dummyInnerError2,
		nil,
		dummyInnerError3,
	}
	var expectedExtraData = map[string]interface{}{
		"foo":   "bar",
		"test":  123,
		"error": errors.New("me"),
	}
	var cleanedInnerErrors = []error{
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	baseAppError.Append(
		dummyInnerErrors...,
	)

	// assert
	assert.Equal(t, 6, len(baseAppError.innerErrors))
	assert.Equal(t, expectedInnerErrors[0], baseAppError.innerErrors[0])
	assert.Equal(t, expectedInnerErrors[1], baseAppError.innerErrors[1])
	assert.Equal(t, expectedInnerErrors[2], baseAppError.innerErrors[2])
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[3])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[4])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[5])

	// verify
	verifyAll(t)
}

func TestAppErrorAttach_NilExtraData(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedExtraData map[string]interface{}
	var dummyName = "some name"
	var dummyValue = uuid.New()

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	baseAppError.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.Equal(t, dummyValue, baseAppError.extraData[dummyName])

	// verify
	verifyAll(t)
}

func TestAppErrorAttach_WithExtraData(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &BaseAppError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
		nil,
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyName = "some name"
	var expectedExtraData = map[string]interface{}{
		dummyName: rand.Int(),
	}
	var dummyValue = uuid.New()

	// mock
	createMock(t)

	// SUT
	var baseAppError = &BaseAppError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
		expectedExtraData,
	}

	// act
	baseAppError.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.Equal(t, dummyValue, baseAppError.extraData[dummyName])

	// verify
	verifyAll(t)
}

func TestGetGeneralFailureError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeGeneralFailure, errorCode)
		assert.Equal(t, "An error occurred during execution", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetGeneralFailureError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetUnauthorized(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeUnauthorized, errorCode)
		assert.Equal(t, "Access denied due to authorization error", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetUnauthorized(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetInvalidOperation(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeInvalidOperation, errorCode)
		assert.Equal(t, "Operation (method) not allowed", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetInvalidOperation(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetBadRequestError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeBadRequest, errorCode)
		assert.Equal(t, "Request URI or body is invalid", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetBadRequestError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetNotFoundError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeNotFound, errorCode)
		assert.Equal(t, "Requested resource is not found in the storage", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetNotFoundError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetCircuitBreakError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeCircuitBreak, errorCode)
		assert.Equal(t, "Operation refused due to internal circuit break on correlation ID", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetCircuitBreakError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetOperationLockError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeOperationLock, errorCode)
		assert.Equal(t, "Operation refused due to mutex lock on correlation ID or trip ID", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetOperationLockError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetAccessForbiddenError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeAccessForbidden, errorCode)
		assert.Equal(t, "Operation failed due to access forbidden", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetAccessForbiddenError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetDataCorruptionError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeDataCorruption, errorCode)
		assert.Equal(t, "Operation failed due to internal storage data corruption", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetDataCorruptionError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetNotImplementedError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, CodeNotImplemented, errorCode)
		assert.Equal(t, "Operation failed due to internal business logic not implemented", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var baseAppError = GetNotImplementedError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetCustomError(t *testing.T) {
	// arrange
	var dummyErrorCode = Code(rand.Intn(255))
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var dummyErrorMessage = "some error message"

	// mock
	createMock(t)

	// expect
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, parameters ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return errors.New(dummyErrorMessage)
	}

	// SUT + act
	var baseAppError, ok = GetCustomError(
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyErrorMessage, baseAppError.error.Error())
	assert.Equal(t, dummyErrorCode, baseAppError.code)
	assert.Equal(t, 0, len(baseAppError.innerErrors))

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_NilInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors []error

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_EmptyInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{}

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_NoValidInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{
		nil,
		nil,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_HasValidInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		nil,
		dummyInnerError2,
		nil,
		dummyInnerError3,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Equal(t, 3, len(result))
	assert.Equal(t, dummyInnerError1, result[0])
	assert.Equal(t, dummyInnerError2, result[1])
	assert.Equal(t, dummyInnerError3, result[2])

	// verify
	verifyAll(t)
}

func TestWrapError_Empty(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{
		nil,
		nil,
		nil,
	}
	var dummyErrorCode = Code(rand.Int())
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var cleanedInnerErrors = []error{}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// SUT + act
	var result = WrapError(
		dummyInnerErrors,
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestWrapError_NotEmpty(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		nil,
		dummyInnerError2,
		nil,
		dummyInnerError3,
	}
	var dummyErrorCode = Code(rand.Int())
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var dummyErrorMessage = "some error message"
	var cleanedInnerErrors = []error{
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// expect
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, parameters ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return errors.New(dummyErrorMessage)
	}

	// SUT + act
	var baseAppError, ok = WrapError(
		dummyInnerErrors,
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyErrorMessage, baseAppError.error.Error())
	assert.Equal(t, dummyErrorCode, baseAppError.code)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestWrapSimpleError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var expectedResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 3, len(innerErrors))
		assert.Equal(t, dummyInnerError1, innerErrors[0])
		assert.Equal(t, dummyInnerError2, innerErrors[1])
		assert.Equal(t, dummyInnerError3, innerErrors[2])
		assert.Equal(t, CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return expectedResult
	}

	// SUT + act
	var baseAppError = WrapSimpleError(
		[]error{
			dummyInnerError1,
			dummyInnerError2,
			dummyInnerError3,
		},
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// assert
	assert.Equal(t, expectedResult, baseAppError)

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_NonAppError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some dummy error")

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, dummyError, errs[0])

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_NoInner(t *testing.T) {
	// arrange
	var dummyError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		nil,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 0, len(errs))

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_WithInner(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("dummy inner error")
	var dummyError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyInnerError},
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, dummyInnerError, errs[0])

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_MultiLayer_NoInner(t *testing.T) {
	// arrange
	var dummyThirdLayerError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		nil,
		nil,
	}
	var dummySecondLayerError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyThirdLayerError},
		nil,
	}
	var dummyError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummySecondLayerError},
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 0, len(errs))

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_MultiLayer_WithInner(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("dummy inner error")
	var dummyThirdLayerError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyInnerError},
		nil,
	}
	var dummySecondLayerError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyThirdLayerError},
		nil,
	}
	var dummyError = &BaseAppError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummySecondLayerError},
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, dummyInnerError, errs[0])

	// verify
	verifyAll(t)
}
