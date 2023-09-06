package console

import "fmt"

func Bg(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
}

func Fg(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func GetFgPrintf(r uint8, g uint8, b uint8) func(format string, a ...any) {
	s := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return func(format string, a ...any) {
		fmt.Printf(s+format+Default(), a...)
	}
}

func GetFgSprintf(r uint8, g uint8, b uint8) func(format string, a ...any) string {
	s := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return func(format string, a ...any) string {
		return fmt.Sprintf(s+format+Default(), a...)
	}
}

func Default() string {
	return "\x1b[0m"
}
