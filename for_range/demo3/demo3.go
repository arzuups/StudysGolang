package for_range

import (
	"fmt"
)

func Demo3() {

	dictionary := map[string]string{"Bowl": "Kase", "Plate": "Tabak"}

	for k, v := range dictionary {
		fmt.Print(k)
		fmt.Println(v)
	}

}

