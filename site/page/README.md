# Pages package
- This package is all about managing templates and server side rendering
- By far this is the most complex thing here
- But just like anything else in this template, feel free to remove and cleanup anything you might not need
- There are two main struct that are used through this feature
  - `Page` for defining an actual page
  - `PageMetaData` is for defining data which will be passed to every single page on the site such as title, menu id and commonly used url variables

## Page fields
- `MenuID` aka a page id is injected into each page meta PageMetaData
  - Also used in the routes map which is accessible in all of the templates so you don't need to hard code any urls in the templates 
  - Keep it camel case if u want to be able to use the [routes](#templating) value in the template without needing to use the `index` function 
- `Frame` used to define whether the page is to be encapsulated by the normal page frame defined in `Site`
- `Template` is to define the file used as the main template of the page
  -  File path will be relative to the folder containing the templates defined in  the `Site.buildRenderer` function
- `Deps` is to hold any page specific dependencies
- `GetPageData` is a function which will get any page data needed then have it embedded in the page templating system
- Handlers for POST, DELETE, and PUT are also present if there is a need to add those handlers for a specific page.
  - E.G. if there is a form on the page, etc...


## Templating
- Renderer built in functions:
  - `include` this is to just include extra templates
  - `includeJs` use this to include a js file, it will automatically wrap script tags around the file contents
  - `includeTs` same as the above, however it prepends the output folder for any compiled typescript and replaces `.ts` with `.js`
- Template variables:
  - `data` will be whatever the `Page.GetPageData` function returns for each page
  - `meta` is generated by `Page.buildBasePageMetaData` function and it returns an instance of `PageMetaData`
  - `auth` will either be empty or will contain the email and username of the currently logged in user
  - `routes` will be `map[string]map[string]`
    - Allowing you to write something similar to the following `{{.routes.site.userPage}}` and get the url that is defined for that page.
    - This also works with the api package, i.e. `{{.routes.api.helloWorld}}`

![routes_struct](https://user-images.githubusercontent.com/17408117/216288313-102ba524-2a1b-497e-9224-0c66ba1de599.png)

# JS/TS minimal ssr framework
## Yet another JS framework?
It seems like I might be re-inventing the wheel here.... but theres a good reason.
I strongly believe that JS is one of the reasons why the web actually sucks nowadays. Which is why it seems like I'm taking a step back in time with using templating. However, we literally have Next.js and other modernly abstracted systems.  
This is me taking a step back in time with some lessons we have learned from the modern web.  
These lessons are:  
- We still could really use some scripting on the browser
- Scripting of any kind SUCKS without types
- Abstraction is BS and makes the web not only slow but also harder debug

## How does it work?
Because of this I have decided to make a bare bones framework which would be designed specifically to accept the `echo.Map` data from the templating system. That data would be injected into some js vars inside the template.
See bellow screenshot from `./framework/framework.gohtml`.   
![image](https://user-images.githubusercontent.com/17408117/216392675-0c33b533-3756-44b8-a86b-13f681e2414f.png)

From here we can already do a bunch of cool stuff allowing us to perform some sort of rendering without having to make any more requests to the server.  

To build on top of that I have made a simple function which is available on every page which allows you to recursively make element trees where needed. Check out the `ts` binding bellow which might explain it.

```typescript 
function elem(
  elem: html,
  inner: HTMLElement | string | number,
  options?: ElemOptions
): HTMLElement
```

All the user needs to do is create a `div` with the id of `page_render` then in JS define a function with the following definition.

```javascript
function render(_data:any): HTMLElement
```
And return a HTML element using the `elem` function.

Any form of reactivity and re-rendering is actually done quite simply. When defining a html element, give it an id (can be done through the options param) and just use the `document.getElementById` api to perform any dom manipulation necessary.  
**Simple.**

## How does TypeScipt come into all of this mess?
Well, there are typescript bindings for all of the different variables, types and functions which exist. See `./ts/global.d.ts` file.
- Bindings for all of the different functions and variables have been made in `./ts/global.d.ts`
- A node package has been made and setup to transpile the ts to js 
- The use of [Tygo](https://github.com/JamesTiberiusKirk/tygo), or at least my version of it to get some resembleness of full stack type safety
  - In the current setup Tygo looks in the `page` package for any structs which have the comment of `\\ TsType` 
  - It then transpiles them to typescript types, and dumps them to the `./tstypes` folder made available to be used in typescript
  - ...If the page data is defined once, why bother doing it again manually in typescript?

## Some reflection
There is definitely some work to be done here both on the typescript implementation and on javascript in general.
However, the whole point of this is that it can be built up as much or as little by the dev using this template.

## Removal instruction
Wanna get rid of this neurodivergent code? I don't blame you, just follow the following.
- Remove the inclusion of the framework in the frame templates.
- Remove the following folders:
  - framework
  - jsdist
  - ts 
  - tstypes
  - node_modules if it exists
- Remove any example use of the framework through out the existing pages 
- Remove the `includeTs` and `includeJs` (if you think you wont be using external js files) template funcs from the `./render/render.go` file
- Remove package-lock.json, package.json and tygo.yaml



