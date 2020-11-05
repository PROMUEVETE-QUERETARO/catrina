// * numberFormat da formato inglés a un número sin formato. Con el parámetro minFractionDigits
// * se selecciona la cantidad de decimales que debe llevar.
export const numberFormat = (number, minFractionDigits) => {
    return new Intl.NumberFormat("en-EN", {
        maximumFractionDigits: '2',
        minimumFractionDigits: minFractionDigits
    }).format(number)
}

export const priceFormat = (number) => {
    return`$ ${numberFormat(number, '2')}`
}

export const priceFormatIVA = (n) => {
    return priceFormat(n *= 16)
}


