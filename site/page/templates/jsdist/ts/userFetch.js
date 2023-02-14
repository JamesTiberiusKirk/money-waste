"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// TODO: Figure out how to get rid of the unused declaration
function render(_data) {
    let styles = {
        border: "solid 2px red",
        color: "blue",
    };
    return elem(html.DIV, elem(html.H2, `This is a hello world from ts ${_data.user.email}`, { style: styles }));
}
