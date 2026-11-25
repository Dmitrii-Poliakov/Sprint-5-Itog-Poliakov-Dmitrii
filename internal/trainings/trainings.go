package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitStr := strings.Split(datastring, ",")
	if len(splitStr) != 3 {
		return errors.New("the slice length should be 3")
	}
	stepsCount, err := strconv.Atoi(splitStr[0])
	if err != nil {
		return errors.New("data type for steps is incorrect (should be int)")
	} else {
		t.Steps = stepsCount
	}
	if stepsCount <= 0 {
		return errors.New("count of steps should be > 0")
	}
	t.TrainingType = splitStr[1]
	timeDuration, err := time.ParseDuration(splitStr[2])
	if err != nil {
		return errors.New("data type for duration is incorrect (should be time.Duration)")
	} else {
		t.Duration = timeDuration
	}
	if timeDuration <= 0 {
		return errors.New("duration should be > 0")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	avrSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calcCalories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calcCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calcCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}

	if err != nil {
		return "", errors.New("calculation error")
	}

	hours := t.Duration.Hours()

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		hours,
		distance,
		avrSpeed,
		calcCalories,
	)
	return result, nil
}
