package news

import (
	"fmt"
	"strings"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/gosimple/slug"
)

type Slug struct {
	value string
}

func NewSlugFromID(id domain.ID, value string) (Slug, error) {
	idSubstr := strings.Split(id.String(), "-")[3]

	if !slug.IsSlug(value) {
		return Slug{}, fmt.Errorf("%s is not a valid slug", value)
	}

	return Slug{value: fmt.Sprintf("%s-%s", value, idSubstr)}, nil
}

func MustNewSlugFromID(id domain.ID, value string) Slug {
	s, err := NewSlugFromID(id, value)
	if err != nil {
		panic(err)
	}

	return s
}

func NewSlugFromString(value string) (Slug, error) {
	if !slug.IsSlug(value) {
		return Slug{}, fmt.Errorf("%s is not a valid slug", value)
	}

	return Slug{value: value}, nil
}

func (s Slug) IsEmpty() bool {
	return s.value == ""
}

func (s Slug) String() string {
	return s.value
}
