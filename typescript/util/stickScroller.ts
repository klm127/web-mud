/**
 * Sticks a scroller to the bottom if it isn't scrolled
 */
export function StickScroller(scroller: HTMLElement) {
	const isAtBottom =
		scroller.scrollHeight - scroller.clientHeight <= scroller.scrollTop + 1
	if (isAtBottom) {
		scroller.scrollTop = scroller.scrollHeight - scroller.clientHeight
	}
}

/**
 * Sticks a scroller to the bottom every N seconds, if it isn't scrolled up.
 */
export function StickScrollerOnInterval(
	scroller: HTMLElement,
	interval: number
) {
	setInterval(() => StickScroller(scroller), interval)
}

export function ScrollerToBottom(scroller: HTMLElement) {
	scroller.scrollTop = scroller.scrollHeight - scroller.clientHeight
}
