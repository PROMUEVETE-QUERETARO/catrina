export class AppHistory {
    constructor(app, button) {
        this.app = app
        this.button = button
        disableArrow(this.button)
    }

    static addScreen(screen, returned, button) {
        if(returned) return

        let history = sessionStorage.getItem('History')

        if(!history){
            sessionStorage.setItem('History', screen)
            return
        }

        let h = history.split(',')
        
        if(screen !== h[0]){
            h.unshift(screen)
            if(h.length > 1) enableArrow(button)
        }

        sessionStorage.setItem('History', h.toString())
    }

    goBack() {
        let history = sessionStorage.getItem('History').split(','),
            screen = history[0]

        history.shift()
        this.app.runScreen(screen, true)

        if(history.length <= 0) {
            disableArrow(this.button)
            history.push(screen)
        } else {
            enableArrow(this.button)
        }

        sessionStorage.setItem('History', history.toString())
    }

}

const enableArrow = (button) =>{
    button.classList.remove('arrow_history--hidden')
    button.removeAttribute('disabled')
}

const disableArrow = (button) => {
    button.classList.add('arrow_history--hidden')
    button.setAttribute('disabled', 'true')
}