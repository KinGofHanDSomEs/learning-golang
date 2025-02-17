package main

import (
	"errors"
	"fmt"
	"sort"
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Company struct {
	Workers []Worker
}

type Worker struct {
	Name            string
	Position        string
	Salary          uint
	ExperienceYears uint
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if _, err := positions[position]; !err {
		return errors.New("invalid position")

	}
	c.Workers = append(c.Workers, Worker{name, position, salary, experience})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	sort.Slice(c.Workers, func(i, j int) bool {
		amount1, amount2 := c.Workers[i].Salary*c.Workers[i].ExperienceYears, c.Workers[j].Salary*c.Workers[j].ExperienceYears
		if amount1 != amount2 {
			return amount1 > amount2
		}
		pos1, pos2 := positions[c.Workers[i].Position], positions[c.Workers[j].Position]
		if pos1 != pos2 {
			return pos1 < pos2
		}
		return true
	})
	result := []string{}
	for _, worker := range c.Workers {
		result = append(result, fmt.Sprintf("%s - %d - %s", worker.Name, worker.Salary*worker.ExperienceYears, worker.Position))
	}
	return result, nil
}

var positions = map[string]int{
	"директор":       1,
	"зам. директора": 2,
	"начальник цеха": 3,
	"мастер":         4,
	"рабочий":        5,
}

func main() {
	company := Company{[]Worker{
		{Name: "Михаил", Position: "директор", Salary: 200, ExperienceYears: 5},
		{Name: "Игорь", Position: "зам. директора", Salary: 180, ExperienceYears: 3},
		{Name: "Николай", Position: "начальник цеха", Salary: 120, ExperienceYears: 2},
		{Name: "Андрей", Position: "мастер", Salary: 90, ExperienceYears: 10},
		{Name: "Виктор", Position: "рабочий", Salary: 80, ExperienceYears: 3},
	}}
	result, _ := company.SortWorkers()
	fmt.Print(result)
}

//[Михаил - 12000 - директор Андрей - 10800 - мастер Игорь - 6480 - зам. директора Николай - 2880 - начальник цеха Виктор - 2880 - рабочий],
// [Михаил - 144000 - директор Андрей - 129600 - мастер Игорь - 77760 - зам. директора Николай - 34560 - начальник цеха Виктор - 34560 - рабочий]
