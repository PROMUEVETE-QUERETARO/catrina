import {urlManager, selectButtonWithURL} from "./url-manager.js";
import {nav, selectButton} from "./nav-and-menu.js";

export class ScreenManager {

    constructor(screens, roles, navBar, returnButton) {
        this.screens = screens
        this.home = roles.home // objeto  AppScreen que actúa como página de inicio.
        this.safePage = roles.safePage // objeto ScreenC que cargará si no se encuentra la pantalla solicitada.
        this.login = roles.login // objeto ScreenC que se utilizará para iniciar sesión.
        this.nav = navBar // Elemento donde se insertarán los botones de navegación.
        this.returnButton = returnButton // Botón para regresar en el historial
    }

    run(){
        urlManager(this.screens, this.safePage, this.returnButton)
        this.nav.innerHTML = ''
        this.nav.appendChild(nav(this.screens))
        let buttons = document.querySelectorAll('.nav__button')
        selectButton(buttons, 'nav__button--selected')
        selectButtonWithURL()
    }


}

