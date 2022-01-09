import * as catrina from "../widgets/catrina.js";

console.log("Running Catrina example...");

const SHOW_LABEL = "Show code";
let id = document.getElementById.bind(document);
catrina.CollapseBox(id("codebox15"), ExtractCode('#checkbox-container'), SHOW_LABEL, 'show-checkbox');
catrina.CollapseBox(id("codebox16"), ExtractCode("#radio-container"), SHOW_LABEL, 'show-radio');
catrina.CollapseBox(id("codebox17"), ExtractCode('#file-container'), SHOW_LABEL, 'show-file');

catrina.collapseBoxes_run();


function n_space(n){
    let str = '';
    for (let i = 0; i <= n; i++) {
        str += ' '
    }
    return str
}

// inspired in https://stackoverflow.com/questions/14129953/how-to-encode-a-string-in-javascript-for-displaying-in-html
function ExtractCode(keyof) {
    let div = document.createElement("div");
    let pre = document.createElement("pre");
    div.appendChild(pre)
    let node = document.querySelector(keyof);
    let content = "";
    if (node != null){
        content = node.innerHTML
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;');


        let arr = content.split('\n');
        let N = [];
        let numbers = [];
        for (let i = 0; i < arr.length; i++) {
            let regex = /\s\s\s/g;
            let n = ((arr[i]||'').match(regex)||[]).length;
            n > 0 ? N.push(n) : true
            numbers.push(n)
        }

        N.sort((a, b)=> {return a-b})

        let n = N[0]
        for (let i = 0; i < arr.length; i++) {
            let spaces = numbers[i] - n;
            if (spaces <= 0) continue;
            arr[i] = `${n_space(spaces)}${arr[i].replace(/(\s*)&lt;/g, '&lt;')}`;
        }
        content = arr.join('\n')

    }
    pre.innerHTML = content;
    return div.innerHTML
}