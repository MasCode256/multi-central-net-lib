package common

import "time"

func CompareTime(timestamp string, referenceTime time.Time) (bool, error) {
	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return false, err
	}

	now := time.Now()
	elapsed := now.Sub(parsedTime)

	// Сравниваем прошедшее время с заданным значением
	return elapsed > referenceTime.Sub(parsedTime), nil
}