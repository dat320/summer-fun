package main

// Task 7: Simple Shell
//
// This task focuses on building a simple shell that accepts
// commands that run certain OS functions or programs. For OS
// functions refer to golang's built-in OS and ioutil packages.
//
// The shell should be implemented through a command line
// application; allowing the user to execute all the functions
// specified in the task. Info such as [path] are command arguments
//
// Important: The prompt of the shell should print the current directory.
// For example, something like this:
//   /Users/meling/Dropbox/work/opsys/2020/meling-stud-labs/lab3>
//
// We suggest using a space after the > symbol.
//
// Your program should be able to at least the following functions:
// 	- exit
// 		- exit the program
// 	- cd [path]
// 		- change directory to a specified path
// 	- ls
// 		- list items and files in the current path
// 	- mkdir [path]
// 		- create a directory with the specified path
// 	- rm [path]
// 		- remove a specified file or folder
// 	- create [path]
// 		- create a file with a specified name
// 	- cat [file]
// 		- show the contents of a specified file
// 			- any file, you can use the 'hello.txt' file to check if your
// 			  implementation works
// 	- help
// 		- show a list of available commands
//
// You may also implement any number of optional functions, here are some ideas:
// 	- help [command]
// 		- give additional info on a certain command
// 	- ls [path]
// 		- make ls allow for a specified path parameter
// 	- rm -r
// 		WARNING: Be aware of where you are when you try to execute this command
// 		- recursively remove a directory
// 			- meaning that if the directory contains files, remove
// 			  all the files within the directory first, then the
// 			  directory itself
// 	- calc [expression]
// 		- Simple calculator program that can calculate a given expression
// 			- example expressions could be + - * \ pow
// 	- ipconfig
// 		- show ip interfaces
// 	- history
// 		- show command history
// 		- Alternatively implement this together with pressing up on your
// 		  keypad to load the previous command
// 		- clrhistory to clear history
// 	- tail [n]
// 		- show last n lines of a file
// 	- head [n]
// 		- show first n lines of a file
// 	- writefile [text]
// 		- write specified text to a specified file
//
// 	Or, alternatively, implement your own functionality not specified as you please
//
// Additional notes:
// 	- If you want to use colors in your terminal program you can see the package
// 		"github.com/fatih/color"
//
// 	- Helper functions may lead to cleaner code
//

// Terminal contains
type Terminal struct {
	// TODO (student): Add field(s) if necessary
}

// Execute executes a given command
func (t *Terminal) Execute(command string) {
	// TODO (student): Make this run a given command
}

// This is the main function of the application.
// User input should be continuously read and checked for commands
// for all the defined operations.
// See https://golang.org/pkg/bufio/#Reader and especially the ReadLine
// function.
func main() {

}
