package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("0.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1) // Вернем ненулевой код выхода при ошибке

	}

	fmt.Println("Точное время:", ntpTime.Format(time.RFC850))
}
