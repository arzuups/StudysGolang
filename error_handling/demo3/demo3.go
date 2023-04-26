package error_handling

import (
	"errors"
	"fmt"
)

func predict(predict int) (string, error) {
	
	numberInMind := 80
	
	if predict < 1 || predict > 100 {
	    return "", errors.New("Enter a number between 1 and 100:")
        }
	
	if predict > numberInMind {
	    return "Enter a smaller number :", nil
        }
	
	return "You know!", nil
}

func
