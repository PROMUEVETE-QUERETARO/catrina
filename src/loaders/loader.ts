import {asyncRemove} from "../core.js";

export function LoaderStart():HTMLElement {
    let wrapper = document.createElement('div');
    wrapper.className = 'c__loader__wrapper';
    const animate_content = document.createElement('div'),
        dot = '<div class="c__loader__dot"></div>';
    animate_content.innerHTML = `<div class="c__loader">${dot}${dot}${dot}${dot}${dot}${dot}</div>`;

    wrapper.appendChild(animate_content);
    document.body.appendChild(wrapper);
    return wrapper;
}

export function LoaderStop(milliseconds: number, loader?: Element) {
    if (loader == null) {
        let l = document.querySelector('.c__loader__wrapper');
        if(l == null) {
            return;
        }
        loader = l;
    }

    asyncRemove(loader, milliseconds)
}