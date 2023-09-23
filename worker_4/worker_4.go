package worker_4

type Worker_4 struct {
	position string
	salary   int
	address  string
}

func NewWorker_4(position string, salary int, address string) *Worker_4 {
	return &Worker_4{
		position: position,
		salary:   salary,
		address:  address}
}
func (w *Worker_4) GetPosition() string {
	return w.position
}
func (w *Worker_4) SetPosition(position string) {
	w.position = position
}
func (w *Worker_4) GetSalary() int {
	return w.salary
}
func (w *Worker_4) SetSalary(salary int) {
	w.salary = salary
}
func (w *Worker_4) GetAddress() string {
	return w.address
}
func (w *Worker_4) SetAddress(address string) {
	w.address = address
}
