package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    var image = make([][]uint8, dx)

    for x := 0; x < dx; x++ {
        image[x] = make([]uint8, dy)
        for y := 0; y < dy; y++ {
            image[x][y] = uint8(x^y)
        }
    }
    return image
}

func main() {
    pic.Show(Pic)
}
