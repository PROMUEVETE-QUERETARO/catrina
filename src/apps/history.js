import {APP, ENV} from "../../main.js";
//import {appLog} from "./logs.js";


export class AppHistory {
    constructor(initialScreen) {
        sessionStorage.setItem('History', initialScreen)
        disableArrow()
    }

    static addScreen(screen, returned) {
        if(returned) return

        let history = sessionStorage.getItem('History')

        if(!history){
            sessionStorage.setItem('History', screen)
            return
        }

        let h = history.split(',')
        
        if(screen !== h[0]){
            h.unshift(screen)
            if(h.length > 1) enableArrow()
        }

        sessionStorage.setItem('History', h.toString())
    }

    static goBack() {
        let history = sessionStorage.getItem('History').split(','),
            screen = history[0]

        history.shift()
        APP.runScreen(screen, true)

        if(history.length <= 0) {
            disableArrow()
            history.push(screen)
        } else {
            enableArrow()
        }

        sessionStorage.setItem('History', history.toString())
    }

}

const enableArrow = () =>{
    ENV.returnButton.classList.remove('arrow_history--hidden')
    ENV.returnButton.removeAttribute('disabled')
}

const disableArrow = () => {
    ENV.returnButton.classList.add('arrow_history--hidden')
    ENV.returnButton.setAttribute('disabled', 'true')
}