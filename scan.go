package ff

import (
	"bufio"
	"os"
	"time"
)

func ScanStdin(separator []byte, dataIn chan []byte, eof chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0), 10*MB)

	data := []byte{}
	scanErrCount := 0
	for {
		if scanErrCount >= 5 {
			eof <- true
			return
		} else if !scanner.Scan() || scanner.Err() != nil {
			scanErrCount++
			time.Sleep(time.Millisecond * 200)
		}
		data = append(data, scanner.Bytes()...)
		data = append(data, byte('\n'))
		if len(data) > len(separator) && string(data[len(data)-len(separator):]) == string(separator) {
			dataIn <- data
			data = []byte{}
		}
	}
}

func ScanStdinBytes() []byte {
	if _, err := os.Stdin.Stat(); err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		data := []byte{}
		for scanner.Scan() {
			data = append(data, scanner.Bytes()...)
			data = append(data, '\n')
		}
		return data
	}
	return []byte{}
}
