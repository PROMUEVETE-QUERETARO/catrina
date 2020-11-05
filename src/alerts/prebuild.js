import {Alert} from "./alert.js";
import {Button} from "../buttons/button.js";

// salert es una función que permite crear Alerts prefabricados
export const salert = (type, title, text, btnFunction) =>{
    if(typeof text !== "string" || typeof btnFunction !== "function"){
        console.log('Parámetro incorrecto. Revisa la documentación')
        return
    }
    let alert = new Alert(
        {
            title: `${title}`,
            type: `${type}`,
            skip: false
        },
        `${text}`,
        [
            new Button(
                {
                    textContent: 'Aceptar'
                },
                {},
                ()=>{
                    btnFunction()
                    Alert.delete(alert)
                }
            )
        ]
    )
}