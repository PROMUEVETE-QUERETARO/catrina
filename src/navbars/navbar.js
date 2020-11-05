const navbarFunctions = () => {
    const iconNavbar = document.querySelector('.pwf__navbar__icon'),
        lateralbar = document.getElementById('pwf_lateralbar'),
        btnNavbar = document.getElementById('button_navbar'),
        lateralbarArrow = document.querySelectorAll('.pwf__lateralbar__icon'),
        lateralbarSubContent = document.querySelectorAll('.pwf__lateralbar__subcontainer')


    btnNavbar.addEventListener('click', ()=>{
        lateralbar.classList.toggle('pwf__lateralbar--open')
        iconNavbar.classList.toggle('pwf__navbar__icon--selected')
    })

    for (let i = 0; i < lateralbarArrow.length; i++) {
        lateralbarArrow[i].addEventListener('click', ()=>{
            lateralbarSubContent[i].classList.toggle('pwf__lateralbar__subcontainer--off')
            lateralbarArrow[i].classList.toggle('pwf__lateralbar__icon--open')
        })
        
    }
}

// * collapseLateralbar permite ocultar la barra lateral al seleccionar otro elemento
// * HTML, debe usarse de la siguiente manera:
// * * * Element.addEventListener('click', collapseLateralbar)

const collapseLateralbar = ()=> {
    const iconNavbar = document.querySelector('.pwf__navbar__icon'),
        lateralbar = document.getElementById('pwf_lateralbar')

    if(lateralbar.classList.contains('pwf__lateralbar--open')){
        lateralbar.classList.toggle('pwf__lateralbar--open')
        iconNavbar.classList.toggle('pwf__navbar__icon--selected')
    }
    
}

export {navbarFunctions, collapseLateralbar}