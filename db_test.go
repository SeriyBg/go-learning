package cms

import (
	"testing"
	"strconv"
	"reflect"
	"github.com/stretchr/testify/assert"
)

var p *Page

func Test_CreatePage(t *testing.T) {
	p = &Page{
		Title:   "test",
		Content: "test",
	}
	id, err := CreatePage(p)
	assert.NoError(t, err)
	p.ID = id
}

func Test_GetPage(t *testing.T)  {
	page, err := GetPage(strconv.Itoa(p.ID))
	assert.NoError(t, err)
	assert.Equal(t, p.ID, page.ID)
	assert.True(t, reflect.DeepEqual(page, p))
}
