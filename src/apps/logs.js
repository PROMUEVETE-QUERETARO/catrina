export const appLog = (log) => {
    const l = localStorage.getItem('Log')
    if(!l) {
        localStorage.setItem('Log', `${log}`)
    }else {
        let logs = l.split(',')
        logs.push(`${log}`)
        localStorage.setItem('Log', logs.toString())
    }
    console.warn(log)
}