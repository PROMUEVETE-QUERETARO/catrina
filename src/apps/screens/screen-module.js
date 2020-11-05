import {USER, APP} from "../../../main.js";
import {Element} from "../../core/index.js";


export class Module {
    constructor(title, hash, icon, container, content, options) {
        this.title = title
        this.icon = icon
        this.hash = hash
        this.container = container
        this.content = content
        this.invisible = false
        this.user = USER
        if (typeof options == "object"){
            this.invisible = options.invisible // bool, permite ocultar del menú un módulo, aunque no esté protegido.
            if (typeof options.protected == "object") {
                this.protected = {
                    usersType: options.protected.usersType, // []<string>, array con los tipos de usuarios con acceso.
                    strict: {
                        value: options.protected.strict.value, // bool, si 'strict.value' no es falso, el objeto Module
                        // solo será visible en el menú por los usuarios con acceso.
                        minimumLevel: options.protected.strict.minimumLevel, // Para mejorar la seguridad, es
                        // posible restringir el acceso solo a usuarios con la combinación correcta de
                        // 'user.type && user.level'.
                    },

                }
                if(this.protected.strict.value){
                    this.invisible = true
                }
            }
        }

    }

    run(){
        this.container.innerHTML = ''
        this.container.appendChild(this.content)
    }
}
