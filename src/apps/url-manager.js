//* urlParams, carga la pantalla correspondiente a la url escrita. Tiene un parámetro obligatorio
//*  y uno opcional: screens y safePage, correspondientemente.
//* + screens es array de objetos tipo Screen,
//* + safePage debe ser un objeto tipo Screen el cual va a ejecutarse en caso de que el hash de la url no coincida
//*   con ningún objeto tipo Screen. En caso de no pasar ningún argumento, o que el argumento sea distinto a un objeto,
//*   el primer objeto parámetro screens será definido como página segura.
import {AppHistory} from "./history.js";

export const urlManager = (screens, safePage, button) => {
    for (let screen of screens) {
        if (location.hash === screen.hash) {
            screen.run()
            AppHistory.addScreen(screen.title, false, button)
            return
        }
    }
    if (typeof(safePage) !== "object") {
        screens[0].run()
        AppHistory.addScreen(screens[0].title, false, button)
    } else {
        safePage.run()
        AppHistory.addScreen(safePage.title, false, button)
    }
}

export const selectButtonWithURL = () => {
    if (document.getElementById(`${location.hash}`)) {
        document.getElementById(`${location.hash}`).classList.add('nav__button--selected')
    }
}
