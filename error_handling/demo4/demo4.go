package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message string
}

func (b borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func Predict2(predict int) (string, error) {
	if predict < 1 || predict > 100 {
	
	
