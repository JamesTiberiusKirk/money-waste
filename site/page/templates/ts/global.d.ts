export { }

declare global {
  // TODO: need to see if there is a way to not doubly define this enum
  enum html {
    H1 = "h1",
    H2 = "h2",
    H3 = "h3",
    H4 = "h4",
    H5 = "h5",
    H6 = "h6",
    H7 = "h7",
    DIV = "div",
    P = "p",
    A = "a",
    FORM = "form",
    INPUT = "input",
    BUTTON = "button",
  }

  interface Data {
  }

  interface Meta {
    MenuID: string
    Title: string
    UrlError: string
    Success: string
  }

  interface Routes {
    site: Map<string, string>
  }

  interface Auth {
    email?: string
    username?: string
  }

  interface ElemOptions{
    style: CSSStyleDeclaration
  }

  let _data: Data
  let _auth: Auth
  let _meta: Meta
  let _rotues: Routes

  function render(_data:any): HTMLElement
  function elem(elem: html, inner: HTMLElement | string | number, options?: ElemOptions): HTMLElement
  function innerTextById(id: string, text: string): void
}

