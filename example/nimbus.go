package main

import (
	"time"

	"github.com/fcangialosi/uiprogress"
	log "github.com/sirupsen/logrus"
)

func main() {
	progress := uiprogress.New()
	progress.SetRefreshInterval(time.Millisecond * 500)
	t := 5
	deadline := time.Now().Add(time.Duration(t) * time.Second)

	bar := progress.AddBar(deadline, t*1000)
	bar.PrependSecRemaining()
	bar.AppendOtherBytes()

	bytes_received := 0
	progress.Start()
	for i := 1; i <= (t * 10); i++ {
		bytes_received += 102400
		bar.Set(i*100, bytes_received)
		time.Sleep(100 * time.Millisecond)
	}
	progress.Stop()
	log.Info("stuff")
}
