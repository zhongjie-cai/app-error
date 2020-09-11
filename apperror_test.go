package apperror

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseAppError(t *testing.T) {
	// arrange
	var dummyCode = Code(rand.Intn(100))
	var dummyMessageFormat = "some format"
	var dummyParameter1 = "some parameter 2"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some error 3")
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 3, len(a))
		assert.Equal(t, dummyParameter1, a[0])
		assert.Equal(t, dummyParameter2, a[1])
		assert.Equal(t, dummyParameter3, a[2])
		return dummyError
	}

	// SUT + act
	var err = NewBaseAppError(
		dummyCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// assert
	assert.Equal(t, dummyError, err.error)
	assert.Equal(t, dummyCode, err.code)
	assert.Empty(t, err.innerErrors)
	assert.Empty(t, err.extraData)

	// verify
	verifyAll(t)
}

func TestFormatExtraData_NilExtraData(t *testing.T) {
	// arrange
	var dummyExtraData map[string]interface{}

	// mock
	createMock(t)

	// SUT + act
	var result = formatExtraData(
		dummyExtraData,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestFormatExtraData_EmptyExtraData(t *testing.T) {
	// arrange
	var dummyExtraData = map[string]interface{}{}

	// mock
	createMock(t)

	// SUT + act
	var result = formatExtraData(
		dummyExtraData,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestFormatExtraData_WithExtraData(t *testing.T) {
	// arrange
	var dummyName1 = "some name 1"
	var dummyValue1 = "some value 1"
	var dummyName2 = "some name 2"
	var dummyValue2 = rand.Int()
	var dummyName3 = "some name 3"
	var dummyValue3 = errors.New("some error 3")
	var dummyExtraData = map[string]interface{}{
		dummyName1: dummyValue1,
		dummyName2: dummyValue2,
		dummyName3: dummyValue3,
	}
	var dummyMessage1 = "some message 1"
	var dummyMessage2 = "some message 2"
	var dummyMessage3 = "some message 3"
	var dummyJoinedMessage = "some joined message"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 4
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		if fmtSprintfCalled < 4 {
			assert.Equal(t, errorExtraDataFormat, format)
			assert.Equal(t, 2, len(a))
			if dummyName1 == a[0] {
				assert.Equal(t, dummyValue1, a[1])
				return dummyMessage1
			} else if dummyName2 == a[0] {
				assert.Equal(t, dummyValue2, a[1])
				return dummyMessage2
			} else if dummyName3 == a[0] {
				assert.Equal(t, dummyValue3, a[1])
				return dummyMessage3
			}
		} else if fmtSprintfCalled == 4 {
			assert.Equal(t, errorJoiningFormat, format)
			assert.Equal(t, 1, len(a))
			assert.Equal(t, dummyJoinedMessage, a[0])
			return dummyResult
		}
		return ""
	}
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		assert.Contains(t, a, dummyMessage1)
		assert.Contains(t, a, dummyMessage2)
		assert.Contains(t, a, dummyMessage3)
		assert.Equal(t, errorSeparator, sep)
		return dummyJoinedMessage
	}

	// SUT + act
	var result = formatExtraData(
		dummyExtraData,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestBaseAppError_PrintError(t *testing.T) {
	// arrange
	var dummyCode = Code(rand.Intn(100))
	var dummyError = errors.New("some error")
	var dummyExtraData = map[string]interface{}{
		"foo":  "bar",
		"test": rand.Int(),
	}
	var dummyExtraDataMessage = "some extra data message"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	formatExtraDataFuncExpected = 1
	formatExtraDataFunc = func(extraData map[string]interface{}) string {
		formatExtraDataFuncCalled++
		assert.Equal(t, dummyExtraData, extraData)
		return dummyExtraDataMessage
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, errorMessageFormat, format)
		assert.Equal(t, 3, len(a))
		assert.Equal(t, dummyCode, a[0])
		assert.Equal(t, dummyError, a[1])
		assert.Equal(t, dummyExtraDataMessage, a[2])
		return dummyResult
	}

	// SUT
	var sut = &BaseAppError{}

	// act
	var result = sut.PrintError(
		dummyCode,
		dummyError,
		dummyExtraData,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestPrintBaseAppError_HappyPath(t *testing.T) {
	// arrange
	var dummyCode = Code(rand.Intn(100))
	var dummyError = errors.New("some error")
	var dummyExtraData = map[string]interface{}{
		"foo":  "bar",
		"test": rand.Int(),
	}
	var dummyExtraDataMessage = "some extra data message"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	formatExtraDataFuncExpected = 1
	formatExtraDataFunc = func(extraData map[string]interface{}) string {
		formatExtraDataFuncCalled++
		assert.Equal(t, dummyExtraData, extraData)
		return dummyExtraDataMessage
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, errorMessageFormat, format)
		assert.Equal(t, 3, len(a))
		assert.Equal(t, dummyCode, a[0])
		assert.Equal(t, dummyError, a[1])
		assert.Equal(t, dummyExtraDataMessage, a[2])
		return dummyResult
	}

	// SUT
	var sut = &BaseAppError{
		dummyError,
		dummyCode,
		nil,
		dummyExtraData,
	}

	// act
	var result = printBaseAppError(
		sut,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestGetErrorMessage_NormalError(t *testing.T) {
	// arrange
	var dummyErrorMessage = "some error message"
	var dummyError = errors.New(dummyErrorMessage)

	// mock
	createMock(t)

	// SUT + act
	var result = getErrorMessage(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyErrorMessage, result)

	// verify
	verifyAll(t)
}

func TestGetErrorMessage_BaseAppError(t *testing.T) {
	// arrange
	var dummyCode = Code(rand.Intn(100))
	var dummyError = errors.New("some error")
	var dummyInnerErrors = []error{
		errors.New("some inner error 1"),
		errors.New("some inner error 2"),
		errors.New("some inner error 3"),
	}
	var dummyBaseAppError = &BaseAppError{
		dummyError,
		dummyCode,
		dummyInnerErrors,
		nil,
	}
	var dummyBaseErrorMessage = "some base error message"
	var dummyInnerErrorMessage = "some inner error message"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	printBaseAppErrorFuncExpected = 1
	printBaseAppErrorFunc = func(baseAppError *BaseAppError) string {
		printBaseAppErrorFuncCalled++
		assert.Equal(t, dummyBaseAppError, baseAppError)
		return dummyBaseErrorMessage
	}
	printInnerErrorsFuncExpected = 1
	printInnerErrorsFunc = func(innerErrors []error) string {
		printInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return dummyInnerErrorMessage
	}
	fmtSprintExpected = 1
	fmtSprint = func(a ...interface{}) string {
		fmtSprintCalled++
		assert.Equal(t, 2, len(a))
		assert.Equal(t, dummyBaseErrorMessage, a[0])
		assert.Equal(t, dummyInnerErrorMessage, a[1])
		return dummyResult
	}

	// SUT + act
	var result = getErrorMessage(
		dummyBaseAppError,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestPrintInnerErrors_NilInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors []error

	// mock
	createMock(t)

	// SUT + act
	var result = printInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestPrintInnerErrors_EmptyInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{}

	// mock
	createMock(t)

	// SUT + act
	var result = printInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestPrintInnerErrors_WithErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{
		errors.New("some inner error 1"),
		errors.New("some inner error 2"),
		errors.New("some inner error 3"),
	}
	var dummyErrorMessages = []string{
		"some error message 1",
		"some error message 2",
		"some error message 3",
	}
	var dummyJoinedMessage = "some joined message"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	getErrorMessageFuncExpected = len(dummyInnerErrors)
	getErrorMessageFunc = func(err error) string {
		getErrorMessageFuncCalled++
		assert.Equal(t, dummyInnerErrors[getErrorMessageFuncCalled-1], err)
		return dummyErrorMessages[getErrorMessageFuncCalled-1]
	}
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		assert.ElementsMatch(t, dummyErrorMessages, a)
		assert.Equal(t, errorSeparator, sep)
		return dummyJoinedMessage
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, errorJoiningFormat, format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyJoinedMessage, a[0])
		return dummyResult
	}

	// SUT + act
	var result = printInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestBaseAppError_Error(t *testing.T) {
	// arrange
	var dummyCode = Code(rand.Intn(100))
	var dummyError = errors.New("some error")
	var dummyInnerErrors = []error{
		errors.New("some inner error 1"),
		errors.New("some inner error 2"),
		errors.New("some inner error 3"),
	}
	var dummyBaseAppError = &BaseAppError{
		dummyError,
		dummyCode,
		dummyInnerErrors,
		nil,
	}
	var dummyBaseErrorMessage = "some base error message"
	var dummyInnerErrorMessage = "some inner error message"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	printBaseAppErrorFuncExpected = 1
	printBaseAppErrorFunc = func(baseAppError *BaseAppError) string {
		printBaseAppErrorFuncCalled++
		assert.Equal(t, dummyBaseAppError, baseAppError)
		return dummyBaseErrorMessage
	}
	printInnerErrorsFuncExpected = 1
	printInnerErrorsFunc = func(innerErrors []error) string {
		printInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return dummyInnerErrorMessage
	}
	fmtSprintExpected = 1
	fmtSprint = func(a ...interface{}) string {
		fmtSprintCalled++
		assert.Equal(t, 2, len(a))
		assert.Equal(t, dummyBaseErrorMessage, a[0])
		assert.Equal(t, dummyInnerErrorMessage, a[1])
		return dummyResult
	}

	// SUT + act
	var result = dummyBaseAppError.Error()

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestBaseAppError_Code(t *testing.T) {
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

func TestBaseAppError_HTTPStatusCode(t *testing.T) {
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

func TestEqualsError_SameError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyTarget = dummyError

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = equalsError(
		dummyError,
		dummyTarget,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestEqualsError_SameMessage(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyTarget = errors.New("some error")

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = equalsError(
		dummyError,
		dummyTarget,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestEqualsError_ErrorIs(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyTarget = errors.New("some target")
	var dummyResult = rand.Intn(100) > 50

	// mock
	createMock(t)

	// expect
	errorsIsExpected = 1
	errorsIs = func(err, target error) bool {
		errorsIsCalled++
		return dummyResult
	}

	// SUT + act
	var result = equalsError(
		dummyError,
		dummyTarget,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestAppErrorContains_HappyPath(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyBaseAppError = &BaseAppError{
		error: dummyError,
	}

	// mock
	createMock(t)

	// expect
	equalsErrorFuncExpected = 1
	equalsErrorFunc = func(err, target error) bool {
		equalsErrorFuncCalled++
		assert.Equal(t, dummyBaseAppError.error, err)
		assert.Equal(t, dummyError, target)
		return true
	}

	// SUT + act
	var result = appErrorContains(
		dummyBaseAppError,
		dummyError,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestInnerErrorContains_NilInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors []error
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// SUT + act
	var result = innerErrorContains(
		dummyInnerErrors,
		dummyError,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestInnerErrorContains_EmptyInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// SUT + act
	var result = innerErrorContains(
		dummyInnerErrors,
		dummyError,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestInnerErrorContains_NormalInnerError(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("some inner error")
	var dummyInnerErrors = []error{
		dummyInnerError,
	}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	equalsErrorFuncExpected = 1
	equalsErrorFunc = func(err, target error) bool {
		equalsErrorFuncCalled++
		assert.Equal(t, dummyInnerError, err)
		assert.Equal(t, dummyError, target)
		return true
	}

	// SUT + act
	var result = innerErrorContains(
		dummyInnerErrors,
		dummyError,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestInnerErrorContains_BaseAppErrorInnerError(t *testing.T) {
	// arrange
	var dummyInnerError = &BaseAppError{
		error: errors.New("some inner error"),
	}
	var dummyInnerErrors = []error{
		dummyInnerError,
	}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	appErrorContainsFuncExpected = 1
	appErrorContainsFunc = func(appError AppError, err error) bool {
		appErrorContainsFuncCalled++
		assert.Equal(t, dummyInnerError, appError)
		assert.Equal(t, dummyError, err)
		return true
	}

	// SUT + act
	var result = innerErrorContains(
		dummyInnerErrors,
		dummyError,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestInnerErrorContains_NoMatchingErrors(t *testing.T) {
	// arrange
	var dummyInnerError1 = &BaseAppError{
		error: errors.New("some inner error 1"),
	}
	var dummyInnerError2 = errors.New("some inner error 2")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		dummyInnerError2,
	}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	appErrorContainsFuncExpected = 1
	appErrorContainsFunc = func(appError AppError, err error) bool {
		appErrorContainsFuncCalled++
		assert.Equal(t, dummyInnerError1, appError)
		assert.Equal(t, dummyError, err)
		return false
	}
	equalsErrorFuncExpected = 1
	equalsErrorFunc = func(err, target error) bool {
		equalsErrorFuncCalled++
		assert.Equal(t, dummyInnerError2, err)
		assert.Equal(t, dummyError, target)
		return false
	}

	// SUT + act
	var result = innerErrorContains(
		dummyInnerErrors,
		dummyError,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestBaseAppError_Contains_DirectEqual(t *testing.T) {
	// arrange
	var dummyBaseAppError = &BaseAppError{
		error: errors.New("some base app error"),
	}

	// mock
	createMock(t)

	// SUT
	var sut = dummyBaseAppError

	// act
	var result = sut.Contains(
		dummyBaseAppError,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestBaseAppError_Contains_ErrorEqual(t *testing.T) {
	// arrange
	var dummyBaseAppError = &BaseAppError{
		error: errors.New("some base app error"),
	}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	equalsErrorFuncExpected = 1
	equalsErrorFunc = func(err, target error) bool {
		equalsErrorFuncCalled++
		assert.Equal(t, dummyBaseAppError.error, err)
		assert.Equal(t, dummyError, target)
		return true
	}

	// SUT
	var sut = dummyBaseAppError

	// act
	var result = sut.Contains(
		dummyError,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestBaseAppError_Contains_InnerErrorEqual(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{
		errors.New("some inner error 1"),
		errors.New("some inner error 2"),
		errors.New("some inner error 3"),
	}
	var dummyBaseAppError = &BaseAppError{
		error:       errors.New("some base app error"),
		innerErrors: dummyInnerErrors,
	}
	var dummyError = errors.New("some error")
	var dummyResult = rand.Intn(100) > 50

	// mock
	createMock(t)

	// expect
	equalsErrorFuncExpected = 1
	equalsErrorFunc = func(err, target error) bool {
		equalsErrorFuncCalled++
		assert.Equal(t, dummyBaseAppError.error, err)
		assert.Equal(t, dummyError, target)
		return false
	}
	innerErrorContainsFuncExpected = 1
	innerErrorContainsFunc = func(innerErrors []error, err error) bool {
		innerErrorContainsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return dummyResult
	}

	// SUT
	var sut = dummyBaseAppError

	// act
	var result = sut.Contains(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyResult, result)

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

func TestAppErrorWrap_NoInnerError(t *testing.T) {
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
	baseAppError.Wrap(
		dummyInnerErrors...,
	)

	// assert
	assert.Equal(t, expectedInnerErrors, baseAppError.innerErrors)

	// verify
	verifyAll(t)
}

func TestAppErrorWrap_HasInnerError(t *testing.T) {
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
	baseAppError.Wrap(
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
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeGeneralFailure, code)
		assert.Equal(t, "An error occurred during execution", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetGeneralFailureError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetUnauthorized(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeUnauthorized, code)
		assert.Equal(t, "Access denied due to authorization error", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetUnauthorized(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetInvalidOperation(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeInvalidOperation, code)
		assert.Equal(t, "Operation (method) not allowed", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetInvalidOperation(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetBadRequestError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeBadRequest, code)
		assert.Equal(t, "Request URI or body is invalid", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetBadRequestError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetNotFoundError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeNotFound, code)
		assert.Equal(t, "Requested resource is not found in the storage", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetNotFoundError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetCircuitBreakError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeCircuitBreak, code)
		assert.Equal(t, "Operation refused due to internal circuit break on correlation ID", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetCircuitBreakError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetOperationLockError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeOperationLock, code)
		assert.Equal(t, "Operation refused due to mutex lock on correlation ID or trip ID", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetOperationLockError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetAccessForbiddenError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeAccessForbidden, code)
		assert.Equal(t, "Operation failed due to access forbidden", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetAccessForbiddenError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetDataCorruptionError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeDataCorruption, code)
		assert.Equal(t, "Operation failed due to internal storage data corruption", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetDataCorruptionError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestGetNotImplementedError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("dummy inner error 1")
	var dummyInnerError2 = errors.New("dummy inner error 2")
	var dummyInnerError3 = errors.New("dummy inner error 3")
	var dummyResult = &BaseAppError{}

	// mock
	createMock(t)

	// expect
	newBaseAppErrorFuncExpected = 1
	newBaseAppErrorFunc = func(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
		newBaseAppErrorFuncCalled++
		assert.Equal(t, CodeNotImplemented, code)
		assert.Equal(t, "Operation failed due to internal business logic not implemented", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyResult
	}
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Error(t, dummyInnerError1, innerErrors[0])
		assert.Error(t, dummyInnerError2, innerErrors[1])
		assert.Error(t, dummyInnerError3, innerErrors[2])
		return innerErrors
	}

	// SUT + act
	var baseAppError, ok = GetNotImplementedError(
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	).(*BaseAppError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyResult, baseAppError)
	assert.Equal(t, 3, len(baseAppError.innerErrors))
	assert.Equal(t, dummyInnerError1, baseAppError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, baseAppError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, baseAppError.innerErrors[2])

	// verify
	verifyAll(t)
}
