package models

import "time"

/*Dates structure of all management of dates of lisander*/
type Dates struct {
	UpdateDate      int       `bson:"updateDate" json:"updateDate"`
	UpdateDateToday int       `bson:"updateDateToday" json:"updateDateToday"`
	DateFrom        time.Time `bson:"dateFrom" json:"dateFrom"`
}
