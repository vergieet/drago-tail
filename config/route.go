package config

type Route struct {
	list map[string]Routes
}



func (r *Route) add(routes Routes)  {
	r.list[routes.url] = routes
}

type Routes struct {
	method string
	url string
	run
}

type run func() string;