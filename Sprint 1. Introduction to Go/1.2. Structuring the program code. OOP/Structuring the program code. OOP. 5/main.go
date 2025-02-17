package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	reportID, _ := strconv.Atoi(strings.ReplaceAll(reportDate, "-", ""))
	return Report{
		User:     user,
		ReportID: reportID,
		Date:     reportDate,
	}
}

func PrintReport(report Report) {
	fmt.Printf("Идентификатор: %d\nИмя: %s\nEmail: %s\nВозраст: %d\nДата создания отчета: %s\n", report.ID, report.Name, report.Email, report.Age, report.Date)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var res []Report
	for _, user := range users {
		res = append(res, CreateReport(user, reportDate))
	}
	return res
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 28},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 34},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 25},
	}

	// Указываем дату для отчета
	reportDate := "2024-11-14"

	// Генерация отчетов для всех пользователей
	reports := GenerateUserReports(users, reportDate)

	// Вывод отчетов
	for _, report := range reports {
		PrintReport(report)
	}
}
