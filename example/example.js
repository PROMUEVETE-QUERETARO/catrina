import * as catrina from "../widgets/catrina.js";

const SHOW_LABEL = "Show code";

window.onload = () => {
    console.log("Running Catrina example...");
    let id = document.getElementById.bind(document);

    catrina.CollapseBox(id("codebox1"), ExtractCode('#buttons-container'), SHOW_LABEL, 'show-buttons');
    catrina.CollapseBox(id("codebox15"), ExtractCode('#checkbox-container'), SHOW_LABEL, 'show-checkbox');
    catrina.CollapseBox(id("codebox16"), ExtractCode("#radio-container"), SHOW_LABEL, 'show-radio');
    catrina.CollapseBox(id("codebox17"), ExtractCode('#file-container'), SHOW_LABEL, 'show-file');
    catrina.CollapseBox(id("codebox18"), ExtractCode('#table-container'), SHOW_LABEL, 'show-table');

    catrina.collapseBoxes_run();

    let button = catrina.PrimaryButton('', 'Accept'),
        button2 = catrina.DangerousButton('', 'Refuse');

    id('run-alert').onclick = ()=> catrina.Alert({title: 'Alert Widget', content: HtmlFromTemplate('#alert1-content'), buttons: []});
    id('run-alert-no-skip').onclick = ()=>{
        let alert = catrina.Alert({title: 'Non-skipped Alert', content: HtmlFromTemplate('#alert2-content'), buttons: [button, button2], skip: false})
        button.onclick = ()=> {
            catrina.FloatNotify(catrina.notifyType.Good, 'Accept button clicked', 3000);
            alert.content.remove()
            alert.shadow.remove()
        }
        button2.onclick = ()=> {
            catrina.FloatNotify(catrina.notifyType.Error, 'Refuse button clicked', 3000);
            alert.content.remove()
            alert.shadow.remove()
        }
    }

    id('loader-btn').onclick = () => {
        let loader = catrina.LoaderStart();
        catrina.LoaderStop(10_000, loader);
    }

    id('modal-btn').onclick = () => {
        let modal = catrina.Modal('Modal Window', 'c__button _no_shadows');
        modal.body.appendChild(id('modal-content').content.cloneNode(true));
        modal.btn.addEventListener('click', ()=>{
            catrina.CloseModal(modal.modal);
            catrina.FloatNotify(catrina.notifyType.Good, 'Modal Closed', 3000)

        });
    }

}

function n_space(n){
    let str = '';
    for (let i = 0; i <= n; i++) {
        str += ' '
    }
    return str
}

// inspired in https://stackoverflow.com/questions/14129953/how-to-encode-a-string-in-javascript-for-displaying-in-html
function ExtractCode(selector) {
    let div = document.createElement("div");
    let pre = document.createElement("pre");
    div.appendChild(pre)
    let node = document.querySelector(selector);
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

function HtmlFromTemplate(selector) {
    let div = document.createElement('div');
    let content = document.querySelector(selector).content.cloneNode(true);
    if (content != null) {
        div.appendChild(content);
    }

    return div.innerHTML;
}