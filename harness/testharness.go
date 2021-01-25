package harness

import (
	"sync"

	qrcode "github.com/skip2/go-qrcode"
)

func Run() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go testRun(wg, test)
	}

	wg.Wait()
}

func testRun(wg *sync.WaitGroup, f func()) {
	defer wg.Done()
	f()
}

func test() {
	qrcode.Encode("this is a test string", qrcode.Medium, 512)
}
