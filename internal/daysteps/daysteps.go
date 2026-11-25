package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitStr := strings.Split(datastring, ",")
	if len(splitStr) != 2 {
		return errors.New("the slice length should be 2")
	}
	stepsCount, err := strconv.Atoi(splitStr[0])
	if err != nil {
		return errors.New("data type for steps is incorrect (should be int)")
	} else {
		ds.Steps = stepsCount
	}
	if stepsCount <= 0 {
		return errors.New("count of steps should be > 0")
	}
	timeDuration, err := time.ParseDuration(splitStr[1])
	if err != nil {
		return errors.New("data type for duration is incorrect (should be time.Duration)")
	} else {
		ds.Duration = timeDuration
	}
	if timeDuration <= 0 {
		return errors.New("duration should be > 0")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	var calcCalories float64
	var err error

	calcCalories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", errors.New("calculation error")
	}
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calcCalories,
	)
	return result, nil
}
