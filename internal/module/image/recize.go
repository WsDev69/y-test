package image

import (
	"github.com/nfnt/resize"
	goim "image"
)

type Recize struct {
}

func (r *Recize) Recize(img goim.Image) goim.Image {
	//todo to config value
	m := resize.Resize(160, 160, img, resize.Lanczos3)
	return m
}

func (r *Recize) RecizeIf(img goim.Image, eql func() bool) goim.Image {
	if eql() {
		return r.Recize(img)
	}
	return img
}

