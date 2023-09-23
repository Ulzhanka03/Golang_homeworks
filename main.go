package main

import (
	"Golang_homeworks/worker_1"
	"Golang_homeworks/worker_2"
	"Golang_homeworks/worker_3"
	"Golang_homeworks/worker_4"
	"Golang_homeworks/worker_5"
	"fmt"
)

func main() {
	for_worker_1 := worker_1.NewWorker_1("Middle UX/UI designer", 300000, "Makataeva 131")

	fmt.Println("Worker 1 Position:", for_worker_1.GetPosition())
	fmt.Println("Worker 1 Salary:", for_worker_1.GetSalary())
	fmt.Println("Worker 1 Address:", for_worker_1.GetAddress())

	for_worker_1.SetPosition("Senior UX/UI designer")
	for_worker_1.SetSalary(450000)
	for_worker_1.SetAddress("Tole bi 59")

	fmt.Println("New position:", for_worker_1.GetPosition())
	fmt.Println("New salary:", for_worker_1.GetSalary())
	fmt.Println("New address:", for_worker_1.GetAddress())

	for_worker_2 := worker_2.NewWorker_2("Junior Web Developer", 180000, "Abaya 34")

	fmt.Println("Worker 2 Position:", for_worker_2.GetPosition())
	fmt.Println("Worker 2 Salary:", for_worker_2.GetSalary())
	fmt.Println("Worker 2 Address:", for_worker_2.GetAddress())

	for_worker_2.SetPosition("Middle Web Developer")
	for_worker_2.SetSalary(250000)
	for_worker_2.SetAddress("Aimanova 126")

	fmt.Println("New position:", for_worker_2.GetPosition())
	fmt.Println("New salary:", for_worker_2.GetSalary())
	fmt.Println("New address:", for_worker_2.GetAddress())

	for_worker_3 := worker_3.NewWorker_3("Senior Front-end", 420000, "Turgut Ozala 307/1")

	fmt.Println("Worker 3 Position:", for_worker_3.GetPosition())
	fmt.Println("Worker 3 Salary:", for_worker_3.GetSalary())
	fmt.Println("Worker 3 Address:", for_worker_3.GetAddress())

	for_worker_3.SetPosition("Lead Front-end")
	for_worker_3.SetSalary(530000)
	for_worker_3.SetAddress("Baitursynova 32")

	fmt.Println("New position:", for_worker_3.GetPosition())
	fmt.Println("New salary:", for_worker_3.GetSalary())
	fmt.Println("New address:", for_worker_3.GetAddress())

	for_worker_4 := worker_4.NewWorker_4("Network manager", 280000, "Zhumanova 110/2")

	fmt.Println("Worker 4 Position:", for_worker_4.GetPosition())
	fmt.Println("Worker 4 Salary:", for_worker_4.GetSalary())
	fmt.Println("Worker 4 Address:", for_worker_4.GetAddress())

	for_worker_4.SetPosition("Network Administrator")
	for_worker_4.SetSalary(330000)
	for_worker_4.SetAddress("Saina 12B")

	fmt.Println("New position:", for_worker_4.GetPosition())
	fmt.Println("New salary:", for_worker_4.GetSalary())
	fmt.Println("New address:", for_worker_4.GetAddress())

	for_worker_5 := worker_5.NewWorker_5("Project manager", 360000, "Raiymbeka 44")

	fmt.Println("Worker 5 Position:", for_worker_5.GetPosition())
	fmt.Println("Worker 5 Salary:", for_worker_5.GetSalary())
	fmt.Println("Worker 5 Address:", for_worker_5.GetAddress())

	for_worker_5.SetPosition("Project Director")
	for_worker_5.SetSalary(480000)
	for_worker_5.SetAddress("Rozybakieva 247a")

	fmt.Println("New position:", for_worker_5.GetPosition())
	fmt.Println("New salary:", for_worker_5.GetSalary())
	fmt.Println("New address:", for_worker_5.GetAddress())

}
