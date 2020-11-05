import {Element, Icon} from "../core/index.js"
// * nav recibe una colección de objetos Screen y crea una barra de navegación con todos los objetos.
export const nav = (screens) => {
    let subContainer = new DocumentFragment()
    for (let screen of screens) {
        let icon = createIcon(screen, {
            icon: "nav__icon",
            button: "nav__button",
            name: "nav__name"
        })
        if (!screen.invisible) {
            subContainer.appendChild(icon)
        }
    }

    return new Element('ul',{className:'global__nav__list'},{child:subContainer})
}
// * menu recibe una colección de objetos module y crea un menú que pueda cargar su contenido
export const menu = (modules) => {
    let subContainer = new DocumentFragment()
    for (let module of modules) {
        let icon = createIcon(module,{
            icon: "",
            button: "global__menu__button",
            name: ""
        })
        if (!module.invisible) {
            subContainer.appendChild(icon)
        }
    }

    return new Element('ul',{},{child:subContainer}
    )
}

// * createIcon recibe un objeto tipo Screen o un objeto tipo Module de PWF y crea un ícono que permitirá
// * ejecutar el objeto.
// * Es necesario especificar el contenedor que delimita la pantalla.
// * El parámetro classNames es un objeto y puede tener tres atributos: icon, button y name.
const createIcon = (object, classNames) => {
    let icon = new Icon(`${object.icon} ${classNames.icon}`),
        button = new Element(
            'li',
            {
                    className:`${classNames.button}`, id:`${object.hash}`,
                    onclick:()=> {
                        object.run()
                    }
                },
            {child: icon}
        ),
        name = new Element(
            'span',
            {className: `${classNames.name}`},
            {textContent:` ${object.title}`}
        )

    button.appendChild(name)
    return button
}

export const selectButton = (buttons, className) => {
    for (let button of buttons){
        button.addEventListener('click', ()=> {
            removeSelect(buttons, className)
            button.classList.add(className)
        })
    }
}

export const removeSelect = (buttons, className) => {
    for (let button of buttons){
        button.classList.remove(className)
    }
}
