package service

import (
	"day01_ManagementSystem/internal/model"
	"fmt"
)

type StudentService struct {
	students []model.Student
	nextID   int
}

func NewStudentService() *StudentService {
	return &StudentService{
		students: []model.Student{},
		nextID:   1,
	}
}

func (s *StudentService) AddStudent(name string, age int) (model.Student, error) {
	if !checkoutName(name) {
		return model.Student{}, fmt.Errorf("[AddStudent] Name error:%v", name)
	}
	if !checkoutAge(age) {
		return model.Student{}, fmt.Errorf("[AddStudent] Age error:%v", age)
	}
	student := model.Student{
		ID:   s.nextID,
		Name: name,
		Age:  age,
	}
	s.students = append(s.students, student)
	s.nextID++
	return student, nil
}

func (s *StudentService) DeleteStudentByID(id int) (bool, error) {
	for i, stu := range s.students {
		if stu.ID == id {
			s.students = append(s.students[:i], s.students[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("[DeleteStudentByID] Can't find student by id: %d...", id)
}

func (s *StudentService) UpdateStudentByID(id int, name string, age int) (bool, error) {
	if !checkoutName(name) {
		return false, fmt.Errorf("[UpdateStudentByID] Name error:%v", name)
	}
	if !checkoutAge(age) {
		return false, fmt.Errorf("[UpdateStudentByID] Age error:%v", age)
	}
	for i, stu := range s.students {
		if stu.ID == id {
			if name != "" {
				s.students[i].Name = name
			}
			if age > 0 {
				s.students[i].Age = age
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("[UpdateStudentByID] Can't find student by id: %d...", id)
}

func (s *StudentService) ListStudents() ([]model.Student, error) {
	return s.students, nil
}

func (s *StudentService) GetStudentByID(id int) (model.Student, error) {
	for _, stu := range s.students {
		if stu.ID == id {
			return stu, nil
		}
	}
	return model.Student{}, fmt.Errorf("[GetStudentByID] Can't find student by id: %d", id)
}

func checkoutName(name string) bool {
	if len(name) <= 0 {
		return false
	}
	return true
}

func checkoutAge(age int) bool {
	if age > 0 && age < 150 {
		return true
	}
	return false
}
