package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// MarusiaSavePictureBuilder builder.
//
// https://vk.com/dev/marusia_skill_docs10
type MarusiaSavePictureBuilder struct {
	api.Params
}

// NewMarusiaSavePictureBuilder func.
func NewMarusiaSavePictureBuilder() *MarusiaSavePictureBuilder {
	return &MarusiaSavePictureBuilder{api.Params{}}
}

// Server parameter.
func (b *MarusiaSavePictureBuilder) Server(v int) *MarusiaSavePictureBuilder {
	b.Params["server"] = v
	return b
}

// Photo parameter.
func (b *MarusiaSavePictureBuilder) Photo(v string) *MarusiaSavePictureBuilder {
	b.Params["photo"] = v
	return b
}

// Hash parameter.
func (b *MarusiaSavePictureBuilder) Hash(v string) *MarusiaSavePictureBuilder {
	b.Params["hash"] = v
	return b
}

// MarusiaDeletePictureBuilder builder.
//
// https://vk.com/dev/marusia_skill_docs10
type MarusiaDeletePictureBuilder struct {
	api.Params
}

// NewMarusiaDeletePictureBuilder func.
func NewMarusiaDeletePictureBuilder() *MarusiaDeletePictureBuilder {
	return &MarusiaDeletePictureBuilder{api.Params{}}
}

// ID parameter.
func (b *MarusiaDeletePictureBuilder) ID(v int) *MarusiaDeletePictureBuilder {
	b.Params["id"] = v
	return b
}
