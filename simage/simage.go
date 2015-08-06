package simage

import (
	"image"
	// Package image/png is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand PNG formatted images.
	"image/png"
	"os"
	"reflect"
	)

type Sippimage struct {
	Img *image.Gray
}

var grayType = reflect.TypeOf(new(image.Gray))

func Read(in *string) (img *Sippimage, err error) {
	reader, err := os.Open(*in)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	
	if reflect.TypeOf(m) != grayType {
		panic("input image must be grayscale!")
	}
		
	img = new(Sippimage)
	img.Img = m.(*image.Gray)
	return
}

func (img *Sippimage) Write(out *string) error {
	writer, err := os.Create(*out) 
	if err != nil {
		return err
	}
	return png.Encode(writer, img.Img)
}