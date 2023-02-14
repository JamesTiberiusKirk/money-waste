# ROADMAP
## V0.1.0
- Version 0.1 will be about having all the bells and whistles that anyone might ever want
- If the user does not want some of those features, they can just remove them manually 

### TODOs:
- [x] Link all of the libraries in "featuring" section
- [x] Actually setup the rootSitePath
- [x] Setup a good way to serve static content such as libraries
- [x] Setup a 404 default re-routing
- [x] Add support for SPAs (vue, angular, react, etc...) inside pages
  - [x] With support for letting the spa use its own router
  - [x] Need to implement some proxy system for being able to use the react dev server in development
- [ ] Figure out how to remove the exports from js dist
  - [x] Added a patch where export as a var is defined in the template import
- [ ] CLI stuff
  - [x] Finish a basic version of the cli with no options or interface
  - [ ] Package up the cli
  - [ ] Implement some form of a TUI around this to pick options and selections (Bubble Tea)
- [x] Make a detailed writeup about what this is meant to be 
  - [x] Including detailed READMEs on every feature
- [ ] Look into actually giving it a proper github url module name

## V0.2.0
- This will be about expanding on the feature set

### TODOs:
- [ ] Add support for web sockets
- [ ] Abstract the db behind a struct so the client can more easily swapped out
- [ ] Add an internal server failure page
- [ ] Add support for WASM inside pages
  - Maybe look at one of the WASM frameworks out there for go

---
# Backlog
## Template TODOs:
- [x] Link all of the libraries in "featuring" section
- [ ] Actually setup the rootSitePath
- [ ] Setup a good way to serve static content such as libraries
- [ ] Maybe think about setting up some theme or CSS library?
- [ ] Maybe find an actual way to load up the .env file in the dev_run.sh
- [x] Make some CLI or script for easy deployment of this Template
  - Not much needed other then copying the folder then using `sed -i` to replace the name to the desired one
- [x] Standalone API handler framework (like site but without the renderer)
- [x] Rethink the GetXMethod handler system
  - It adds too much boilerplate 
- [x] Maybe think of a way to hardcode in the path to be used outside each page for redirection
  - Probably just consts 
- [x] Add a menu system to frame and pages
- [ ] Add support for SPAs (vue, angular, react, etc...) inside pages
- [ ] Add support for WASM inside pages
  - Maybe look at one of the WASM frameworks out there for go
- [ ] Add support for web sockets
- [x] EITHER replace echoview or re-write it myself, it's been a year since last commit and its missing some needed features
  - Probably the best thing would be to implement it myself within my own site struct with complete support for stuff like:
    - [x] default 404s 
    - [x] options to exclude master frame per page 
    - [x] custom template file includes
    - [x] render function for the frame (so that base data can be included such as authed user, altho might still be better done in the GetPageHandler function instead)
- [x] Add typescript support to the js proto framework
  - [x] Go struct transpiration to typescript
  - [x] Typescript typings for the proto framework 
  - [ ] Try to implement the base framework functions in TS
  - [ ] Get rid of lint/lsp warning for unused declarations in TS
- [ ] Possibly throw the page status code in the get page data function so that each page can return their own status codes
- [ ] Add an internal server failure page
- [ ] Break out the template, spa, and static serving into their own sub structures maybe?

