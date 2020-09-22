package helpers

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func GetAlphabetChar(i int) string {
	return string(rune(int('A') + i))
}

func GetCharNumber(i string) int {
	return int([]rune(i)[0]) - int('A')
}
