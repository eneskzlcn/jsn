package clipboard

import (
	"fmt"
	"os/exec"
)

func Add(str string) error {
	return exec.Command("sh", "-c",
		fmt.Sprintf(`echo '%s' | pbcopy`, str)).Run()
}
