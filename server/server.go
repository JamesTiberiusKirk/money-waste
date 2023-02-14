package server

type RoutesMap map[string]string

type Server interface {
	Serve()
	GetRoutes() RoutesMap
	SetRoutes(t string, r RoutesMap)
}
