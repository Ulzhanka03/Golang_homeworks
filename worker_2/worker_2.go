package worker_2

type Worker_2 struct {
	position string
	salary   int
	address  string
}

func NewWorker_2(position string, salary int, address string) *Worker_2 {
	return &Worker_2{
		position: position,
		salary:   salary,
		address:  address}
}
func (w *Worker_2) GetPosition() string {
	return w.position
}
func (w *Worker_2) SetPosition(position string) {
	w.position = position
}
func (w *Worker_2) GetSalary() int {
	return w.salary
}
func (w *Worker_2) SetSalary(salary int) {
	w.salary = salary
}
func (w *Worker_2) GetAddress() string {
	return w.address
}
func (w *Worker_2) SetAddress(address string) {
	w.address = address
}
