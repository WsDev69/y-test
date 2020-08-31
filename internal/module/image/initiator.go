package image

import "image"

type Service interface {
	Recize(img image.Image) image.Image
	RecizeIf(img image.Image, eql func() bool) image.Image
}
