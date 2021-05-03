package challenges

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/youpy/go-wav"
)

// Write a wav file out, reading each sample and converting to bigendian
// Straightforward when you figure it out - you have to be willing to dive into how wav files work though
// You also have to be able to stomach that writing the samples out in bigendian can be deduced from a map of India
// http://www.pythonchallenge.com/pc/hex/bin.html
func Solve19() {
	contents, err := os.Open("./challenges/data/chal19input.txt") // I'm not parsing html comments out

	if err != nil {
		panic(err)
	}

	defer contents.Close()

	reader := base64.NewDecoder(base64.RawStdEncoding, contents)

	bufio.NewReader(reader)

	data, _ := io.ReadAll(reader)

	wread := wav.NewReader(bytes.NewReader(data))

	wavFormat, _ := wread.Format()

	//Raw Wav data starting point
	rawData := data[44:]

	bytesSampleSize := int(wavFormat.BitsPerSample) / 8

	samples := []wav.Sample{}

	for i := 0; i < len(rawData); i += bytesSampleSize {
		rawSample := rawData[i : i+bytesSampleSize]
		sampleValue := binary.BigEndian.Uint16(rawSample)
		samples = append(samples, wav.Sample{Values: [2]int{int(sampleValue), int(sampleValue)}})
	}

	out2, err := os.Create("answers/chal19_answer.wav")

	if err != nil {
		panic(err)
	}

	defer out2.Close()

	wav2Writer := wav.NewWriter(out2, uint32(len(samples)), wavFormat.NumChannels, wavFormat.SampleRate, wavFormat.BitsPerSample)
	wav2Writer.WriteSamples(samples)

	fmt.Println("Answer saved to: " + out2.Name())
}
