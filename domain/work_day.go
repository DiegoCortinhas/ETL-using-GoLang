package domain

import "time"

type WorkDay struct{
	Employee	Employee
	Date		time.Time
	Punches		[] time.Time
}