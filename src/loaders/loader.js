/*
* La clase Loader puede posicionar distintos tipos de loaders o animaciones de carga.
* Por defecto crea una pantalla sobre todos los objetos del DOM e inserta la animación
* elegida en el centro. Si el parámetro «parent» es un objeto del DOM', la animación
* se insertará en él.
* Si no se dispone de una animación, debe definirse el parámetro «defaultAnimation»
* como 'true' y se insertará una animación simple con color 'c-yellow1'.
* */
import {Element} from "../core/index.js";

export class Loader {
    constructor(animation, parent, defaultAnimation) {
        this.content = animation
        this.parent = parent
        this.defaultAnimation = !!defaultAnimation;
        this.loader = ''
        if (!animation){
            this.content = new DocumentFragment()
        }

        if (!this.parent) {
            this.loader = new Element('div', {className:'c__loader__background'},
                {child: this.content})
            if(this.defaultAnimation) {
                insertLoader(this.loader)
            }

            document.body.appendChild(this.loader)
            this.loader.style.display = 'none'
            return
        }


    }

    run(){
        this.loader.style.display = 'grid'
    }

    stop(delay){
        if(!delay){
            delay = 1
        }
        setTimeout(()=>{this.loader.remove()},delay)
    }
}



const insertLoader = (parentNode) => {
    let container = document.createElement('div'),
        loader = document.createElement('div')
    
    parentNode.innerHTML = ''

    loader.className = 'c__loader'
    container.className = 'c__loader__container'
    container.appendChild(loader)

    parentNode.appendChild(container)
}

