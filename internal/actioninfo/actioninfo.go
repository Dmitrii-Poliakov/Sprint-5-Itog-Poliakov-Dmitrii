package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(v string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println("parsing error:", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("format error:", err)
			continue
		}
		fmt.Println(info)
	}
}
