
export function Alert({title, content, buttons, skip = true}: {title:string, content:string, buttons: HTMLButtonElement[], skip:boolean}):{ shadow: HTMLElement; content: HTMLElement } {
    let fragment = new DocumentFragment(),
        wrapper = __alertWrapper(),
        alert = __alertHtml(title, content, skip);
    wrapper.appendChild(alert);
    fragment.appendChild(wrapper)
    document.body.appendChild(fragment);

    // skipped alert
    if (skip) {
        wrapper.addEventListener('click', (e)=> {
            if (e.target) {
                let target = e.target as HTMLElement;
                if (target.classList.contains('c__alert_wrapper') || target.classList.contains('c__alert__button')) {
                    wrapper.remove();
                    alert.remove();
                }
            }
        })
    } else {
        wrapper.classList.add('c__alert_wrapper--no_skip');
        let btn = alert.querySelector('.c__alert__button');
        if (btn) {
            btn.remove();
        }
    }

    //Add buttons
    let buttonsArea = alert.querySelector('.c__alert__buttons');
    if (buttonsArea != null){
        buttons.forEach(b => {
            // @ts-ignore
            buttonsArea.appendChild(b)
        })
    }

    return {
        shadow: wrapper,
        content: alert
    }
}


function __alertHtml(title: string, content: string, skip: boolean): HTMLElement {
    let container = document.createElement('div');
    container.className = 'c__alert'
    let headerClass = skip ? 'c__alert__header' : 'c__alert__header c__alert__header--no_skip';
    container.innerHTML = `
        <div class="${headerClass}">
            <p>${title}</p>
            <div><button class="c__button c__alert__button icon-close"></button></div>
        </div>
        <div class="c__alert__body">${content}</div>
        <div class="c__alert__buttons"></div>
    `;

    return container
}


function __alertWrapper(): HTMLElement {
    let wrapper = document.createElement('div');
    wrapper.className = 'c__alert_wrapper'
    if(document.querySelector('.c__alert_wrapper')) {
        wrapper.classList.add('c__alert_wrapper--clear')
    }
    return wrapper;
}
