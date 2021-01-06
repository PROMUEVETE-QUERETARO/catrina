export class Element {
    constructor(type, attributes, content) {
        this.type = type // string
        this.attributes = {
            id: attributes.id,
            className: attributes.className,
            name: attributes.name,
            type: attributes.type, // input
            placeholder: attributes.placeholder, // input
            disabled: attributes.disabled, // input
            readOnly: attributes.readOnly, // input
            size: attributes.size, // input
            maxLength: attributes.maxLength, // input
            min: attributes.min, // input
            max: attributes.max, // input
            step: attributes.step, // input.number
            scope: attributes.scope, // th
            title: attributes.title,
            value: attributes.value, // option || input
            list: attributes.list, // input text
            src: attributes.src, // img
            alt: attributes.alt, // img
            href: attributes.href, // a
            label: attributes.label, // optgroup
            HTMLFor: attributes.HTMLFor, // label || output
            action: attributes.action, // form
            target: attributes.target, // form
            method: attributes.method, // form
            autocomplete: attributes.autocomplete, // form
            acceptCharset: attributes.acceptCharset, // form
            multiple: attributes.multiple, // input || form
            pattern: attributes.pattern, // input
            required: attributes.required, // input
            autofocus: attributes.autofocus, // input
            rows: attributes.rows, // textarea
            cols: attributes.cols, // textarea
            onclick: attributes.onclick,
            oninput: attributes.oninput,
            onchange: attributes.onchange
        }
        if (typeof content == "object"){
            this.content = {
                textContent: content.textContent,
                child: content.child, // Contenido del elemento, se usará appendChild()
                children: content.children, // Contenidos del elemento, se usará appendChild() en ciclo for
                innerText: content.innerText,
                innerHTML: content.innerHTML
            }
        }
        return this.create()
    }
    create(){
        let element = document.createElement(`${this.type}`),
            a = this.attributes,
            c = this.content

        if(a.id) element.setAttribute('id', `${a.id}`)

        if(a.className) element.setAttribute( 'class', `${a.className}`)

        if(a.name) element.setAttribute( 'name', `${a.name}`)

        if(a.type) element.setAttribute( 'type', `${a.type}`)

        if(a.placeholder) element.setAttribute( 'placeholder', `${a.placeholder}`)

        if(a.disabled) element.setAttribute( 'disabled', `${a.disabled}`)

        if(a.readOnly) element.setAttribute( 'readonly', `${a.readOnly}`)

        if(a.size) element.setAttribute( 'size', `${a.size}`)

        if(a.maxLength) element.setAttribute( 'maxlength', `${a.maxLength}`)

        if(a.min) element.setAttribute( 'min', `${a.min}`)

        if(a.max) element.setAttribute( 'max', `${a.max}`)

        if(a.step) element.setAttribute( 'step', `${a.step}`)

        if(a.scope) element.setAttribute( 'scope', `${a.scope}`)

        if(a.title) element.setAttribute( 'title', `${a.title}`)

        if(a.value) element.setAttribute( 'value', `${a.value}`)

        if(a.list) element.setAttribute( 'list', `${a.list}`)

        if(a.src) element.setAttribute( 'src', `${a.src}`)

        if(a.alt) element.setAttribute( 'alt', `${a.alt}`)

        if(a.href) element.setAttribute( 'href', `${a.href}`)

        if(a.label) element.setAttribute( 'label', `${a.label}`)

        if(a.HTMLFor) element.setAttribute( 'for', `${a.HTMLFor}`)

        if(a.action) element.setAttribute( 'action', `${a.action}`)

        if(a.target) element.setAttribute( 'target', `${a.target}`)

        if(a.method) element.setAttribute( 'method', `${a.method}`)

        if(a.autocomplete) element.setAttribute( 'autocomplete', `${a.autocomplete}`)

        if(a.acceptCharset) element.setAttribute( 'accept-charset', `${a.acceptCharset}`)

        if(a.multiple) element.setAttribute( 'multiple', `${a.multiple}`)

        if(a.pattern) element.setAttribute( 'pattern', `${a.pattern}`)

        if(a.required) element.setAttribute( 'required', `${a.required}`)

        if(a.autofocus) element.setAttribute( 'autofocus', `${a.autofocus}`)

        if(a.rows) element.setAttribute( 'rows', `${a.rows}`)

        if(a.cols) element.setAttribute( 'cols', `${a.cols}`)

        // --------------- Actions --------------- \\
        if(typeof a.onclick === "function") element.onclick = a.onclick
        if(typeof a.oninput === "function") element.oninput = a.oninput
        if(typeof a.onchange === "function") element.onchange = a.onchange

        // --------------- Content --------------- \\
        if (c) {
            if (c.innerHTML) element.innerHTML = c.innerHTML
            else if (c.innerText) element.innerText = c.innerText
            else if (c.textContent) element.textContent = c.textContent
            else if (typeof c.child === "object") element.append(c.child)
            else if (typeof c.children === "object") for(let child of c.children) element.append(child)
        }

        return element
    }
}
//@stop

export class Icon {
    constructor(className) {
        return new Element('i',{ className: `${className}`})
    }
}
//@stop

export const setCssVariable = (property, value)=> document.documentElement.style.setProperty(property, value)
//@stop

export const getRandomInt = (max) => Math.floor(Math.random() * Math.floor(max))
//@stop

export const getRandomArrayInt = (bytes, length) =>{
    let array
    if (bytes === 8) array = new Uint8Array(length)
    else if (bytes === 16) array = new Uint16Array(length)
    else if (bytes === 32) array = new Uint32Array(length)

    return window.crypto.getRandomValues(array);
}
//@stop


export const priceFormatIVA = (n) => priceFormat(n *= 16)

export const priceFormat = (number) => `$ ${numberFormat(number, '2')}`

// * numberFormat da formato inglés a un número sin formato. Con el parámetro minFractionDigits
// * se selecciona la cantidad de decimales que debe llevar.
export const numberFormat = (number, minFractionDigits) => {
    return new Intl.NumberFormat("en-EN", {
        maximumFractionDigits: '2',
        minimumFractionDigits: minFractionDigits
    }).format(number)
}
//@stop

export const capitalizeString = (s) => s.trim().toLocaleLowerCase().replace(/\w\S*/g,(w)=>(w.replace(/^\w/, (c)=> c.toLocaleUpperCase())))
//@stop
