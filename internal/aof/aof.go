package aof

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type Aof struct {
	file *os.File
	rd   *bufio.Reader
	mu   sync.Mutex
}

func NewAof(path string) (*Aof, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	aof := &Aof{
		file: f,
		rd:   bufio.NewReader(f),
	}

	// start go routine to sync aof to disk every 1 second
	go func() {
		for {
			aof.mu.Lock()

			aof.file.Sync()

			aof.mu.Unlock()

			time.Sleep(time.Second)
		}
	}()

	return aof, nil
}

func (aof *Aof) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	return aof.file.Close()
}

func (aof *Aof) Write(respValue []byte) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()
	print(string(respValue))
	_, err := aof.file.Write(respValue)
	if err != nil {
		return err
	}

	return nil
}

// func (aof *Aof) Read(fn func(value resp.Value)) error {
// 	aof.mu.Lock()
// 	defer aof.mu.Unlock()

// 	aof.file.Seek(0, io.SeekStart)

// 	reader := resp.NewResp(aof.file)

// 	for {
// 		value, err := reader.Read()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}

// 			return err
// 		}

// 		fn(value)
// 	}

// 	return nil
// }
