package pkg

import (
	"fmt"
)

func RemoveLinesUp(lines int) {

	for i := 0; i < lines; i++ {
		fmt.Print("\033[1A\033[2K")
	}
}

func MaxLen(Items []string) (max int) {

	for _, l := range Items {

		if len(l) > max {
			max = len(l)
		}
	}

	return

}
