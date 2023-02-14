# API
- The `Api` struct allows you to define api endpoints
- It holds any dependencies the handlers will need (db, clients, etc...)
- It holds the `sessionManager` for defining authenticated routes
- `routes` is for use for the site, it allows the templating engine to have a map of all of the defined routes
  - The following is passed to templating as part of the meta data
  - key: `route.RouteID`
  - value: full api server path

## Routes and SubRoutes
- Each route has fields for the different types of HTTP requests
  - `GET` `POST` `DELETE` `PUT`
  - Feel free to modify this list as you wish
- `RouteID`, this is for just internal use  
- `Path` is for defining the literal url path of the new route
  - Remember, that's concatenated on any previous route (if its a SubRoute)or root path for the api
  - Also, path can be anything that is accepted by Echo (i.e. "/:routeParam", or wild cards "/*")
- `Subroute` is a pointer to another `Route` allowing for route nesting
  - These could either be defined in the parent routes, or to more easily decouple dependencies just pass the sub route through the params to parent route
  - See example of this in users.go
- Examples for this can be found in `api/route/user.go` (sub route) and `api/route/users.go` (parent route)

# Removal instructions
- Remove this package
- Remove the appropriate code from main.go
