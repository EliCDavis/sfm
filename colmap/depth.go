package colmap

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readHeader(in io.Reader, deliminator byte, out []int) error {
	buf := make([]byte, 1)

	literal := &strings.Builder{}
	currentEntry := 0
	for currentEntry < len(out) {
		_, err := io.ReadFull(in, buf)
		if err != nil {
			return err
		}

		if buf[0] == deliminator {
			parsedInt, err := strconv.ParseInt(literal.String(), 10, 64)
			if err != nil {
				return err
			}

			out[currentEntry] = int(parsedInt)
			currentEntry++
			literal.Reset()
			continue
		}

		literal.WriteByte(buf[0])
	}
	return nil
}

type Depthmap struct {
	Width  int
	Height int
	Data   []float32
}

func (dm Depthmap) Value(x, y int) float32 {
	return dm.Data[x+(y*dm.Width)]
}

func (dm Depthmap) MaxValue() float32 {
	var f float32 = 0
	for _, v := range dm.Data {
		f = max(f, v)
	}
	return f
}

func ReadDepthmap(in io.Reader) (*Depthmap, error) {
	header := make([]int, 3)
	err := readHeader(in, '&', header)
	if err != nil {
		return nil, err
	}

	width := header[0]
	height := header[1]
	channels := header[2]
	if channels != 1 {
		return nil, fmt.Errorf("unexpected channel count %d. did you pass us a normal map?", channels)
	}

	data := make([]float32, width*height)
	err = binary.Read(in, binary.LittleEndian, &data)
	if err != nil {
		return nil, err
	}

	return &Depthmap{
		Width:  width,
		Height: height,
		Data:   data,
	}, nil
}

func LoadDepthmap(file string) (*Depthmap, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadDepthmap(bufio.NewReader(f))
}
