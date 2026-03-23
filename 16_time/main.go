package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	t := time.Date(2026, time.March, 23, 15, 4, 5, 0, time.UTC)
	fmt.Println(t)

	t1 := time.Now()
	t2 := t1.Add(2 * time.Hour)
	fmt.Println(t1.Before(t2)) // true
	fmt.Println(t1.After(t2))  // false
	fmt.Println(t1.Equal(t2))  // false

	t3 := t.Add(24 * time.Hour) // добавить сутки
	t4 := t.AddDate(1, 0, 0)    // добавить 1 год
	fmt.Println(t3, t4)

	year, month, day := t.Date()
	hour, minuts, sec := t.Clock()
	unix := t.Unix() // секунды с 1970-01-01
	unixNano := t.UnixNano()

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))      // 2026-03-23 15:04:05
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006")) // Mon Mar 23 15:04:05 2026
	fmt.Println(t.Format("02/01/2006 03:04:05 PM"))   // 23/03/2026 03:04:05 PM

	timer := time.NewTimer(2 * time.Second)
	<-timer.C // блокируется до истечения времени
	fmt.Println("прошло 2 секунды")

	time.Sleep(2 * time.Second) // тоже блокирует, но не даёт возможности отменить

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("тик") // Каждую секунду
	}
}
