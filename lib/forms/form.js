import { Element, Icon } from "../core/export.js";

export class Form {
    constructor(attributes, inputs, options) {
        this.attributes = attributes // Object
        this.inputs = inputs // Input[]
        if (typeof options === "object"){
            this.fielset = {
                className: options.fielset.className,
                legend: {
                    textContent: options.fielset.legend.textContent,
                    className: options.fielset.legend.className
                }
            }
        }

        let form =  new Element(
            'form',
            attributes
            ),
            fragment = new DocumentFragment()

        for (let input of inputs) {
            fragment.appendChild(input)
        }

        if (this.fielset){
            let fileset = new Element(
                    'fileset',
                    {className:`${this.fielset.className}`},
                    {child: new Element(
                            'legend',
                            {className: `${this.fielset.legend.className}`},
                            {textContent:`${this.fielset.legend.textContent}`}
                            )
                    }
                )

            fileset.appendChild(fragment)
            form.appendChild(fileset)

            return form
        }



        form.appendChild(fragment)

        return form
    }
}

export class Input {
    constructor(type, attributes, options) {
        this.type =  type
        this.attributes = attributes // Object

        if (typeof options === "object"){
            if (options.label) {
                this.label = {
                    className: options.label.className,
                    icon: options.label.icon

                }
            }
            if (options.fielset) {
                this.fielset = {
                    className: options.fielset.className,
                    legend: {
                        textContent: options.fielset.legend.textContent,
                        className: options.fielset.legend.className
                    }
                }
            }
        }

        let input = new Element(
            'input',
            this.attributes
        )
        input.type = this.type

        if (this.label){
            let label = new Element(
                    'label',
                    {className: `${this.label.className}`, HTMLFor: `${this.attributes.id}`}
                ),
                container = new Element(
                    'div',
                    {className:'c__form__subContainer'},
                    {child: label}
                )

            if (typeof this.label.icon === "string"){
                label.appendChild(new Icon(`${this.label.icon}`))
            } else if (typeof this.label.icon === "object"){
                label.appendChild(this.label.icon)
            }

            container.appendChild(input)

            return container
        }

        if (this.fielset){
            let fileset = new Element(
                    'fileset',
                    {className:`${this.fielset.className}`},
                    {child: new Element(
                            'legend',
                            {className: `${this.fielset.legend.className}`},
                            {textContent:`${this.fielset.legend.textContent}`}
                    )}
                )

            fileset.appendChild(input)
            return fileset
        }

        return input
    }
}

export const passwordInputs = () =>{
    let buttons = document.querySelectorAll('.c__form__button__password'),
        inputs = document.querySelectorAll('.c__form__input__password')

    for(let button of buttons){
        button.addEventListener('click', (e)=>{
            e.preventDefault()
            if(button.previousElementSibling.type === 'text'){
                button.innerHTML = '<i class="icon-eye-slash"></i>'
                button.previousElementSibling.type = 'password'
                button.previousElementSibling.focus()
            } else {
                button.innerHTML = '<i class="icon-eye"></i>'
                button.previousElementSibling.type = 'text'
                button.previousElementSibling.focus()
            }

        })
    }

    for(let input of inputs){
        input.addEventListener('focus', ()=>{
            input.parentNode.classList.add('c__form__passwordSubContainer--focus')
        })
        input.addEventListener('blur', ()=>{
            input.parentNode.classList.remove('c__form__passwordSubContainer--focus')
        })
    }
}

export const customSelectInputs = () => {
    let inputs = document.querySelectorAll('.c__form__c_select__input'),
        boxes = document.querySelectorAll('.c__form__c_select__box')

    for (let input of inputs) {
        input.addEventListener('click', ()=> {
            input.nextElementSibling.classList.add('c__form__c_select__box--active')
        })

        input.addEventListener('blur', ()=> {
            setTimeout(()=> {
                input.nextElementSibling.classList.remove('c__form__c_select__box--active')
            }, 500)
        })
    }

    for(let box of boxes) {
        let children = box.children
        for(let child of children) {
            child.addEventListener('click', ()=>{
                child.parentNode.classList.remove('c__form__c_select__box--active')
                child.parentNode.previousElementSibling.value = child.innerHTML
                child.parentNode.previousElementSibling.focus()
            })
        }
    }

}