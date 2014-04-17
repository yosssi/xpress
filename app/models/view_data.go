package models

type ViewData struct {
	App  *Application
	RCtx *RequestContext
	Data *interface{}
}

// NewViewData generates a ViewData and returns it.
func NewViewData(app *Application, rCtx *RequestContext, data *interface{}) *ViewData {
	return &ViewData{App: app, RCtx: rCtx, Data: data}
}
