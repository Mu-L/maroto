package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/image"
	"github.com/johnfercher/v2/maroto/props"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.NewMaroto(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(40,
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Left:    10,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Top:     10,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Left:    15,
			Top:     15,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(8, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(6, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Top:     5,
			Left:    10,
		}),
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Top:     5,
		}),
		image.NewFromFileCol(2, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left:    5,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(6, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		image.NewFromFileCol(2, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(8, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/imagegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}