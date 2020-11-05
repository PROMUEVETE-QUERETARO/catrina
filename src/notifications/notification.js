import {Element} from "../core/index.js";

export const notify = (attributes, content, parentNode) => {
    let notification = new Element('div', {className:'c__notification'}, {innerHTML:`${content}`})
    notification.style.color = `${ attributes.color }` //también puede ser una variable
    notification.style.backgroundColor = `${ attributes.background }`

    if(attributes.type === 'float') {
        notification.classList.add('c__notification--float')
        if(attributes.position === 'bottom') {
            console.log()
            notification.style.top = '80%'
        }
        document.body.appendChild(notification)
    } else {
        parentNode.appendChild(notification)
    }

    notification.addEventListener('click', ()=>{
        notification.remove()
    })
    deleteNotify(notification, attributes.duration)

    return notification
}
// En deleteNotify la notificación debe durar por lo menos un segundo, por lo tanto si el argumento 'duration'
// es menor a 1000 la notificación solo se podrá eliminar haciendo click sobre ella.
const deleteNotify = (notification, duration) => {
    if (duration != null){
        if (duration > 999){
            setTimeout(()=>{
            notification.remove()
            }, duration)
        }
    }
}
