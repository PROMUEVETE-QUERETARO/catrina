export const HIDDEN_CLASS = "_hidden";

export function Button(id: string, text: string): HTMLButtonElement {
    return __button(id, text, 'c__button')
}

export function PrimaryButton(id: string, text: string): HTMLButtonElement {
    return __button(id, text, 'c__button c__button--primary')
}

export function DangerousButton(id: string, text: string): HTMLButtonElement {
    return __button(id, text, 'c__button c__button--dangerous')
}

function __button(id:string, text:string, classList: string): HTMLButtonElement {
    let button = document.createElement('button');
    button.id = id;
    button.innerText = text;
    button.className = classList;
    return button;
}

export function asyncRemove(node: Element, milliseconds: number) {
    setTimeout(()=>{
        node.remove();
    }, milliseconds);
}