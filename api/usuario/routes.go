package usuario

// Routes struct
type Routes struct{}

// List -> url list endpoint
func (r *Routes) List() string {
	return "/list"
}

// Get -> url get endpoint
func (r *Routes) Get() string {
	return "/get"
}
