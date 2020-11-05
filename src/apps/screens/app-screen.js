import {Module} from "./screen-module.js";
import {salert} from "../../alerts/prebuild.js";
import {AppHistory} from "../history.js";
import {menu, selectButton} from "../nav-and-menu.js";
import {ENV, APP} from "../../../main.js";


export class AppScreen extends Module {
    constructor(title, icon, hash, content, subScreens, options) {
        super(title, hash, icon, ENV.screenContainer, content, options);
        this.menuContainer = ENV.screenMenuContainer // DOM Object
        this.subScreens = subScreens // []subScreen
    }

    run(historyReturn) {
        if(this.protected){
            if (this.protected.usersType.findIndex(t => t === this.user.type) === -1) {
                if(this.protected.strict.value){
                    APP.runScreen(`${APP.roles.safePage.title}`)
                    return
                }
                salert('bad', 'Inicia Sesión', 'Inicia Sesión para ver esta página', ()=>{
                    APP.runScreen(`${APP.roles.safePage.title}`)
                })
                return
            }
        }


        let titleSpace = document.getElementById('title_screen')

        location.hash = this.hash
        document.title = `${this.title} | ${APP.name}`
        AppHistory.addScreen(this.title, historyReturn)

        if (titleSpace) {
            titleSpace.innerText = this.title
        }

        this.container.innerHTML = ''
        if (typeof this.content === "object") {
            this.container.appendChild(this.content)
        }else {
            this.container.innerHTML = this.content
        }

        if(this.subScreens) {
            this.menuContainer.classList.remove('screen__menu__container--disabled')
            if (this.subScreens.length > 0) {
                this.menuContainer.innerHTML = ''
                this.menuContainer.appendChild(menu(this.subScreens))
                let buttons = document.querySelectorAll('.global__menu__button')
                selectButton(buttons, 'button__selected')

                this.subScreens[0].run()
                buttons[0].classList.add('button__selected')
            } else {
                this.menuContainer.innerHTML = ''
            }
        } else {
            this.menuContainer.innerHTML = ''
        }
    }
}