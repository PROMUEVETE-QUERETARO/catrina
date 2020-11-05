import {ScreenManager} from "./screen-manager.js"
import {removeSelect} from "./nav-and-menu.js";
import {ENV} from "../../main.js";

export class App {
    constructor(name, screens, roles, serviceWorker) {
        this.name = name
        this.screens = screens // Screen[]
        this.roles = roles
        this.serviceWorker = serviceWorker // bool, determina si se ha de cargar algún serviceWorker desde la raíz
    }

    run(){
        new ScreenManager(this.screens, this.roles, ENV.menuContainer)
            .run()
        if(this.serviceWorker) {
            navigator.serviceWorker.register('./sw.js')
                .then(response => console.log('SW ok', response))
                .catch(error => console.error("SW failed", error))
        }
    }

    runScreen(nameScreen, historyReturn){
        let screens = this.screens
        removeSelect(
            document.querySelectorAll('.nav__button'),
            'nav__button--selected'
        )
        for (let screen of screens){
            if(screen.title === nameScreen){
                screen.run()
                if(document.getElementById(`${screen.hash}`)) {
                    document.getElementById(`${screen.hash}`).classList.add('nav__button--selected')
                }
                return screen
            }
        }
        this.roles.safePage.run(historyReturn)
    }

    runModule(nameScreen, nameModule){
        let screen = this.runScreen(`${nameScreen}`),
            modules = screen.modules
        for (let module of modules){
            if (module.title === nameModule){
                module.run()
                document.getElementById(`${screen.hash}`).classList.add('nav__button--selected')
                removeSelect(
                    document.querySelectorAll('.global__menu__button'),
                    'button__selected'
                )
                document.getElementById(`${module.hash}`).classList.add('button__selected')
                return
            }
        }
        this.roles.safePage.run()
    }
}