
export const getRandomInt = (max) => Math.floor(Math.random() * Math.floor(max))

export const getRandomArrayInt = (bytes, length) =>{
    let array
    if (bytes === 8){
        array = new Uint8Array(length)
    }else if (bytes === 16){
        array = new Uint16Array(length)
    }else if (bytes === 32){
        array = new Uint32Array(length)
    }

    return window.crypto.getRandomValues(array);
}

