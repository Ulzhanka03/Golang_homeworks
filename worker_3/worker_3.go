package worker_3

type Worker_3 struct {
	position string
	salary   int
	address  string
}

func NewWorker_3(position string, salary int, address string) *Worker_3 {
	return &Worker_3{
		position: position,
		salary:   salary,
		address:  address}
}
func (w *Worker_3) GetPosition() string {
	return w.position
}
func (w *Worker_3) SetPosition(position string) {
	w.position = position
}
func (w *Worker_3) GetSalary() int {
	return w.salary
}
func (w *Worker_3) SetSalary(salary int) {
	w.salary = salary
}
func (w *Worker_3) GetAddress() string {
	return w.address
}
func (w *Worker_3) SetAddress(address string) {
	w.address = address
}
