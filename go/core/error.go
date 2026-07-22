package core

type BusinessDayCalculatorError struct {
	IsBusinessDayCalculatorError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewBusinessDayCalculatorError(code string, msg string, ctx *Context) *BusinessDayCalculatorError {
	return &BusinessDayCalculatorError{
		IsBusinessDayCalculatorError: true,
		Sdk:              "BusinessDayCalculator",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *BusinessDayCalculatorError) Error() string {
	return e.Msg
}
