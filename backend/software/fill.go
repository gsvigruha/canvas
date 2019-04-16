package softwarebackend

import (
	"image/color"

	"github.com/tfriedel6/canvas/backend/backendbase"
)

func (b *SoftwareBackend) Clear(pts [4][2]float64) {
	iterateTriangles(pts[:], func(tri [][2]float64) {
		b.fillTriangle(tri, func(x, y int) {
			if b.clip.AlphaAt(x, y).A == 0 {
				return
			}
			b.Image.SetRGBA(x, y, color.RGBA{})
		})
	})
}

func (b *SoftwareBackend) Fill(style *backendbase.FillStyle, pts [][2]float64) {
	iterateTriangles(pts[:], func(tri [][2]float64) {
		b.fillTriangle(tri, func(x, y int) {
			if b.clip.AlphaAt(x, y).A == 0 {
				return
			}
			b.Image.SetRGBA(x, y, style.Color)
		})
	})
}

func (b *SoftwareBackend) ClearClip() {
	p := b.clip.Pix
	for i := range p {
		p[i] = 255
	}
}

func (b *SoftwareBackend) Clip(pts [][2]float64) {
	p2 := b.clip2.Pix
	for i := range p2 {
		p2[i] = 0
	}

	iterateTriangles(pts[:], func(tri [][2]float64) {
		b.fillTriangle(tri, func(x, y int) {
			b.clip2.SetAlpha(x, y, color.Alpha{A: 255})
		})
	})

	p := b.clip.Pix
	for i := range p {
		if p2[i] == 0 {
			p[i] = 0
		}
	}
}
