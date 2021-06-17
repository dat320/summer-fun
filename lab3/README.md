# Lab 3: Introduction to Go Programming

| Lab 3: | Introduction to Go Programming |
| ---------------------    | --------------------- |
| Subject:                 | DAT320 Operating Systems and Systems Programming |
| Grading:                 | Pass/fail |
| Submission:              | Individually |

## Table of Contents

1. [Introduction](#introduction)
2. [Go Resources](#go-resources)
3. [Primary Resources](#primary-resources)
4. [Additional Resources](#additional-resources)
5. [Installing Go](#installing-go)
6. [Tasks](#tasks)
7. [Task 1: Exercises from Tour of Go](#task-1-exercises-from-tour-of-go)
8. [Task 2: Go Language Questions](#task-2-go-language-questions)
9. [Task 3: Go Exercises](#task-3-go-exercises)

## Introduction

In this lab you will be learning the Go programming language, which we will be using throughout the course.
Go is a very nice language in many ways.
It provides built in primitives to design concurrent programs using lightweight threads, called goroutines.
In addition, these goroutines can communicate using channels, instead of by means of shared memory which is the most common approach in other languages such as C and Java.
Usage of goroutines and channels will be covered in depth in later labs.
For now you just need to know that they exist.
Go is also a very small language in the sense that it has very few keywords, and as such it does not take too long to learn the fundamentals of the language.

## Go Resources

Below we provide some resources which will be useful when you start programming in Go.

### Primary Resources

[The Go Language web page](https://golang.org/) provides a variety of helpful documentation, including:

* Learning the basics with ["A Tour of Go"](http://tour.golang.org/)
* [The Go documentation](https://golang.org/doc/)
* [Standard library package documentation](https://golang.org/pkg/)

### Additional Resources

* [Effective Go](https://golang.org/doc/effective_go.html) gives tips for writing clear, idiomatic Go code
* [Frequently Asked Questions](https://golang.org/doc/faq)
* [The Go Blog](https://blog.golang.org/index) contains many great articles that describe important idioms and intricacies of the language.
* [Collection of Videos about Go](https://github.com/golang/go/wiki/GoTalks)
* [The Go Programming Language (book)](http://www.gopl.io)
* [Golang Tutorial Series](https://golangbot.com/learn-golang-series/)

**Troubleshooting tip:**
When searching for information about Go, use the term *golang* instead of go.

## Installing Go

Please refer to [our guide](https://github.com/dat320-2020/course-info/blob/master/setup-go.md) for installing Go.
 and setting up a text editor/IDE for development.

## Tasks

### Task 1: Exercises from Tour of Go

Start learning the basics of Go by completing ["A Tour of Go"](http://tour.golang.org/).
You should do at least the following exercises.

* [Exercise: Loops and Functions](https://tour.golang.org/flowcontrol/8)
* [Exercise: Slices](https://tour.golang.org/moretypes/18)
* [Exercise: Maps](https://tour.golang.org/moretypes/23)
* [Exercise: Errors](https://tour.golang.org/methods/20)
* [Exercise: rot13Reader](https://tour.golang.org/methods/23)

Note that you can change the code inline in the browser and run the code to see the results.

### Task 2: Go Language Questions

Answer these multiple choice questions about [Go programming](go_questions.md).

### Task 3: Go Exercises

Before you start working on the assignments below, make sure that your local working copy has all the latest changes from the course [assignments](https://github.com/dat320-2020/assignments) repository.
Instructions for fetching the latest changes are [here](https://github.com/dat320-2020/course-info/blob/master/lab-submission.md#update-local-working-copy-from-course-assignments).

1. In the following, we will use `sequence/triangular.go` exercise as an example.
   The file contains the following skeleton code and task description:

    ```golang
    package sequence

    // Task 1: Triangular numbers
    //
    // triangular(n) returns the n-th Triangular number, and is defined by the
    // recurrence relation F_n = n + F_n-1 where F_0=0 and F_1=1
    //
    // Visualization of numbers:
    // n = 1:    n = 2:     n = 3:      n = 4:    etc...
    //   o         o          o           o
    //            o o        o o         o o
    //                      o o o       o o o
    //                                 o o o o
    func triangular(n uint) uint {
        return 0
    }
    ```

2. Implement the function body according to the specification so that all the tests in `sequence/triangular_test.go` passes.
   The test file looks like this:

    ```golang
    package sequence

    import (
        "testing"

        "github.com/google/go-cmp/cmp"
    )

    var triangularTests = []struct {
        in, want uint
    }{
        {0, 0},
        {1, 1},
        {2, 3},
        {3, 6},
        {4, 10},
        {5, 15},
        {6, 21},
        {7, 28},
        {8, 36},
        {9, 45},
        {10, 55},
        {20, 210},
    }

    func TestTriangular(t *testing.T) {
        for _, test := range triangularTests {
            if diff := cmp.Diff(test.want, triangular(test.in)); diff != "" {
                t.Errorf("triangular(%d): (-want +got):\n%s", test.in, diff)
            }
        }
    }
    ```

3. There are several ways to run the tests. If you run:

   ```console
   go test
   ```

   the Go tool will run all tests found in files whose file name ends with `_test.go` (in the current directory).
   Similarly, you can also run a specific test as follows:

   ```console
   go test -run TestTriangular
   ```

4. You should ***not*** edit files or code that are marked with a `// DO NOT EDIT` comment.
   Please make separate `filename_test.go` files if you wish to write and run your own tests.

5. When you have completed a task and sufficiently many local tests pass, you may push your code to GitHub.
   This will trigger QuickFeed which will then run a separate test suite on your code.

   Using `sequence/triangular.go` as an example, use the following procedure to commit and push your changes to GitHub and QuickFeed:

    ```console
    $ git add triangular.go
    $ git commit
    // This will open an editor for you to write a commit message
    // Use for example "Implemented Assignment 2"
    $ git push
    ```

6. QuickFeed will now build and run a test suite on the code you submitted.
   You can check the output by going to the [QuickFeed web interface](https://uis.itest.run).
   The results (build log) is available from the Labs menu.
   Note that the results shows output for all the tests in current lab assignment.
   You will want to focus on the output for the specific test results related to the task you're working on.

7. Follow the same process for the other tasks included in this lab assignment.
   Each task contains a single `.go` template file, along with a task description and a `_test.go` file with tests.

8. Finally, complete the task in `cmd/terminal/main.go`.

   *Note: You must print the current working directory at the start of every prompt.*
   For example, if the working directory is `/home/username`, your prompt should be something like the following:

   ```console
   /home/username>
   ```

   Then we can input commands after the prompt, e.g. the `ls` command:

   ```console
   /home/username> ls
   ```

   You can use the [`os.Getwd` function](https://golang.org/pkg/os/#Getwd) to find the current working directory.
   This could be implemented in a function named `printPrompt()`.

9. When you are done with all assignments and want to submit the final version, please follow these [instructions](https://github.com/dat320/course-info/blob/master/lab-submission.md#final-submission-of-labx).
