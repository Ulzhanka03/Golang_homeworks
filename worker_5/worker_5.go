package worker_5

type Worker_5 struct {
	position string
	salary   int
	address  string
}

func NewWorker_5(position string, salary int, address string) *Worker_5 {
	return &Worker_5{
		position: position,
		salary:   salary,
		address:  address}
}
func (w *Worker_5) GetPosition() string {
	return w.position
}
func (w *Worker_5) SetPosition(position string) {
	w.position = position
}
func (w *Worker_5) GetSalary() int {
	return w.salary
}
func (w *Worker_5) SetSalary(salary int) {
	w.salary = salary
}
func (w *Worker_5) GetAddress() string {
	return w.address
}
func (w *Worker_5) SetAddress(address string) {
	w.address = address
}
