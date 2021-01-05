import { Element } from "../DOM/elements.js";

export class Icon {
    constructor(className) {
        return new Element(
            'i',
            {
                className: `${className}`
            }
        )
    }
}