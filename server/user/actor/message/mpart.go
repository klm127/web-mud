package message

/*
	Represents a part of a message.
*/
type messagePart struct {
	TextValue    *string        `json:"Text,omitempty"`
	ColorValue   *string        `json:"Color,omitempty"`
	NewLineValue *int8          `json:"NewLine,omitempty"`
	TabValue     *int8          `json:"Indent,omitempty"`
	LinkValue    *string        `json:"Link,omitempty"`
	ClassesValue []*string      `json:"Css,omitempty"`
	owner        *messageHolder `json:"-"`
}

func (m *messagePart) Class(txt string) {
	m.ClassesValue = append(m.ClassesValue, &txt)
}

func (m *messagePart) Classes(txt []string) {
	for _, v := range txt {
		m.ClassesValue = append(m.ClassesValue, &v)
	}
}
