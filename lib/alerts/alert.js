import {Element,setCssVariable} from "../core/core.js";
import {Button} from "../buttons/button.js";

// salert es una función que permite crear Alerts prefabricados
export const salert = (type, title, text, btnFunction) =>{
    if(typeof text !== "string" || typeof btnFunction !== "function"){
        console.log('Parámetro incorrecto. Revisa la documentación')
        return
    }
    let alert = new Alert(
        {
            title: `${title}`,
            type: `${type}`,
            skip: false
        },
        `${text}`,
        [
            new Button(
                {
                    textContent: 'Aceptar'
                },
                {},
                ()=>{
                    btnFunction()
                    Alert.delete(alert)
                }
            )
        ]
    )
}

// Se recomienda que solo se utilicen un máximo de tres botones por alert, por cuestiones estéticas.
export class Alert{
    constructor(header, content, buttons, options){
        this.header = {
            title: header.title,// string li
            type: header.type, // string
            skip: header.skip, // bool
            button: {}
        }
        this.content = content // Object (DOM Node)
        this.buttons = buttons // array de buttons
        if (typeof options === "object") {
            this.options = {
                width: options.width, // number (Porcentaje de la página)
                font: options.font // string
            }
        }

        let container = new DocumentFragment(),
            shadow = this.createShadow(),
            alert = this.createAlert()

        container.append(shadow, alert)

        if(this.header.skip){
            shadow.addEventListener('click', ()=>{
                document.body.removeChild(alert)
                document.body.removeChild(shadow)
            })
            this.header.button.addEventListener('click', ()=>{
                document.body.removeChild(alert)
                document.body.removeChild(shadow)
            })
        } else {
            shadow.classList.add('pwf__alert__background--noSkip')
        }

        document.body.appendChild(container)

        return {
            shadow: shadow,
            content: alert
        }
    }

    createShadow(){
        let className = 'pwf__alert__background'
        if (document.querySelector('.pwf__alert__background')){
            className += ' pwf__alert__background--noShadow'
        }

        return new Element(
            'div',
            {className: `${className}`}
        )
    }

    createAlert(){
        let header = alertHeader(this.header.type, this.header.title, this.header.skip),
            alert = new Element(
            'div',
            {className: 'pwf__alert__canvas'},
            {
                child: header.header
            }
        )
        this.header.button = header.button

        if(this.options) {
            // Se permite seleccionar el ancho del alert siempre y cuando este sea mayor que 45 y menor que 90. El valor se trasladará a porcentaje
            if (typeof this.options.width == "number" && (this.options.width > 45 && this.options.width < 91)) {
                setCssVariable('--width-alert', `${this.options.width}%`)
            }

            if (typeof this.options.font == "string") {
                setCssVariable('--alert-font', `${this.options.font}`)
            } else {
                setCssVariable('--alert-font', 'armata')

            }
        }
        alert.append(alertBody(this.content), alertFooter(this.buttons))
        return alert
    }

    static delete(alert){
        alert.shadow.remove()
        alert.content.remove()
    }

    static delAlerts(){
        let backgrounds = document.querySelectorAll('.pwf__alert__background'),
            canvas = document.querySelectorAll('.pwf__alert__canvas')
        for (let i = 0; i < backgrounds.length; i++) {
            backgrounds[i].remove()
            canvas[i].remove()
        }
    }
}

const alertBody = (content) => {
    let body = document.createElement('div')

    if(typeof(content) == 'string'){
        body.innerHTML = content
    } else if (typeof(content) == 'object') {
        body.appendChild(content)
    } else {
        body.innerHTML = '<p>Pon aquí tu mensaje</p>'
    }

    body.setAttribute('class', 'pwf__alert__body')

    return body
}

const alertFooter = (alertButtons) => {
    let buttons = new DocumentFragment()
    if(alertButtons) {
        for (let i = alertButtons.length; i--;) {
            buttons.appendChild(alertButtons[i])
        }
    }

    return new Element(
        'div',
        { className: "pwf__alert__footer" },
        {children:[
                new Element(
                    'hr',
                    { className: "pwf__alert__footer__hr" }
                ),
                buttons
            ]}
    )
}

const alertHeader = (type, title, skip) => {
    let header = new Element(
        'div',
        {
            className: "pwf__alert__header"
        },
        {
            children:[
                new Element(
                    'div',
                    {className: "pwf__alert__header__section pwf__alert__header__section--icon"},
                    {child: defineIcon(type)}
                ),
                new Element(
                    'div',
                    {className: "pwf__alert__header__section "},
                    {child: new Element(
                                'h2',
                                {className: 'pwf__alert__header__title'},
                                {textContent: title}
                    )}
                ),
            ]
        }
    )

    let button = new Button(
            {
                icon: 'icon-close'
            },
            {
                className: "pwf__alert__header__button"
            }
        ),
        scapeButton = new Element(
            'div',
            {
                className: "pwf__alert__header__section pwf__alert__header__section--close"
            },
            {
                child: button
            }
        )

    if(skip){
        header.appendChild(scapeButton)
    }
    return {
        header: header,
        button: button
    }
}

const defineIcon = (type) => {
    if (type === 'none' || !type){
        return
    }
    let icon = document.createElement('i')
    switch (type){
        case 'error':
            icon.setAttribute('class', 'icon-delete-circle')
            break;

        case 'confirmation':
            icon.setAttribute('class', 'icon-check-circle')
            break;

        case 'login':
            icon.setAttribute('class', 'icon-lock')
            break;

        case 'suscribe':
            icon.setAttribute('class', 'icon-edit')
            break;

        default:
            icon.setAttribute('class', 'icon-exclamation-circle')
            break;
    }

    return icon
}
//@stop

