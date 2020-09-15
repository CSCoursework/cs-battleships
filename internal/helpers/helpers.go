package helpers

import (
	"os"
	"os/exec"
)

func ClearConsole() {
	out, _ := exec.Command("cls").Output() // Windows only
	_, _ = os.Stdout.Write(out)
}

func GetAlphabetChar(i int) string {
	return string(rune(int('A') + i))
}
