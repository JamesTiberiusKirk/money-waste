# Go Web Template

This template is now feature complete. But more is coming.
If curious you can see my [roadmap here](./ROADMAP.md)

## The Cli
The Cli for this template can be found [here](https://github.com/JamesTiberiusKirk/go-template).  
I kept going around in circles trying to make a good cli which would fully work and would be properly packaged.
During my research I have stumbled across [go-template](https://github.com/SchwarzIT/go-template), which is a really nice template but more importantly comes with an extremely well made cli around it which frankly does everything I wanted todo and does it better, so in the spirit of open source....I've cannibalised it for my own template.  

Again the original CLI is really well made, so please have a look at the original repo!

## Featuring
- Mostly out of your way opinionated and extensible framework for web templating and api
- Echo framework for the HTTP server backbone
- Started off with [Goview](https://github.com/foolin/goview) to help with templating
  - Now a modified and stripped down version the source code is embedded into the renderer package 
- GORM with Postgres
- Session authentication
- [Reflex](github.com/cespare/reflex) for hotloading
- Forked [tygo](github.com/gzuidhof/tygo) and added a new feature for Typescript type generation from go structs
  - [Fork here](github.com/JamesTiberiusKirk/tygo)

## Template Architecture Features:
- [API](api/)
- [Site](site/)
  - [SPA hosting](site/spa/)
  - [SSR/Templating](site/page/)
    - [JS/TS templating framework with full stack type safety](site/page#jsts-minimal-ssr-framework)

### Caveat 
As of this version of this template, I have not setup any form of toggling for different template features.
Meaning that when you init a new project with this template you will be setup with all of it's features and sub-features (i.e. api, site, spa, ssr, etc).  
So at the moment, please refer to the sub docs for the feature in which you are trying to remove. There you will see a "Removal" section which will describe how to get rid of that specific feature.


# For getting the dev script to run 
```sh 
go install github.com/cespare/reflex@latest
go install github.com/JamesTiberiusKirk/tygo@v0.2.5
```
