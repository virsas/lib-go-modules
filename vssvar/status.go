package vssvar

import "net/http"

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

const STATUS_OK_STRING = "OK"
const STATUS_OK_CODE = http.StatusOK
const STATUS_SERVER_ERROR_STRING = "apiError"
const STATUS_SERVER_ERROR_CODE = http.StatusInternalServerError
const STATUS_DB_ERROR_STRING = "dbError"
const STATUS_DB_ERROR_CODE = http.StatusInternalServerError
const STATUS_EMAIL_ERROR_STRING = "emailError"
const STATUS_EMAIL_ERROR_CODE = http.StatusInternalServerError
const STATUS_FORBIDDEN_ERROR_STRING = "notAllowed"
const STATUS_FORBIDDEN_ERROR_CODE = http.StatusForbidden
const STATUS_EXPIRED_ERROR_STRING = "expiredError"
const STATUS_EXPIRED_ERROR_CODE = http.StatusForbidden
const STATUS_LOGIN_ERROR_STRING = "loginError"
const STATUS_LOGIN_ERROR_CODE = http.StatusUnauthorized
const STATUS_REQUEST_ERROR_STRING = "badRequest"
const STATUS_REQUEST_ERROR_CODE = http.StatusBadRequest
const STATUS_VALIDATION_ERROR_STRING = "validationError"
const STATUS_VALIDATION_ERROR_CODE = http.StatusBadRequest
const STATUS_NOTFOUND_ERROR_STRING = "notFound"
const STATUS_NOTFOUND_ERROR_CODE = http.StatusNotFound
const STATUS_DUPLICATION_ERROR_STRING = "duplicationError"
const STATUS_DUPLICATION_ERROR_CODE = http.StatusConflict
const STATUS_CONFLICT_ERROR_STRING = "conflictError"
const STATUS_CONFLICT_ERROR_CODE = http.StatusConflict
const STATUS_EMPTY_ERROR_STRING = "notEmpty"
const STATUS_EMPTY_ERROR_CODE = http.StatusConflict

func StatusOK(err error) error {
	return &StatusErr{STATUS_OK_CODE, STATUS_OK_STRING, err}
}
func StatusServerError(err error) error {
	return &StatusErr{STATUS_SERVER_ERROR_CODE, STATUS_SERVER_ERROR_STRING, err}
}
func StatusDBError(err error) error {
	return &StatusErr{STATUS_DB_ERROR_CODE, STATUS_DB_ERROR_STRING, err}
}
func StatusEmailError(err error) error {
	return &StatusErr{STATUS_EMAIL_ERROR_CODE, STATUS_EMAIL_ERROR_STRING, err}
}
func StatusForbiddenError(err error) error {
	return &StatusErr{STATUS_FORBIDDEN_ERROR_CODE, STATUS_FORBIDDEN_ERROR_STRING, err}
}
func StatusLoginError(err error) error {
	return &StatusErr{STATUS_LOGIN_ERROR_CODE, STATUS_LOGIN_ERROR_STRING, err}
}
func StatusExpiredError(err error) error {
	return &StatusErr{STATUS_EXPIRED_ERROR_CODE, STATUS_EXPIRED_ERROR_STRING, err}
}
func StatusRequestError(err error) error {
	return &StatusErr{STATUS_REQUEST_ERROR_CODE, STATUS_REQUEST_ERROR_STRING, err}
}
func StatusValidationError(err error) error {
	return &StatusErr{STATUS_VALIDATION_ERROR_CODE, STATUS_VALIDATION_ERROR_STRING, err}
}
func StatusNotFoundError(err error) error {
	return &StatusErr{STATUS_NOTFOUND_ERROR_CODE, STATUS_NOTFOUND_ERROR_STRING, err}
}
func StatusDuplicationError(err error) error {
	return &StatusErr{STATUS_DUPLICATION_ERROR_CODE, STATUS_DUPLICATION_ERROR_STRING, err}
}
func StatusConflictError(err error) error {
	return &StatusErr{STATUS_CONFLICT_ERROR_CODE, STATUS_CONFLICT_ERROR_STRING, err}
}
func StatusEmptyError(err error) error {
	return &StatusErr{STATUS_EMPTY_ERROR_CODE, STATUS_EMPTY_ERROR_STRING, err}
}
