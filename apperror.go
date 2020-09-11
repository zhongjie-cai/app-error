package apperror

// AppError is the error wrapper interface for all WebServiceTemplate service generated errors
type AppError interface {
	// Golang internal error interface
	error
	// PrintError prints given error data to a string; override this if you want a different format than default style as "(Code) Message [Attached Data]"
	PrintError(code Code, err error, extraData map[string]interface{}) string
	// Code returns the string representation of the error code enum
	Code() string
	// HTTPStatusCode returns the corresponding HTTP status code mapped to the error code value
	HTTPStatusCode() int
	// Contains checks if the current error object or any of its inner errors contains the given error object
	Contains(err error) bool
	// Wrap wraps the given list of inner errors into the current app error object
	Wrap(innerErrors ...error)
	// Attach adds the given value to the current app error's extra data map by given name
	Attach(name string, value interface{})
}

// These are print formatting related constants
const (
	errorMessageFormat   string = "(%v) %v%v" // (Code) Message [extra data]
	errorExtraDataFormat string = "%v = %+v"  // name = value
	errorJoiningFormat   string = " [ %v ]"   // [ content ]
	errorPointer         string = " : "
	errorSeparator       string = " | "
)

// BaseAppError instantiates the AppError interface and provides a base for inheritance
type BaseAppError struct {
	error
	code        Code
	innerErrors []error
	extraData   map[string]interface{}
}

// NewBaseAppError creates an instance of BaseAppError object using given data
func NewBaseAppError(code Code, messageFormat string, parameters ...interface{}) *BaseAppError {
	return &BaseAppError{
		fmtErrorf(
			messageFormat,
			parameters...,
		),
		code,
		[]error{},
		map[string]interface{}{},
	}
}

func formatExtraData(extraData map[string]interface{}) string {
	if len(extraData) == 0 {
		return ""
	}
	var extraDataMessages = []string{}
	for name, value := range extraData {
		extraDataMessages = append(
			extraDataMessages,
			fmtSprintf(
				errorExtraDataFormat,
				name,
				value,
			),
		)
	}
	return fmtSprintf(
		errorJoiningFormat,
		stringsJoin(
			extraDataMessages,
			errorSeparator,
		),
	)
}

// PrintError prints given error data to a string; override this if you want a different format than default style as "(Code) Message [Attached Data]"
func (baseAppError *BaseAppError) PrintError(code Code, err error, extraData map[string]interface{}) string {
	var extraDataMessage = formatExtraDataFunc(
		extraData,
	)
	return fmtSprintf(
		errorMessageFormat,
		code,
		err,
		extraDataMessage,
	)
}

func printBaseAppError(baseAppError *BaseAppError) string {
	return baseAppError.PrintError(
		baseAppError.code,
		baseAppError.error,
		baseAppError.extraData,
	)
}

func getErrorMessage(err error) string {
	return err.Error()
}

func printInnerErrors(innerErrors []error) string {
	if len(innerErrors) == 0 {
		return ""
	}
	var innerErrorMessages = []string{}
	for _, innerError := range innerErrors {
		if innerError != nil {
			innerErrorMessages = append(
				innerErrorMessages,
				getErrorMessageFunc(
					innerError,
				),
			)
		}
	}
	return fmtSprintf(
		errorJoiningFormat,
		stringsJoin(
			innerErrorMessages,
			errorSeparator,
		),
	)
}

func (baseAppError *BaseAppError) Error() string {
	var baseErrorMessage = printBaseAppErrorFunc(
		baseAppError,
	)
	var innerErrorMessage = printInnerErrorsFunc(
		baseAppError.innerErrors,
	)
	return fmtSprint(
		baseErrorMessage,
		innerErrorMessage,
	)
}

// Code returns string representation of the error code of the app error
func (baseAppError *BaseAppError) Code() string {
	return baseAppError.code.String()
}

// HTTPStatusCode returns HTTP status code according to the error code of the app error
func (baseAppError *BaseAppError) HTTPStatusCode() int {
	return baseAppError.code.HTTPStatusCode()
}

func equalsError(err, target error) bool {
	return err == target ||
		err.Error() == target.Error() ||
		errorsIs(err, target)
}

func appErrorContains(appError AppError, err error) bool {
	return appError.Contains(err)
}

func innerErrorContains(innerErrors []error, err error) bool {
	for _, innerError := range innerErrors {
		var typedError, isTyped = innerError.(AppError)
		if isTyped {
			if appErrorContainsFunc(
				typedError,
				err,
			) {
				return true
			}
		} else if equalsErrorFunc(innerError, err) {
			return true
		}
	}
	return false
}

// Contains checks if the current error object or any of its inner errors contains the given error object
func (baseAppError *BaseAppError) Contains(err error) bool {
	if baseAppError == err ||
		equalsErrorFunc(baseAppError.error, err) {
		return true
	}
	return innerErrorContainsFunc(
		baseAppError.innerErrors,
		err,
	)
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

// Wrap wraps the given list of inner errors into the current app error object
func (baseAppError *BaseAppError) Wrap(innerErrors ...error) {
	var cleanedInnerErrors = cleanupInnerErrorsFunc(
		innerErrors,
	)
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
	var baseAppError = newBaseAppErrorFunc(
		CodeGeneralFailure,
		"An error occurred during execution",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetUnauthorized creates an error related to Unauthorized
func GetUnauthorized(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeUnauthorized,
		"Access denied due to authorization error",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetInvalidOperation creates an error related to InvalidOperation
func GetInvalidOperation(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeInvalidOperation,
		"Operation (method) not allowed",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetBadRequestError creates an error related to BadRequest
func GetBadRequestError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeBadRequest,
		"Request URI or body is invalid",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetNotFoundError creates an error related to NotFound
func GetNotFoundError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeNotFound,
		"Requested resource is not found in the storage",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetCircuitBreakError creates an error related to CircuitBreak
func GetCircuitBreakError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeCircuitBreak,
		"Operation refused due to internal circuit break on correlation ID",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetOperationLockError creates an error related to OperationLock
func GetOperationLockError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeOperationLock,
		"Operation refused due to mutex lock on correlation ID or trip ID",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetAccessForbiddenError creates an error related to AccessForbidden
func GetAccessForbiddenError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeAccessForbidden,
		"Operation failed due to access forbidden",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetDataCorruptionError creates an error related to DataCorruption
func GetDataCorruptionError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeDataCorruption,
		"Operation failed due to internal storage data corruption",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}

// GetNotImplementedError creates an error related to NotImplemented
func GetNotImplementedError(innerErrors ...error) AppError {
	var baseAppError = newBaseAppErrorFunc(
		CodeNotImplemented,
		"Operation failed due to internal business logic not implemented",
	)
	baseAppError.Wrap(
		innerErrors...,
	)
	return baseAppError
}
