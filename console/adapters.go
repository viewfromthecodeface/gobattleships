package console

import (
	"bufio"
	"fmt"
	"os"
)

type ConsoleOutput struct {
	// no fields
}

func NewConsoleOutput() ConsoleOutput {
	return ConsoleOutput{}
}

func (o ConsoleOutput) Show(text string) {
	fmt.Println(text)
}

type KeyboardPlayerInput struct {
	scanner *bufio.Scanner
}

func NewKeyboardPlayerInput() *KeyboardPlayerInput {
	scanner := bufio.NewScanner(os.Stdin)
	return &KeyboardPlayerInput{scanner: scanner}
}

func (k *KeyboardPlayerInput) Fetch() string {
	k.scanner.Scan()
	return k.scanner.Text()
}
