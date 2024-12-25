package colmap

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/EliCDavis/vector/vector3"
)

type Normalmap struct {
	Width  int
	Height int
	Data   []float32
}

// https://github.com/colmap/colmap/blob/main/scripts/python/read_write_dense.py#L40
/*
import numpy as np

arr = np.array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23])

width = 4
height = 2
channels = 3

arr = arr.reshape((width, height, channels), order="F")
print(arr)

> [[[ 0  8 16]
>   [ 4 12 20]]
>
>  [[ 1  9 17]
>   [ 5 13 21]]
>
>  [[ 2 10 18]
>   [ 6 14 22]]
>
>  [[ 3 11 19]
>   [ 7 15 23]]]

arr = np.transpose(arr, (1, 0, 2)).squeeze()
print(arr)

> [[[ 0  8 16]
>   [ 1  9 17]
>   [ 2 10 18]
>   [ 3 11 19]]
>
>  [[ 4 12 20]
>   [ 5 13 21]
>   [ 6 14 22]
>   [ 7 15 23]]]

*/

func (nm Normalmap) Value(x, y int) vector3.Float32 {
	size := nm.Width * nm.Height
	i := (x + (y * nm.Width))
	return vector3.New(nm.Data[i], nm.Data[i+size], nm.Data[i+(size*2)])
}

func ReadNormalmap(in io.Reader) (*Normalmap, error) {
	header := make([]int, 3)
	err := readHeader(in, '&', header)
	if err != nil {
		return nil, err
	}

	width := header[0]
	height := header[1]
	channels := header[2]
	if channels != 3 {
		return nil, fmt.Errorf("unexpected normal map channel count %d", channels)
	}

	data := make([]float32, width*height*3)
	err = binary.Read(in, binary.LittleEndian, &data)
	if err != nil {
		return nil, err
	}

	return &Normalmap{
		Width:  width,
		Height: height,
		Data:   data,
	}, nil
}

func LoadNormalmap(file string) (*Normalmap, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadNormalmap(f)
}
