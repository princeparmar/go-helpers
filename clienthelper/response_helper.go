package clienthelper

// Response struct to be used in client response
type Response struct {
	StatusCode int         `json:"status_code"`
	Errors     *Errors     `json:"errors"`
	Data       interface{} `json:"data"`
}

// NewResponse is a factory function to create a new Response struct
func NewResponse() *Response {
	return &Response{}
}

// SetStatusCode sets the status code in the Response struct
func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

// GetStatusCode returns the status code from the Response struct
func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

// SetErrors sets the Errors struct in the Response struct
func (r *Response) SetErrors(errors *Errors) *Response {
	r.Errors = errors
	return r
}

// GetErrors returns the Errors struct from the Response struct
func (r *Response) GetErrors() *Errors {
	return r.Errors
}

// SetData sets the data in the Response struct
func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

// GetData returns the data from the Response struct
func (r *Response) GetData() interface{} {
	return r.Data
}

// Errors struct for client response
type Errors struct {
	ClientMessage interface{} `json:"client_message"`
	Err           []*Error    `json:"err"`
}

// NewErrors is a factory function to create a new Errors struct
func NewErrors() *Errors {
	return &Errors{}
}

// SetClientMessage sets the client message in the Errors struct
func (e *Errors) SetClientMessage(clientMessage interface{}) *Errors {
	e.ClientMessage = clientMessage
	return e
}

// AddError adds an error to the Errors struct
func (e *Errors) AddError(err *Error) *Errors {
	e.Err = append(e.Err, err)
	return e
}

// GetClientMessage returns the client message from the Errors struct
func (e *Errors) GetClientMessage() interface{} {
	return e.ClientMessage
}

// GetErrors returns the errors from the Errors struct
func (e *Errors) GetErrors() []*Error {
	return e.Err
}

// ErrorLevel for client response
type ErrorLevel string

// Constants for different error levels
const (
	Level_Error   ErrorLevel = "ERROR"
	Level_Warning ErrorLevel = "WARNING"
)

// Error struct for client response
type Error struct {
	Message     string     `json:"message"`
	MessageCode string     `json:"message_code"`
	Level       ErrorLevel `json:"level"`
}

// NewError is a factory function to create a new Error struct
func NewError() *Error {
	return &Error{}
}

// NewError is a factory function to create a new Error struct
func NewErrorFromErr(err error) *Error {
	return &Error{
		Message: err.Error(),
		Level:   Level_Error,
	}
}

// SetMessage sets the error message in the Error struct
func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

// GetMessage returns the error message from the Error struct
func (e *Error) GetMessage() string {
	return e.Message
}

// SetMessageCode sets the error message code in the Error struct
func (e *Error) SetMessageCode(code string) *Error {
	e.MessageCode = code
	return e
}

// GetMessageCode returns the error message code from the Error struct
func (e *Error) GetMessageCode() string {
	return e.MessageCode
}

// SetLevel sets the error level in the Error struct
func (e *Error) SetLevel(level ErrorLevel) *Error {
	e.Level = level
	return e
}

// GetLevel returns the error level from the Error struct
func (e *Error) GetLevel() ErrorLevel {
	return e.Level
}
