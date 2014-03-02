package models

type ViewData struct {
	App  *Application
	Data *interface{}
}

// NewViewData generates a ViewData and returns it.
func NewViewData(app *Application, data *interface{}) *ViewData {
	return &ViewData{App: app, Data: data}
}
