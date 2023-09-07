
export namespace Dom {
/** Creates a new HTML element with the tag given
 * 
 * @param {string} tagname The HTML tag. EG "div", "button"
 * @param {string[] | string} cssClass CSS Class or classes to apply to the element
 * @param {Object} styles Style overrides
 */
export function el<T extends keyof HTMLElementTagNameMap>(tagName: T, cssClass?:string | string[], styles?:{[key:string]:string}) {
    let el = document.createElement(tagName)
    
    if(styles !== undefined) {
        for(let key of Object.keys(styles)) {
            el.style[key as any] = styles[key] as any
        }
    }    
    
    if(cssClass !== undefined && cssClass !== "") {
        if(Array.isArray(cssClass)) {
            for(let c of cssClass) {
                el.classList.add(c)
            }
        } else {
            el.classList.add(cssClass)
        }
    }
    
    return el as HTMLElementTagNameMap[T]
}

/** Gets a new HTML Element with its textContent set to text.
 * 
 * @param {string} text The text to put in the element
 * @param {string} tagname The HTML tag. EG "div", "button"
 * @param {string[] | string} cssClass CSS Class or classes to apply to the element
 * @param {Object} styles Style overrides
 */
export function textEl<T extends keyof HTMLElementTagNameMap>(tagName: T, text?: string, cssClass?:string | string[], styles?:{[key:string]:string}) {
    let x = el<T>(tagName, cssClass, styles)
    x.textContent = text ? text : ""
    return x

}

}

