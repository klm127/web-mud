package message

import "fmt"

type MessageBuilder struct {
	message messageHolder
}

func New() *MessageBuilder {
	m := MessageBuilder{}
	m.message.Parts = make([]*messagePart, 0)
	m.message.Current = &messagePart{}
	return &m
}

func (mb *MessageBuilder) Next() *MessageBuilder {
	mb.message.Next()
	return mb
}

func (mb *MessageBuilder) Text(txt string) *MessageBuilder {
	mb.message.Current.TextValue = &txt
	return mb
}

func (mb *MessageBuilder) Textf(format string, a ...any) *MessageBuilder {
	s := fmt.Sprintf(format, a...)
	mb.message.Current.TextValue = &s
	return mb
}

func (mb *MessageBuilder) Color(txt string) *MessageBuilder {
	mb.message.Current.ColorValue = &txt
	return mb
}

func (mb *MessageBuilder) NewLine(num int8) *MessageBuilder {
	mb.message.Current.NewLineValue = &num
	return mb
}

func (mb *MessageBuilder) Link(link_cmd string) *MessageBuilder {
	mb.message.Current.LinkValue = &link_cmd
	return mb
}

func (mb *MessageBuilder) Indent(num int8) *MessageBuilder {
	mb.message.Current.TabValue = &num
	return mb
}

func (mb *MessageBuilder) Class(css_class string) *MessageBuilder {
	mb.message.Current.Class(css_class)
	return mb
}

func (mb *MessageBuilder) Classes(css_classes []string) *MessageBuilder {
	mb.message.Current.Classes(css_classes)
	return mb
}

func (mb *MessageBuilder) Bytes() []byte {
	return mb.message.String()
}
