package register

import (
	"fmt"
	"os"
)

func Exec() error {
	d, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Print(d)

	return nil
}
