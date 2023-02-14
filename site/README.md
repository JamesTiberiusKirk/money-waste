# Site
- `Site` is for defining a site with one of the following features
  - Standard HTML and JS templating
    - With a lot of nice bells and whistles (refer tho the templating and SSR docs)
  - Static serving of files
  - Single page app setup (SPA with something like React, Vue, etc...)

## Site fields
- `dev` proxying connections to dev servers with an SPA
- `frameTmpls` is for defining a frame templates
  - There are multiple for being able to toggle them from the page
  - In this state it is used to either have a frame (header, footer, etc)
  - Or have an empty frame.
  - This system could be modified to allow for multiple custom frames.
 
## Sub-features
- [SPA](spa/README.md)
- [Static serving](#static-serving)
- [SSR/Templating](page/README.md)

# Removal instructions
- Remove the entire `Site` package
- Remove `./package.json`
- Remove `./tsconfig.json`
- Remove `./tygo.yaml`
- Feel free to remove the `routes` map and `RoutesID` from the `Api` and `Routes` packages
- Remove the appropriate initialisation code from `./main.go`

# Static serving
- Not much to say there, add the static folders that you want served in the `staticFolders` map
  - key: path to be served at 
  - value: the folder to be served (from project root)

## Removal instructions
- Remove `Site.staticFolders` field
- Remove any initialisations of said field
- Remove any static assets folders 
- Remove `Site.mapStatic` folder
