package apperror

// AppError is the error wrapper interface for all WebServiceTemplate service generated errors
type AppError interface {
	// Golang internal error interface
	error
	// Code returns the string representation of the error code enum
	Code() string
	// HTTPStatusCode returns the corresponding HTTP status code mapped to the error code value
	HTTPStatusCode() int
	// InnerErrors returns the inner errors array
	InnerErrors() []error
	// Messages returns the string representations of all inner errors
	Messages() []string
	// ExtraData returns the serialized map of all attached extra data
	ExtraData() map[string]string
	// Append adds the given list of inner errors into the current app error object
	Append(innerErrors ...error)
	// Attach adds the given value to the current app error's extra data map by given name
	Attach(name string, value interface{})
}

// These are print formatting related constants
const (
	errorPrintFormat   string = "(%v) {%v}" // (Code)Message
	errorPointer       string = " -> "
	errorHolderLeft    string = "[ "
	errorHolderRight   string = " ]"
	errorSeparator     string = " | "
	errorMessageIndent string = "  "
)

// BaseAppError instantiates the AppError interface and provides a base for inheritance
type BaseAppError struct {
	error
	code        Code
	innerErrors []error
	extraData   map[string]interface{}
}

// Code returns string representation of the error code of the app error
func (baseAppError *BaseAppError) Code() string {
	return baseAppError.code.String()
}

// HTTPStatusCode returns HTTP status code according to the error code of the app error
func (baseAppError *BaseAppError) HTTPStatusCode() int {
	return baseAppError.code.HTTPStatusCode()
}

func (baseAppError *BaseAppError) Error() string {
	var fullMessage = fmtSprintf(
		errorPrintFormat,
		baseAppError.code,
		baseAppError.error.Error(),
	)
	var innerMessages []string
	for _, innerError := range baseAppError.innerErrors {
		innerMessages = append(
			innerMessages,
			errorHolderLeft+innerError.Error()+errorHolderRight,
		)
	}
	var innerMessage = stringsJoin(innerMessages, errorSeparator)
	if innerMessage != "" {
		fullMessage += errorPointer + errorHolderLeft + innerMessage + errorHolderRight
	}
	return fullMessage
}

// InnerErrors returns all contained inner errors of the app error
func (baseAppError *BaseAppError) InnerErrors() []error {
	return baseAppError.innerErrors
}

// Messages returns a list of error messages based on the app error itself and its contained inner errors
func (baseAppError *BaseAppError) Messages() []string {
	var messages = []string{
		fmtSprintf(
			errorPrintFormat,
			baseAppError.Code(),
			baseAppError.error.Error(),
		),
	}
	for _, innerError := range baseAppError.innerErrors {
		var typedError, isTyped = innerError.(AppError)
		if isTyped {
			var innerMessages = typedError.Messages()
			for _, innerMessage := range innerMessages {
				messages = append(
					messages,
					errorMessageIndent+innerMessage,
				)
			}
		} else {
			messages = append(
				messages,
				errorMessageIndent+innerError.Error(),
			)
		}
	}
	return messages
}

// ExtraData returns the full attachments added to the app error (with all values in string representations)
func (baseAppError *BaseAppError) ExtraData() map[string]string {
	var result = map[string]string{}
	for name, value := range baseAppError.extraData {
		var bytes, _ = jsonMarshal(value)
		result[name] = string(bytes)
	}
	return result
}

// Append allows consumer to add a list of errors to the app error
func (baseAppError *BaseAppError) Append(innerErrors ...error) {
	var cleanedInnerErrors = cleanupInnerErrorsFunc(innerErrors)
	if len(cleanedInnerErrors) == 0 {
		return
	}
	baseAppError.innerErrors = append(
		baseAppError.innerErrors,
		cleanedInnerErrors...,
	)
}

// Attach allows consumer to add/update a key-value pair to the app error
func (baseAppError *BaseAppError) Attach(name string, value interface{}) {
	if baseAppError.extraData == nil {
		baseAppError.extraData = map[string]interface{}{}
	}
	baseAppError.extraData[name] = value
}

// GetGeneralFailureError creates a generic error based on GeneralFailure
func GetGeneralFailureError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeGeneralFailure,
		"An error occurred during execution",
	)
}

// GetUnauthorized creates an error related to Unauthorized
func GetUnauthorized(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeUnauthorized,
		"Access denied due to authorization error",
	)
}

// GetInvalidOperation creates an error related to InvalidOperation
func GetInvalidOperation(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeInvalidOperation,
		"Operation (method) not allowed",
	)
}

// GetBadRequestError creates an error related to BadRequest
func GetBadRequestError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeBadRequest,
		"Request URI or body is invalid",
	)
}

// GetNotFoundError creates an error related to NotFound
func GetNotFoundError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeNotFound,
		"Requested resource is not found in the storage",
	)
}

// GetCircuitBreakError creates an error related to CircuitBreak
func GetCircuitBreakError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeCircuitBreak,
		"Operation refused due to internal circuit break on correlation ID",
	)
}

// GetOperationLockError creates an error related to OperationLock
func GetOperationLockError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeOperationLock,
		"Operation refused due to mutex lock on correlation ID or trip ID",
	)
}

// GetAccessForbiddenError creates an error related to AccessForbidden
func GetAccessForbiddenError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeAccessForbidden,
		"Operation failed due to access forbidden",
	)
}

// GetDataCorruptionError creates an error related to DataCorruption
func GetDataCorruptionError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeDataCorruption,
		"Operation failed due to internal storage data corruption",
	)
}

// GetNotImplementedError creates an error related to NotImplemented
func GetNotImplementedError(innerErrors ...error) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeNotImplemented,
		"Operation failed due to internal business logic not implemented",
	)
}

// GetCustomError creates a customized error with given code and formatted message
func GetCustomError(errorCode Code, messageFormat string, parameters ...interface{}) AppError {
	return &BaseAppError{
		fmtErrorf(messageFormat, parameters...),
		errorCode,
		nil,
		nil,
	}
}

func cleanupInnerErrors(innerErrors []error) []error {
	var cleanedInnerErrors = []error{}
	for _, innerError := range innerErrors {
		if innerError != nil {
			cleanedInnerErrors = append(
				cleanedInnerErrors,
				innerError,
			)
		}
	}
	return cleanedInnerErrors
}

// WrapError wraps an inner error with given message as a new error with given error code
func WrapError(innerErrors []error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
	var cleanedInnerErrors = cleanupInnerErrorsFunc(
		innerErrors,
	)
	if len(cleanedInnerErrors) == 0 {
		return nil
	}
	return &BaseAppError{
		fmtErrorf(messageFormat, parameters...),
		errorCode,
		cleanedInnerErrors,
		nil,
	}
}

// WrapSimpleError wraps an inner error with given message as a new general failure error
func WrapSimpleError(innerErrors []error, messageFormat string, parameters ...interface{}) AppError {
	return wrapErrorFunc(
		innerErrors,
		CodeGeneralFailure,
		messageFormat,
		parameters...,
	)
}

// GetInnermostErrors finds the innermost error wrapped within the given error
func GetInnermostErrors(err error) []error {
	var innermostErrors []error
	var baseAppError, ok = err.(AppError)
	if ok {
		for _, innerError := range baseAppError.InnerErrors() {
			innermostErrors = append(
				innermostErrors,
				GetInnermostErrors(innerError)...,
			)
		}
	} else {
		innermostErrors = append(
			innermostErrors,
			err,
		)
	}
	return innermostErrors
}
