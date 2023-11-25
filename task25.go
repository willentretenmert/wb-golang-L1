package main

import "time"

// Таймер (sleep) реализую через проверку времени в данном моменте, когда время совпадет с переданным, выполнение программы продолжится
func task25(d time.Duration) {
	endTime := time.Now().Add(d)
	for time.Now().Before(endTime) {
	}
}
