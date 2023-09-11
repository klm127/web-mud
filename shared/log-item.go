package shared

import (
	"fmt"
	"time"
)

type log struct {
	time time.Time
	txt  string
}

func newlog(txt string) log {
	alog := log{}
	alog.write(txt)
	return alog
}

func (l *log) write(txt string) {
	l.time = time.Now()
	l.txt = txt
}

func (l *log) String() string {
	return fmt.Sprintf("%d:%d:%d %s", l.time.Hour(), l.time.Minute(), l.time.Second(), l.txt)
}

func (l *log) Time() time.Time {
	return l.time
}
func (l *log) TextOnly() string {
	return l.txt
}
