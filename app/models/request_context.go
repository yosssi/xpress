package models

// A RequestContext represents a request context.
type RequestContext struct {
	User *User
}

// NewRequestContext generages a RequestContext and returns it.
func NewRequestContext() *RequestContext {
	return &RequestContext{}
}
