export function Modal(title:string, btnClass: string):{ body: HTMLElement; btn: HTMLElement; modal: HTMLElement } {
    let fragment = new DocumentFragment(),
        window = document.createElement('div');
    window.className = 'c__modal';
    window.innerHTML = `
        <div class="c__modal__header">
            <button class="${btnClass} c__modal__btn icon-arrow-left"></button>
            <div class="c__modal__title">${title}</div>
        </div>
        <div class="c__modal__body"></div>
    `;

    fragment.appendChild(window);
    document.body.appendChild(fragment);
    document.body.style.overflow = 'hidden';

    let btn = window.querySelector('.c__modal__body') as HTMLElement;

    return {
        modal: window,
        btn: window.querySelector('.c__modal__btn') as HTMLElement,
        body: btn,
    }
}

export function CloseModal(modal: HTMLDivElement) {
    if(modal.classList.contains('c__modal')){
        modal.remove()
        document.body.style.overflow = 'inherit';
    }
}