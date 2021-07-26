/*
* tooltipsFunc, activa la animación de todos los tooltips creados con base en el siguiente patrón:
*   <html>
*       ...
*          <div class="c__tooltip__container">
*            <button><i class="icon-save"></i></button>
*            <span class="pcb_quote__control__tooltip">Guardar</span>
*          </div>
*       ...
*   </html>
* No es importante el tipo de elementos, lo importante el uso de las clases css, de igual forma,
* es necesario usar cargar el archivo 'tooltips.css'.
* */
export const tooltipsFunc = () => {
    const tooltips = document.querySelectorAll('.c__tooltip')
    for (let tooltip of tooltips) {
        tooltip.previousSibling.previousSibling.addEventListener('mouseover', ()=>{
            tooltip.classList.add('c__tooltip--active')
            if(/Firefox/.test(navigator.userAgent)) tooltip.classList.add('c__tooltip--active_mz')
        })
        tooltip.previousSibling.previousSibling.addEventListener('mouseout', ()=>{
            tooltip.classList.remove('c__tooltip--active')
            if(/Firefox/.test(navigator.userAgent)) tooltip.classList.remove('c__tooltip--active_mz')
        })
    }
}
//@stop