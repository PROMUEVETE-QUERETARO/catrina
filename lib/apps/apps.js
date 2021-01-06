export class App {
    constructor(name, screens, user) {
        this.name = name
        this.screens = screens
        this.user = user
    }

    runScreen(name){
        for (const screen of this.screens) {
            if (screen.name !== name) continue
            screen.run(this.user)
        }
    }

}
//@stop

export class UrlManager {
    constructor(app, noFoundFunc) {
        this.app = app
        this.noFound = ()=> alert("404 No found")
        if(typeof noFoundFunc === "function") this.noFound = noFoundFunc
    }

    run(){
        const l = location
        for (const screen of this.app.screens) {
            if(l.hash !== screen.hash) continue
            screen.run(this.app.user)
            return
        }
        this.noFound()
    }

    static getParams(paramsList, clearUrl) {
        const search = location.search
        if(search !== '') {
            const urlP = new URLSearchParams(search)
            let c = '{'
            for (const param of paramsList) {
                c += `"${param}":"${urlP.get(param)}"`
            }
            c = c.slice(0, -1)
            c += '}'
            sessionStorage.setItem('UrlParams', c)

            if(!clearUrl) return

            let l = location.toString()
            l = l.replace(`${search.toString()}`, '')
            history.pushState('', document.title, l)
        }
    }
}
//@stop

export class Screen {
    constructor(name, hash, buildFunc, logic) {
        this.name = name
        this.hash = `#/${hash}`
        this.build = ()=> console.log('Hello World!')
        this.logic = ()=> console.log('Hello World!')

        if (typeof buildFunc === "function") this.build = buildFunc
        if (typeof logic === "function") this.logic = logic
    }


    run(user) {
        document.title = this.name
        location.hash = this.hash
        this.build(user)
        this.logic(user)
    }
}
//@stop

export class User {
    constructor(name, guess, user, admin) {
        this.name = name // String
        this.guess = guess // boolean
        this.user = user // boolean
        this.admin = admin //boolean
    }

    isUser(){ return (this.user === true && !this.guess && !this.admin) }

    isAdmin() { return (this.admin === true && !this.isUser() && !this.guess) }

    isGuess() { return !this.isUser() && !this.isAdmin() }
}
//@stop
