package vars

type StatusErr struct {
	statusCode int
	statusMsg  string
	err        error
}

func (r *StatusErr) Error() string {
	return r.statusMsg
}
func (r *StatusErr) Code() int {
	return r.statusCode
}
func (r *StatusErr) Unwrap() error {
	return r.err
}

// Good request
const STATUS_OK_CODE = 200
const STATUS_OK_STRING = "OK"

func StatusOK(err error) error {
	return &StatusErr{STATUS_OK_CODE, STATUS_OK_STRING, err}
}

// Server side issue
const STATUS_SERVER_ERROR_CODE = 500
const STATUS_SERVER_ERROR_STRING = "apiError"

func StatusServerError(err error) error {
	return &StatusErr{STATUS_SERVER_ERROR_CODE, STATUS_SERVER_ERROR_STRING, err}
}

const STATUS_DB_ERROR_CODE = 500
const STATUS_DB_ERROR_STRING = "dbError"

func StatusDBError(err error) error {
	return &StatusErr{STATUS_DB_ERROR_CODE, STATUS_DB_ERROR_STRING, err}
}

const STATUS_LOGIN_ERROR_CODE = 401
const STATUS_LOGIN_ERROR_STRING = "loginError"

func StatusLoginError(err error) error {
	return &StatusErr{STATUS_LOGIN_ERROR_CODE, STATUS_LOGIN_ERROR_STRING, err}
}

const STATUS_REQUEST_ERROR_CODE = 400
const STATUS_REQUEST_ERROR_STRING = "badRequest"

func StatusRequestError(err error) error {
	return &StatusErr{STATUS_REQUEST_ERROR_CODE, STATUS_REQUEST_ERROR_STRING, err}
}

const STATUS_VALIDATION_ERROR_CODE = 400
const STATUS_VALIDATION_ERROR_STRING = "validationError"

func StatusValidationError(err error) error {
	return &StatusErr{STATUS_VALIDATION_ERROR_CODE, STATUS_VALIDATION_ERROR_STRING, err}
}

const STATUS_NOTFOUND_ERROR_CODE = 404
const STATUS_NOTFOUND_ERROR_STRING = "notFound"

func StatusNotFoundError(err error) error {
	return &StatusErr{STATUS_NOTFOUND_ERROR_CODE, STATUS_NOTFOUND_ERROR_STRING, err}
}

const STATUS_DUPLICATION_ERROR_CODE = 409
const STATUS_DUPLICATION_ERROR_STRING = "duplicationError"

func StatusDuplicationError(err error) error {
	return &StatusErr{STATUS_DUPLICATION_ERROR_CODE, STATUS_DUPLICATION_ERROR_STRING, err}
}

const STATUS_EXPIRED_ERROR_CODE = 403
const STATUS_EXPIRED_ERROR_STRING = "expiredError"

func StatusExpiredError(err error) error {
	return &StatusErr{STATUS_EXPIRED_ERROR_CODE, STATUS_EXPIRED_ERROR_STRING, err}
}
