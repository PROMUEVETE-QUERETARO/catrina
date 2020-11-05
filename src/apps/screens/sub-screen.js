import {Module} from "./screen-module.js";
import {Element} from "../../core/index.js";


export class SubScreen extends Module {
    run(){
        this.container.innerHTML = ''
        if(this.protected){
            if (this.protected.usersType.findIndex(t => t === this.user.type) === -1) {
                if(this.protected.strict.value){
                    APP.runScreen(`${APP.roles.safePage.title}`)
                    return
                }
                this.content = new Element('p',{}, {textContent:'Para ver esta sección, inicia sesión'})
            }
        }

        if (typeof this.content === 'string'){
            this.container.innerHTML = this.content
        }   else {
            this.container.appendChild(this.content)
        }
    }
}

