package dto

import "fmt"

// UnauthorizedError represents an error when user tries to access
// guarded resources without autohrization information
type UnauthorizedError struct {
}

func (e UnauthorizedError) Error() string {
	return "user is not authorized"
}

// ForbiddenError represents an error when authorized user tries to access
// resources that he has no access to
type ForbiddenError struct {
	ResourceType string
	ResourceID   uint
}

func (e ForbiddenError) Error() string {
	return fmt.Sprintf("access to %s#%d is denied", e.ResourceType, e.ResourceID)
}

// NotFoundError represents an error when user tries to access
// resources that does not exist
type NotFoundError struct {
	ResourceType string
	ResourceID   uint
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("resource %s#%d not found", e.ResourceType, e.ResourceID)
}

// BadRequestError represents a general bad request error with custom message
type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

// WrongRequestParametersError represents an error when user provides wrong
// parameters in a request
type WrongRequestParametersError struct {
}

func (e WrongRequestParametersError) Error() string {
	return fmt.Sprintf("invalid request parameters")
}
