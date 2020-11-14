export class User {
    constructor() {
        this.name = ''
        this.type = ''
        this.level = 0
    }

}

export const parseJWT = (token) => {
    let base64URL = token.split('.')[1],
        base64 = base64URL.replace(/-/g, '+').replace(/_/g, '/'),
        jsonPayload = decodeURIComponent(atob(base64).split('')
            .map( (c) => {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
            })
            .join(''))
    return JSON.parse(jsonPayload)
}
