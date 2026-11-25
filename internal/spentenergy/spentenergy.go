package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps: 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight: 0")
	}
	if height <= 0 {
		return 0, errors.New("height: 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration: 0")
	}
	avrSpeed := MeanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	calcCalories := (weight * avrSpeed * durationMin) / minInH
	calcSpentCalories := calcCalories * walkingCaloriesCoefficient
	return calcSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps: 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight: 0")
	}
	if height <= 0 {
		return 0, errors.New("height: 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration: 0")
	}
	avrSpeed := MeanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	calcSpentCalories := (weight * avrSpeed * durationMin) / minInH
	return calcSpentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	avrSpeed := distance / duration.Hours()
	return avrSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := (float64(steps) * (height * stepLengthCoefficient)) / float64(mInKm)
	return stepLength
}
