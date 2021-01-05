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


