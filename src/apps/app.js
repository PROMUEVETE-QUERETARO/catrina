import {ScreenManager} from "./screen-manager.js"
import {removeSelect} from "./nav-and-menu.js";
import {AppHistory} from "./history.js";

export class App {
    constructor(name, screens, roles, serviceWorker, env) {
        this.name = name
        this.screens = screens // Screen[]
        this.roles = roles
        this.serviceWorker = serviceWorker // bool, determina si se ha de cargar algún serviceWorker desde la raíz
        this.env = env // env contiene los objetos de la interfaz
    }

    run(){
        new ScreenManager(this.screens, this.roles, this.env.menuContainer)
            .run()
        if(this.serviceWorker) {
            navigator.serviceWorker.register('./sw.js')
                .then(response => console.log('SW ok', response))
                .catch(error => console.error("SW failed", error))
        }
    }

    runScreen(nameScreen, returnHistory){
        let screens = this.screens
        removeSelect(
            document.querySelectorAll('.nav__button'),
            'nav__button--selected'
        )
        for (let screen of screens){
            if(screen.title === nameScreen){
                screen.run()
                if(!returnHistory)AppHistory.addScreen(screen.title)

                const button = document.getElementById(`${screen.hash}`)
                if(button) button.classList.add('nav__button--selected')

                return screen
            }
        }
        this.roles.safePage.run()
        AppHistory.addScreen(screen.title)
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
        modules[0].run()
    }
}