export const capitalizeString = (s) =>{
    return s.trim().toLocaleLowerCase().replace(/\w\S*/g,(w)=>(w.replace(/^\w/, (c)=> c.toLocaleUpperCase())))
}
