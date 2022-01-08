import * as catrina from "../widgets/catrina.js";

const SHOW_LABEL = "Show code";
let id = document.getElementById.bind(document);
let datas = [id('pre-radio'),id('pre-checkbox'), id('pre-file')];
catrina.CollapseBox(id("codebox15"), id('pre-checkbox').innerHTML, SHOW_LABEL, 'show-radio15');
catrina.CollapseBox(id("codebox16"), id('pre-radio').innerHTML, SHOW_LABEL, 'show-radio16');
catrina.CollapseBox(id("codebox17"), id('pre-file').innerHTML, SHOW_LABEL, 'show-radio17');
catrina.collapseBoxes_run();
datas.map(d => d.remove());
console.log("Catrina example");