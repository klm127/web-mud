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

func (self *log) write(txt string) {
	self.time = time.Now()
	self.txt = txt
}

func (self *log) String() string {
	return fmt.Sprintf("%d:%d:%d %s", self.time.Hour(), self.time.Minute(), self.time.Second(), self.txt)
}

func (self *log) Time() time.Time {
	return self.time
}
func (self *log) TextOnly() string {
	return self.txt
}
