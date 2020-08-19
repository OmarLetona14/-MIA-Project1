package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"MIA/DiskMagnament/function"
)

var equalizer string = "->"
var com string = ""

func main() {
	fmt.Println("Welcome to the console! (Press x to finish)")
	reader := bufio.NewReader(os.Stdin)
	finish_app := false
	for !finish_app {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = get_text(input)
		if strings.ToLower(input) != "x"{
			if !strings.HasPrefix(input, "#") { 
				execute_console(input)
			}
		} else {
			fmt.Println("Finishing the app...")
			finish_app = true
		}
	}
}

func execute_console(i string) {
	fmt.Println(i)
	com += i
	if(!strings.HasSuffix(i, "/*")){
		recognize_command(splitter(get_text(com)))
		com = ""
	}else{
		com = strings.TrimRight(com, "/*")
		fmt.Println(com)
	}
}

func get_text(txt string) string {
	if runtime.GOOS == "windows" {
		txt = strings.TrimRight(txt, "\r\n")
	} else {
		txt = strings.TrimRight(txt, "\n")
	}
	return txt
}

func recognize_command(commands []string) {
	switch strings.ToLower(commands[0]) {
	case "mkdisk":
		function.Exec_mkdisk(commands)
	case "exec":
		sub_command := strings.Split(commands[1], equalizer)
		fmt.Println(strings.ToLower(sub_command[1]))
		if strings.ToLower(sub_command[0]) == "-path" {
			readFile(sub_command[1])
		} else {
			fmt.Println("Not supported command! ")
			fmt.Println("You may say -path, press -help to see the list of commands avalibles")
		}
	case "rmdisk":
		function.Exec_mrdisk(commands)
	case "fdisk":
		function.Exec_fdisk(commands)
	case "pause":
		fmt.Print("Exection paused \nPress any key to continue... ")
		reader := bufio.NewReader(os.Stdin)
		x, _ := reader.ReadString('\n')
		x += ""
	default:
		fmt.Println("Not supported command! ")
	}
}

func readFile(file_name string) {
	f, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var command string =""
	for scanner.Scan() {
		command += strings.TrimRight(scanner.Text(), " ")
		if !strings.HasPrefix(command, "#"){
			if(!strings.HasSuffix(command, "/*")){
				fmt.Println("Executing ", command, "... ")
				execute_console(command)
				command = ""
			}else{
				command = strings.TrimRight(command, "/*")
				fmt.Println(command)
			}
		}else{
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func splitter(txt string) []string {
	commands := strings.Split(txt, " ")
	return commands
}

