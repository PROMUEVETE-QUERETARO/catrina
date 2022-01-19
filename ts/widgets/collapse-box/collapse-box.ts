import {HIDDEN_CLASS} from "../../catrina.js";

export function CollapseBox(node: HTMLElement, content: string, header:string, idSwitch: string) {
    node.innerHTML =
        `<div class="c__header_collapse"><label class="c__label" for="${idSwitch}">${header}</label>${__switchHTML(idSwitch)}</div>
        <div class="c__collapse_box ${HIDDEN_CLASS}">${content}</div>`;
}

function __switchHTML(id: string): string {
    return `<label class="c__switch" for="${id}">
    <input class="c__switch__input" id="${id}" type="checkbox">
    <span class="c__switch__slider"></span>
    </label>`
}

export function collapseBoxes_run() {
    let boxes = document.querySelectorAll(".c__collapse_box");
    let headers = document.querySelectorAll(".c__header_collapse");

    headers.forEach((h, i) => {
        let input = h.querySelector(".c__switch__input");
        if (input == null) {
            return
        }

        input.addEventListener("change", ()=>{

            //@ts-ignore
            if (input.checked) {
                boxes[i].classList.remove(HIDDEN_CLASS);
            } else {
                boxes[i].classList.add(HIDDEN_CLASS);
            }
        });
    });
}