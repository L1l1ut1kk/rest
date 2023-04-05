package invert

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)
   // "github.com/L1l1ut1kk/rest/models"

   // "github.com/jinzhu/gorm"



func invertImage(filename string) error {
    // Open the original image file
    origFile, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer origFile.Close()

    // Decode the original image file
    origImg, _, err := image.Decode(origFile)
    if err != nil {
        return err
    }

    // Invert the color channels of the original image
    bounds := origImg.Bounds()
    inv := image.NewRGBA(bounds)
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
        for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
            origColor := origImg.At(x, y)
            r, g, b, a := origColor.RGBA()
            invColor := color.RGBA{255 - uint8(r>>8), 255 - uint8(g>>8), 255 - uint8(b>>8), uint8(a>>8)}
            inv.Set(x, y, invColor)
        }
    }

    // Save the inverted image to disk
    outFilename := "inverted_" + filename
    outFile, err := os.Create(outFilename)
    if err != nil {
        return err
    }
    defer outFile.Close()
    err = jpeg.Encode(outFile, inv, &jpeg.Options{Quality: 100})
    if err != nil {
        return err
    }

    // Save the filenames to the database
    image := &Image{OriginalFileName: filename, InvertedFileName: outFilename}
    err = db.Create(image).Error
    if err != nil {
        return err
    }

    return nil
}
