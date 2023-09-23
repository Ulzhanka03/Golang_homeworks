package worker_1

type Worker_1 struct {
	position string
	salary   int
	address  string
}

func NewWorker_1(position string, salary int, address string) *Worker_1 {
	return &Worker_1{
		position: position,
		salary:   salary,
		address:  address}
}
func (w *Worker_1) GetPosition() string {
	return w.position
}
func (w *Worker_1) SetPosition(position string) {
	w.position = position
}
func (w *Worker_1) GetSalary() int {
	return w.salary
}
func (w *Worker_1) SetSalary(salary int) {
	w.salary = salary
}
func (w *Worker_1) GetAddress() string {
	return w.address
}
func (w *Worker_1) SetAddress(address string) {
	w.address = address
}
