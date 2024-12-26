package repeatable

import "time"

// это утилита, которая принимает в себя функцию, количество попыток и время на попытку.
// нужна для того, чтобы делать к примеру попытки подключения к БД

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}
