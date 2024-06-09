package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	gops "github.com/mitchellh/go-ps"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд
*/

// Выводим в константы все нужные команды
const (
	CD   = "cd"
	PWD  = "pwd"
	ECHO = "echo"
	KILL = "kill"
	PS   = "ps"
	EXEC = "exec"
	FORK = "fork"
	QUIT = "\\q"
)

// Определяем какая команда была введена
func processCommand(s string) {
	cmd := strings.Split(s, " ")
	switch cmd[0] {
	case CD:
		cd(cmd)
	case PWD:
		pwd()
	case ECHO:
		echo(cmd)
	case KILL:
		kill(cmd)
	case PS:
		ps()
	case EXEC:
		execute(cmd)
	case FORK:
		forkCommand(cmd)
	case QUIT:
		os.Exit(0)
	}
}

func cd(cmd []string) {
	if len(cmd) == 1 {
		home := os.Getenv("HOME")
		_ = os.Chdir(home)
	}
	if len(cmd) == 2 {
		err := os.Chdir(cmd[1])
		if err != nil {
			fmt.Printf("cd: no such file or directory: %s", cmd[1])
		}
	}
}

func pwd() {
	curr, _ := os.Getwd()
	fmt.Println(curr)
}

func echo(cmd []string) {
	fmt.Println(strings.Join(cmd[1:], " "))
}

func kill(cmd []string) {
	procId, _ := strconv.Atoi(cmd[1])
	process, err := os.FindProcess(procId)
	if err != nil {
		fmt.Printf("kill: %s failed: no such process", cmd[1])
	}

	err = process.Kill()
	if err != nil {
		fmt.Printf("kill: %s failed: %s", cmd[1], err.Error())
	}
}

func ps() {
	processes, _ := gops.Processes()
	for _, p := range processes {
		fmt.Printf("name: %-15s\tpid: %d\n", p.Executable(), p.Pid())
	}
}

func execute(args []string) {
	if len(args) == 0 {
		fmt.Println("exec: no command specified")
	}

	cmd := exec.Command(args[1], args[2:]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func forkCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("fork: enter process to fork")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.New("fork: couldn't get pwd")
	}

	cmd, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf("fork: couldn't find path for %v: %w", args[0], err)
	}
	if cmd == "" {
		return "", fmt.Errorf("fork: couldn't find path for %v", args[0])
	}
	args[0] = cmd

	_, err = syscall.ForkExec(args[0], args, &syscall.ProcAttr{
		Dir:   pwd,
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}, // print message to the same pty
	})

	if err != nil {
		return "", fmt.Errorf("fork: could't fork: %w", err)
	}

	return "", nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		processCommand(sc.Text())
	}
}
