// Package clienthelper provides utility functions for creating and handling API responses.

package clienthelper

import (
	"net/http"

	"github.com/princeparmar/go-helpers.git/context"
)

// APIExecutor defines an interface for executing API requests. It defines methods
// for executing the controller logic, and handling the response. Implementing this
// interface allows for modular and extensible API request handling, allowing developers
// to customize response handling as needed.
type APIExecutor interface {
	// Controller executes the business logic for the API request and returns the result
	// of the operation as an interface and any errors that occur during execution.
	// It takes in a context object and returns an interface and an error.
	Controller(ctx context.IContext) (interface{}, error)

	// HandleResponse takes in the result of the Controller method, any errors that occur,
	// the context object, and the APIResponse object to generate the HTTP response.
	// It takes in an interface, an error, a context object, an APIResponse pointer and
	// does not return anything.
	HandleResponse(ctx context.IContext, response interface{}, err error, apiResponse *APIResponse)
}

// RequestParser defines an interface for parsing the API request.
type RequestParser interface {
	// ParseRequest parses the HTTP request and extracts any relevant data into the
	// context object. It takes in the context, response writer, and HTTP request object
	// and returns any errors that occur during parsing.
	ParseRequest(ctx context.IContext, w http.ResponseWriter, r *http.Request) error
}

// RequestValidator defines an interface for validating the API request.
type RequestValidator interface {
	// ValidateRequest validates the data in the context object and returns any errors
	// that occur during validation. It takes in the context object and returns an error.
	ValidateRequest(ctx context.IContext) error
}

func getContextFromRequest(r *http.Request) context.IContext {
	ctx := r.Context()
	if v, ok := ctx.(context.IContext); ok {
		return v
	}

	return context.NewContextWithParent(ctx)
}

// GetHandler returns a new http.HandlerFunc that executes the provided APIExecutor.
// The returned function is responsible for parsing the request, validating it, invoking
// the controller and handling the response.
//
// The provided APIExecutor must implement the Controller method.
//
// If the APIExecutor also implements the RequestParser and/or RequestValidator interfaces,
// the request parsing and validation steps are also performed.
//
// If any of the above steps fail, an appropriate error response is returned to the client.
func GetHandler(factory func() APIExecutor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new context
		ctx := getContextFromRequest(r)

		executor := factory()

		// If the provided APIExecutor implements the RequestParser interface, call ParseRequest
		if parser, ok := executor.(RequestParser); ok {
			if err := parser.ParseRequest(ctx, w, r); err != nil {
				// Handle parse request error
				// This can be done by returning an error response or logging the error
				NewAPIResponse(w).SetStatusCode(http.StatusBadRequest).
					GetResponse().SetStatusCode(http.StatusBadRequest).
					GetErrors().SetClientMessage("request parsing failed").AddError(NewErrorFromErr(err))
				return
			}
		}

		// Check if the provided APIExecutor implements the RequestValidator interface
		if validator, ok := executor.(RequestValidator); ok {
			// Validate the request
			if err := validator.ValidateRequest(ctx); err != nil {
				// Handle validation error
				// Create a new APIResponse and set appropriate status code and error message
				// This can be done by returning an error response or logging the error
				NewAPIResponse(w).SetStatusCode(http.StatusBadRequest).
					GetResponse().SetStatusCode(http.StatusBadRequest).
					GetErrors().SetClientMessage("request validation failed").AddError(NewErrorFromErr(err))
				return
			}
		}

		// execute cotroller
		response, err := executor.Controller(ctx)

		// Handle the response
		executor.HandleResponse(ctx, response, err, NewAPIResponse(w))
	}
}

// baseAPIExecutor is an implementation of the APIExecutor interface that provides
// a default implementation of the Controller method and only implements the
// HandleResponse method.
type BaseAPIExecutor struct{}

// NewBaseAPIExecutor creates a new instance of baseAPIExecutor.
func NewBaseAPIExecutor() *BaseAPIExecutor {
	return &BaseAPIExecutor{}
}

// HandleResponse takes in the result of the Controller method, any errors that occur,
// and the APIResponse object to generate the HTTP response. It takes in an interface,
// an error, and an APIResponse pointer and does not return anything.
func (e *BaseAPIExecutor) HandleResponse(ctx context.IContext, response interface{}, err error, apiResponse *APIResponse) {
	if err != nil {
		// Set the error response
		apiResponse.SetStatusCode(http.StatusInternalServerError).
			GetResponse().SetStatusCode(http.StatusInternalServerError).
			GetErrors().SetClientMessage("internal server error").AddError(NewErrorFromErr(err))
	} else {
		// Set the success response
		apiResponse.SetStatusCode(http.StatusOK).
			GetResponse().SetStatusCode(http.StatusOK).
			SetData(response)
	}

	// Send the response
	if err := apiResponse.Send(); err != nil {
		// Handle error sending response
		// This can be done by logging the error
	}
}
