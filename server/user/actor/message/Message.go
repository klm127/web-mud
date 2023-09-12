package message

import "encoding/json"

type messageHolder struct {
	Parts   []*messagePart `json:"Parts"`
	Current *messagePart   `json:"-"`
}

func (m *messageHolder) Next() *messagePart {
	m.Parts = append(m.Parts, m.Current)
	m.Current = &messagePart{}
	return m.Current
}

func (m *messageHolder) String() []byte {
	m.Next()
	b, err := json.Marshal(m)
	if err != nil {
		return []byte("Error marshalling message.")
	} else {
		return b
	}
}
