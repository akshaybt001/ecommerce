package controllers

import "fmt"

func GetDate(year, month, week, day int) string {
	if day != 0 && month != 0 && year != 0 {
		return fmt.Sprintf("%d-%d-%d", year,month,day)
	}else if week !=0 && month !=0 && year !=0{
		var weekDay1,weekDay2 int
		if  week<5{
			weekDay1,weekDay2=(week-1)*7+1, week*7
		} else if week==5 {
			weekDay1,weekDay2=(week-1)*7+1,31
		}
		return fmt.Sprintf("'%d-%d-%d 00:00:00'::timestamp AND '%d-%d-%d 23:59:59'::timestamp", year, month, weekDay1, year, month, weekDay2)
	}
	return ""
}