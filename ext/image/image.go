package goproxy_image

import (
	"image"
	_ "image/gif"

	. "github.com/elazarl/goproxy"
)

var RespIsImage = ContentTypeIs("image/gif",
	"image/jpeg",
	"image/pjpeg",
	"application/octet-stream",
	"image/png")

// "image/tiff" tiff support is in external package, and rarely used, so we omitted it

func HandleImage(f func(img image.Image, ctx *ProxyCtx) image.Image) RespHandler {
	_ = "STUB: not implemented"
	return *new(RespHandler)
}

// we might get 304 - not modified response without data

// No gif image encoder in go - convert to png
