export class User {
    constructor(loginAPI, verifyFunc, defaultFunc) {
        this.name = ''
        this.type = ''
        this.level = 0
        this.org = ''
        this.idOrg = ''
        this.loginApi = loginAPI
        this.verify = verifyFunc // Función que permite verificar el token

        const token = localStorage.getItem('token')
        if (token) {
            this.verify()
            let payload = parseJWT(`${token}`)
            this.name = payload.user
            this.type = payload.category
            this.level = payload.level
            this.org = payload.org
            this.idOrg = payload.idOrg
        } else {
            defaultFunc()
        }
    }

    login(bodyRequest, saveResponse,catchFunction){
        fetch(`${this.loginApi}`, {
            method: 'POST',
            body: JSON.stringify(bodyRequest)

        })
            .then(data=>data.json())
            .then(data=>{
                localStorage.setItem('token', `${data.token}`)
                sessionStorage.setItem('user', `${data.user}`)
                if (saveResponse) {
                    localStorage.setItem('SStorage', `${JSON.stringify(data)}`)
                }
                location.reload()
            })
            .catch(()=>{
                if(typeof catchFunction == "function"){
                    catchFunction()
                }
            })
    }

    logout(){
        localStorage.setItem('token', '')

        if (localStorage.getItem('SStorage')){
            localStorage.removeItem('SStorage')
        }

        sessionStorage.clear()
        location.reload()
    }

    static renew(verifyApi, saveResponse, catchFunction){
        let token = localStorage.getItem('token')
        if (!token || token === ''){
            return false
        }
         fetch(`${verifyApi}`, {
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + `${token}`,
                'Content-Type': 'application/json'
            }

        })
            .then(data => data.json())
            .then(data => {
                sessionStorage.setItem('User', `${data.user}`)
                localStorage.setItem('token', `${data.token}`)
                if(saveResponse){
                    localStorage.setItem('SStorage', `${JSON.stringify(data)}`)
                }
            })
            .catch(() => {
                if(typeof catchFunction == "function"){
                    catchFunction()
                }
                console.error('Token no válido')
                sessionStorage.clear()
                localStorage.setItem('token', '')
                localStorage.setItem('SStorage', '')
                location.reload()
            })

        return true
    }

}

const parseJWT = (token) => {
    let base64URL = token.split('.')[1],
        base64 = base64URL.replace(/-/g, '+').replace(/_/g, '/'),
        jsonPayload = decodeURIComponent(atob(base64).split('')
            .map( (c) => {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
            })
            .join(''))
    return JSON.parse(jsonPayload)
}
