export const HIDDEN_CLASS = "_hidden";

export async function GetTemplate(templateSrc: string): Promise<string> {
    return fetch(templateSrc).then(d => d.text())

}

