package gopractice

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func goChunkReader() {

	fptr := flag.String("/tmp/tmp.dhxPCM6SHO/", "demo2.log", "/tmp/tmp.dhxPCM6SHO/demo2.log")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		fmt.Println(err)
	}

	h, _ := ReadChunk(f, 500)

	v, _ := ReadLogs(h)

	fmt.Println("=======>", v)

	fmt.Println(":============================:")

	fmt.Println(h)
}

// func ReadChunk(file *os.File, chunkSize int) ([]string, error) {
// 	chunk := make([]byte, chunkSize)
// 	_, err := file.Read(chunk)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var lines []string
// 	reader := bytes.NewReader(chunk)
// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	fmt.Println("===>", chunk)
// 	fmt.Println("=======>", lines)
// 	return lines, scanner.Err()

// }

func ReadChunk(file *os.File, chunkSize int) ([]byte, error) {
	chunk := make([]byte, chunkSize)

	var wg sync.WaitGroup

	n, err := file.Read(chunk)
	if n == 0 {
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if err == io.EOF {
			return nil, err
		}
		wg.Add(1)

		return nil, err
	}

	wg.Wait()
	return chunk, nil
}

func ReadLogs(bytefile []byte) ([]string, error) {
	var lines []string
	reader := bytes.NewReader(bytefile)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
