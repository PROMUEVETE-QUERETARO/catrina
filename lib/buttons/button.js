import {Element,Icon} from "../core/export.js";

export class Button {
    constructor( content, attributes, onclick) {
        this.content = {
            icon: content.icon, // class del icon
            textContent: content.textContent
        }
        this.attributes = {
            color: attributes.color,
            id: attributes.id,
            className: attributes.className,
            title: attributes.title,
            disabled: attributes.disabled
        }
        this.onclick = onclick

        return this.create()
    }

    create(){
        let className = 'c__button',
            text = document.createElement('span'),
            button,
            onclick = ()=>{}

        if ( typeof this.onclick == "function" ) {
            onclick = this.onclick
        }

        switch ( this.attributes.color ){
            case 'red':
                className += " c__button--red"
                break
            case 'yellow':
                className += " c__button--yellow"
                break
            case 'exit':
                className += " c__button--red"
                break
            default:
                className += " c__button--white"
                break
        }

         button = new Element(
            'button',
            {
                className: className,
                onclick: onclick
            }
        )
        if(this.content.icon){
            button.appendChild(new Icon(`${this.content.icon}`))
        } else {
            text.innerText = ' My Button'
            button.appendChild(text)
        }

        if (this.content.textContent) {
            text.innerText = ' ' + this.content.textContent
            button.appendChild(text)
            if(!this.attributes.title) {
                button.title = this.content.textContent
            }
        } else {
            if(!this.attributes.title) {
                button.removeAttribute('title')
            } else {
                button.title = this.attributes.title
            }
        }

        if (typeof this.attributes == "object"){
            if(this.attributes.id){
                button.setAttribute('id', `${this.attributes.id}`)
            }
            if (this.attributes.className){
                button.setAttribute('class', `${this.attributes.className}`)
            }
            if (this.attributes.disabled){
                button.disabled = true
            }
        }

        return button
    }
}
//@stop