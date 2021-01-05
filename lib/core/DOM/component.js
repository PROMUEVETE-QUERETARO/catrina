// Component permite crear elementos de la interfaz complejos, importables con pocas líneas de
// código.
// El parámetro ui debe ser un string que se pueda renderizar a HTML. El parámetro events
// debe contener el código js a ejecutar en ese componente; elements debe contener los HTMLElement
// exportables para que otros componentes puedan interactuar con ellos.
export class Component{
    constructor(template, elements, events) {
        this.template = template
        this.elements = elements
        this.events = events
    }

    reset(){
        return this.template
    }
}