package api

import (
	"github.com/stretchr/objx"

	"github.com/PeerXu/meepo/pkg/meepo"
)

type NewServerOption = func(objx.Map)

func WithMeepo(meepo *meepo.Meepo) NewServerOption {
	return func(o objx.Map) {
		o["meepo"] = meepo
	}
}
