import {asyncRemove} from "../../catrina";

export function FloatNotify(type: notifyType, content: string, milliseconds: number) {
    let n = Notify({type: type, content: content, isFloat: true});
    n.addEventListener('click', ()=> n.remove());
    asyncRemove(n as HTMLElement, milliseconds);
}

export function Notify({type, content, isFloat, parent}:{type: notifyType, content: string, isFloat:boolean, parent?: HTMLElement}): HTMLElement {
    let notification = document.createElement('div');
    notification.className =  isFloat ? 'c__notification c__notification--float': 'c__notification';

    switch (type) {
        case notifyType.Error:
            notification.classList.add('_colorbox_error');
            break;
        case  notifyType.Warning:
            notification.classList.add('_colorbox_warning');
            break;
        case  notifyType.Good:
            notification.classList.add('_colorbox_good');
            break;
        case notifyType.Emphasis:
            notification.classList.add('_colorbox_emphasis');
            break;
        default:
            notification.classList.add('_colorbox_contrast');
            break;
    }
    notification.classList.add('_color_contrast');
    notification.innerHTML = content;

    if (!isFloat && parent != null) {
        parent.appendChild(notification);
    } else {
        document.body.appendChild(notification);
    }

    return notification;
}

export enum notifyType {
    Error = -1,
    Neutral,
    Warning,
    Good,
    Emphasis,
}

