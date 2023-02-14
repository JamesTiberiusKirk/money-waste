
// all the HTML elements listed for easier use with the elem() func
const html = {
  H1: "h1",
  H2: "h2",
  H3: "h3",
  H4: "h4",
  H5: "h5",
  H6: "h6",
  H7: "h7",
  DIV: "div",
  P: "p",
  A: "a",
  FORM: "form",
  INPUT: "input",
  BUTTON: "button",
}

// TODO: probably needs to be removed as its not actually used
// This is for adding .format() function on any string
// This pretty much is used as a printf function
// EXAMPLE USE: "Some test string {0} {1}".format("one var", "second var")
// OUTPUT: "Some test string one var second var"
// if (!String.prototype.format) {
//   String.prototype.format = function(...args) {
//     return this.replace(/(\{\d+\})/g, function(a) {
//       return args[+(a.substr(1, a.length - 2)) || 0];
//     });
//   };
// }

function assignOptionsToElem(options = {}, element) {
  for (const key in options) {
    if (typeof options[key] === "object" && key in element) {
      console.log("nested")
      element[key] = assignOptionsToElem(options[key], element[key])
    } else if (typeof options[key] === "string" && key in element) {
      element[key] = options[key]
    }
  }

  return element
}

function innerTextById(id, text) {
  const elem = document.getElementById(id)
  elem.innerText = text
}

// elem - string - is the enum html element type
// inner - HTML element or string- is the inner complete html string element
// options - style field inside of type CSSStyleDeclaration
// returns HTML element - allowing for easier concatination of multiple of this function
function elem(elem, inner, options = {}) {
  const { style } = options

  element = document.createElement(elem)

  if (typeof inner === "string" || typeof inner === "number") {
    element.innerText = inner
  } else {
    // Assuming that if its not a string it will be a HTML element
    if (Array.isArray(inner)) {
      inner.forEach(child => {
        element.appendChild(child)
      })
    } else {
      element.appendChild(inner)
    }
  }

  // Have to do this bs to actually be able to assign the styles over
  // Might have to do the same with the rest of the attribs
  // This does not work for objects so needs to be manually done for no
  for (const key in style) {
    element.style[key] = style[key]
  }

  // This does not work for nested objects
  for (const key in options) {
    if (typeof options[key] === "object") continue
    element[key] = options[key]
  }

  return element
}

contentDiv = document.getElementById("page_render")
if (typeof contentDiv !== 'undefined' && typeof render !== 'undefined') {
  contentDiv.appendChild(render(_data))
}
