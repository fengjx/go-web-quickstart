package utils

import "time"

var Now = time.Now()

func init() {

	go func() {
		for {
			Now = time.Now()
			time.Sleep(time.Millisecond * 50)
		}
	}()

}
