package errs

import "net/http"

type errMessage struct {
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
	ErrMessage string `json:"message"`
}

type ErrMessage interface {
	Status() int
	Error() string
	Message() string
}

func (e *errMessage) Status() int {
	return e.ErrStatus
}

func (e *errMessage) Error() string {
	return e.ErrMessage
}

func (e *errMessage) Message() string {
	return e.ErrMessage
}

func NewInternalServerError(message string) ErrMessage {
	return &errMessage{
		ErrStatus:  500,
		ErrError:   "INTERNAL_SERVER_ERROR",
		ErrMessage: message,
	}
}

func NewUnauthenticatedError(message string) ErrMessage {
	return &errMessage{
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "NOT_AUTHENTICATED",
		ErrMessage: message,
	}
}

func NewUnauthorizedError(message string) ErrMessage {
	return &errMessage{
		ErrStatus:  http.StatusForbidden,
		ErrError:   "NOT_AUTHORIZED",
		ErrMessage: message,
	}
}

func NewBadRequestError(message string) ErrMessage {
	return &errMessage{
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
		ErrMessage: message,
	}
}

func NewUnprocessableEntityError(message string) ErrMessage {
	return &errMessage{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "UNPROCESSABLE_ENTITY",
		ErrMessage: message,
	}
}
