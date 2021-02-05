package etl

import (
	"Transformer/etl/domain"
	"Transformer/etl/vo"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
	"time"
)

func convert(entries [] vo.TimeEntry) [] domain.WorkDay {

	var result [] domain.WorkDay
	var employees = map[string]domain.Employee {}

	for i := 0; i < len(entries); i++ {

		employeeString := entries[i].Employee
		dateString := entries[i].Date + "/2019"
		hoursString := entries[i].Hours

		employee := employees[employeeString]

		if employee == (domain.Employee{}) {
			employee = domain.Employee{Name: employeeString}
			employees[employeeString] = employee
		}

		layout := "02/01/2006"
		date, _ := time.Parse(layout, dateString)

		var punches [] time.Time
		punchesArray := strings.Split(hoursString, " ")

		for i := 0; i < len(punchesArray); i++ {

			layout := "15:04"
			punch, _ := time.Parse(layout, punchesArray[i])
			punches = append(punches, punch)
		}

		workDay := domain.WorkDay{Employee: employee, Date: date, Punches: punches}
		result = append(result, workDay)

		logrus.Debugln(workDay)
	}
	log.Printf("Foram transformados: %v\n", len(entries))
	return result
}