package object

import (
	"fmt"

	"github.com/eitarox/eitaroxgit/sha"
)

type Object struct {
	Hash sha.SHA1
	Type Type
	Size int
	Data []byte
}

func (o *Object) Header() []byte {
	return []byte(fmt.Sprintf("%s %d\x00", o.Type, o.Size))
}
