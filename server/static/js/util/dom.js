export var Dom;
(function (Dom) {
    /** Creates a new HTML element with the tag given
     *
     * @param {string} tagname The HTML tag. EG "div", "button"
     * @param {string[] | string} cssClass CSS Class or classes to apply to the element
     * @param {Object} styles Style overrides
     */
    function el(tagName, cssClass, styles) {
        let el = document.createElement(tagName);
        if (styles !== undefined) {
            for (let key of Object.keys(styles)) {
                el.style[key] = styles[key];
            }
        }
        if (cssClass !== undefined && cssClass !== "") {
            if (Array.isArray(cssClass)) {
                for (let c of cssClass) {
                    el.classList.add(c);
                }
            }
            else {
                el.classList.add(cssClass);
            }
        }
        return el;
    }
    Dom.el = el;
    /** Gets a new HTML Element with its textContent set to text.
     *
     * @param {string} text The text to put in the element
     * @param {string} tagname The HTML tag. EG "div", "button"
     * @param {string[] | string} cssClass CSS Class or classes to apply to the element
     * @param {Object} styles Style overrides
     */
    function textEl(tagName, text, cssClass, styles) {
        let x = el(tagName, cssClass, styles);
        x.textContent = text ? text : "";
        return x;
    }
    Dom.textEl = textEl;
})(Dom || (Dom = {}));
