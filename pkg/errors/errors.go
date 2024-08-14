package errors

var (
	ErrInternal         = Make("Hệ thống đang bận, vui lòng thử lại sau ít phút", 500)

	ErrForbidden      = Make("Access credentials are not sufficient to access this resource", 403)
	ErrInternalServer = Make("internal server error", 500)
	ErrUnauthorized   = Make("Access credentials are invalid", 401)

	ErrInvalidError = Make("Error code is invalid", 1001)

	ErrUUID = New("Wrong UUID format")


	// ErrSaveDb indicates error while saving to database
	ErrSaveDb = New("save entity to db error")
	// ErrUpdateDb indicates error while updating database
	ErrUpdateDb = New("update entity in db error")
	// ErrSelectDb indicates error while reading from database
	ErrSelectDb = New("select entity from db error")
	// ErrDeleteDb indicates error while deleting from database
	ErrDeleteDb = New("delete entity from db error")

	// ErrUnsupportedMediaType indicates unsupported media type
	ErrUnsupportedMediaType = New("unsupported media type")

	// ErrMalformedEntity indicates malformed entity specification (e.g)
	ErrMalformedEntity = Make("malformed entity specification", 2001)

	ErrBadRequest = Make("Invalid input format", 400)

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = Make("non-existent entity", 2003)

	// ErrConflict indicates that entity already exists.
	ErrConflict = Make("entity already exists", 2004)

	ErrIdentifyProject = Make("Failed to identify project", 2055)

	ErrInvalidId = Make("Invalid ID", 1001)

	ErrHttpTimeout = Make("Giao dịch vượt quá thời gian", 1000113)

	ErrInvalidSubmitData = Make("Submit data is invalid", 80001)

	// ErrMalformedEntity indicates malformed entity specification (e.g. invalid username or password).
	ErrUnauthorizedAccess = New("unauthorized")

)

// Error specifies an API that must be fullfiled by error type
type Error interface {

	// Error implements the error interface.
	Error() string

	// Msg returns error message
	Msg() string

	Code() int

	// Err returns wrapped error
	Err() Error
}

var _ Error = (*customError)(nil)

// customError struct represents a viot error
type customError struct {
	msg  string
	code int
	err  Error
}

func (ce *customError) Code() int {
	return ce.code
}

func (ce *customError) Msg() string {
	return ce.msg
}

func (ce *customError) Err() Error {
	return ce.err
}
func (ce *customError) Error() string {
	if ce == nil {
		return ""
	}
	if ce.err == nil {
		return ce.msg
	}
	return ce.msg + " : " + ce.err.Error()
}

// Contains inspects if e2 error is contained in any layer of e1 error
func Contains(e1 error, e2 error) bool {
	if e1 == nil || e2 == nil {
		return e2 == e1
	}
	ce, ok := e1.(Error)
	if ok {
		if ce.Msg() == e2.Error() {
			return true
		}
		return Contains(ce.Err(), e2)
	}
	return e1.Error() == e2.Error()
}

// Wrap returns an Error that wrap err with wrapper
func Wrap(wrapper error, err error) error {
	if wrapper == nil || err == nil {
		return wrapper
	}
	if w, ok := wrapper.(Error); ok {
		var code = w.Code()
		if code == 0 {
			code = 1
		}
		return &customError{
			msg:  w.Msg(),
			err:  cast(err),
			code: code,
		}
	}
	return &customError{
		msg: wrapper.Error(),
		err: cast(err),
	}
}

func cast(err error) Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(Error); ok {
		return e
	}
	return &customError{
		msg: err.Error(),
		err: nil,
	}
}

// New returns an Error that formats as the given text.
func New(text string) Error {
	return &customError{
		msg: text,
		err: nil,
	}
}

func Make(msg string, code int) Error {
	return &customError{
		msg:  msg,
		code: code,
		err:  nil,
	}
}

type ErrorBody struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
