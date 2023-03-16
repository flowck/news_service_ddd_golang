package news

type Topic struct {
	value string
}

func (t Topic) String() string {
	return t.value
}
