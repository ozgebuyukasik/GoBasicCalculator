package calc

type Calculation interface {
	Calculate() float64
}

type Addition struct {
	firstOperand  float64
	secondOperand float64
}
type Subtraction struct {
	firstOperand  float64
	secondOperand float64
}
type Multiplication struct {
	firstOperand  float64
	secondOperand float64
}
type Division struct {
	firstOperand  float64
	secondOperand float64
}

func (add Addition) Calculate() float64 {

	return add.firstOperand + add.secondOperand
}
func (sub Subtraction) Calculate() float64 {

	return sub.firstOperand + sub.secondOperand
}
func (mul Multiplication) Calculate() float64 {

	return mul.firstOperand + mul.secondOperand
}
func (div Division) Calculate() float64 {

	return div.firstOperand + div.secondOperand
}
