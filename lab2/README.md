# Lab 2: Introduction to the C Programming Language

| Lab 2: | Introduction to the C Programming Language |
| ---------------------    | --------------------- |
| Subject:                 | DAT320 Operating Systems and Systems Programming |
| Grading:                 | Pass/fail |
| Submission:              | Individually |

## Table of Contents

1. [Introduction](#introduction)
2. [The C Programming Language](#the-c-programming-language)
3. [Installing a C Compiler](#installing-a-c-compiler)
4. [Hello, World in C](#hello,-world-in-c)
5. [Command-Line Arguments](#command-line-arguments)
6. [Pointers in C](#pointers-in-c)
7. [Memory management](#memory-management)
8. [Data Structures in C](#data-structures-in-c)
9. [Dynamic Memory Allocation with `malloc`](#dynamic-memory-allocation-with-malloc)
10. [Header Files and `#include` Statements](#header-files-and-include-statements)
11. [Bitwise Operators](#bitwise-operators)
12. [Tasks](#tasks)
13. [Task 1: Multiple Choice Questions](#task-1-multiple-choice-questions)
14. [Task 2: Implement a Hello World Program in C](#task-2-implement-a-hello-world-program-in-c)
15. [References for Further Reading](#references-for-further-reading)

## Introduction

This lab gives a brief introduction to the C programming language.
While you will not be asked to program in C in this course, C is the language used for most examples in the textbook, and it is also widely used in real operating system kernels and other low-level programs.
Because of this it is important to be able to *read* C code to understand what is going on.
This lab does not offer an exhaustive overview of everything C, however if you have some experience with other programming languages it should offer enough insight to be able to understand the C code used in the course's textbook.
Those that want to learn C programming *for real*, should consult the book The C Programming Language by Kernighan and Ritchie.

## The C Programming Language

### Installing a C Compiler

There are many C compilers out there, but `gcc` (GNU Compiler Collection) and `clang` (C language family frontend for LLVM) are two popular choices, especially on Unix systems.
To check whether or not you already have a C compiler installed, you can run:

```console
gcc -help
```

If this outputs lots of command-line flags, then you're good to go and can skip the rest of this section.

Most Linux distributions come with `gcc` installed; if not you can use the package manager of your Linux distribution to install `gcc`.
On macOS you may need to install the command-line tools package, if you haven't already done so:

```console
xcode-select --install
```

We haven't got specific instructions for Windows here, but it shouldn't be too difficult to find instructions online.

### Hello, World in C

To write a C program that prints a simple message on the console is only slightly more involved than what you would do in Python, or with a shell.

```c
#include <stdio.h>

int main()
{
    printf("Hello, system programmers\n");
}
```

If you create and save this to a file called `hello_programmers.c`, you can compile it with the GNU Compiler Collection, or `gcc` command:

```console
cd <directory that contains hello_programmers.c>
gcc hello_programmers.c
```

This will produce an executable named `a.out` that you can run like this:

```console
./a.out
```

To give the executable a specific name you can instead compile it with the `-o` option:

```console
gcc hello_programmers.c -o hello_programmers
```

And run it like this:

```console
./hello_programmers
```

*Question why do we have `./` before the program's name?*
This is necessary to execute files in the current directory, which on the command line is referred to with the dot (.).
However, this was not necessary when executing `gcc`.
This is because `gcc` and other commands are installed in directories that have been added to the search `PATH`.

### Command-Line Arguments

We can expand on the previous example by adding support for command-line arguments.
Command-line arguments passed to a C program can be accessed through the `argc` and `argv` parameters of the `main` function as shown below:

```c
#include <stdio.h>

// argc and argv are set automatically
int main(int argc, char* argv[])
{
    // print the command-line arguments in order
    for (int i = 0; i < argc; i++) {
        printf("Argument %d: %s \n", i, argv[i]);
    }
    return 0;
}
```

* `argc` is the number of arguments given to the program.
  There will always be at least one argument, i.e. the name of the program itself.

* `argv` is an array of strings containing the program arguments.
  The first element in the array is the name of the program itself, followed by the arguments provided on the command-line.

Assuming the program above is compiled to an executable called `printargs`:

```console
gcc printargs.c -o printargs
```

And executed with the arguments `arg1`, `arg2` and `arg3`:

```console
./printargs arg1 arg2 arg3
```

It will print the following:

```console
Argument 0: ./printargs
Argument 1: arg1
Argument 2: arg2
Argument 3: arg3
```

### Pointers in C

Pointers is a feature of many programming languages and is especially useful when creating data structures, e.g. linked lists and maps.
The following code snippets show the most basic uses.
You can think of a pointer as a variable that instead of holding a value, it holds a pointer (or memory location), where some value or data structure is stored.
A pointer typically has a *type*, which corresponds to the type of the data or data structure that it points to.

Working with pointers in C and Go uses the same syntax, specifically, the `*` and the `&` symbols.
The `*` symbol is used to declare a variable as a pointer and to dereference its value, meaning to access the value pointed to by the pointer.

The `&` symbol allows you to obtain the address of a non-pointer variable.
This is useful to pass the address of your variable to a function, instead of passing the value (e.g. a large data structure) to the function.
Thus, the called function can then operate on the same data structure without copying.

```c
// declare a pointer variable p1 that points to an int
int *p1;

int i = 0;
printf("%d\n", i); // prints 0

// &i gives the address of the variable i
p1 = &i;
// p1 now points to the address of i

// *p1 dereferences the pointer p1:
// -> can now modify/read contents of memory location it points to
*p1 = 1337;
printf("%d\n", i); // prints 1337
```

`printf` works like many other programming languages you might know.
More details about `printf` are available [here](https://www.tutorialspoint.com/c_standard_library/c_function_printf.htm).

### Memory management

There are three types of variables in C:

* **Static variables/global variables**:
  Static variables are stored in the "initialized data segment", which is a portion of the virtual address space of a program.
  These variables are not deallocated after a function ends, and they keep a consistent state across several function calls and can be altered at run time.

    ```c
    void fn() {
        // automatically allocated
        static int static_var = 2;
    } // static_var won't be deallocated, but stay in the data segment
    ```

* **Local variables**:
  Local variables are defined locally within a function.
  They are automatically allocated when the function is called, and automatically deallocated when the function returns.

    ```c
    void fn() {
        // automatically allocated
        int local_var = 1;
    } // local_var is deallocated
    ```

* **Dynamic variables**:
  These variables are explicitly allocated and deallocated by the programmer using `malloc` and `free`.
  `malloc` returns a pointer to a chunk of memory in the heap (or `NULL` if a problem occurs), and `free` takes a pointer and deallocates it.
  The use of `malloc` and `free` is covered in-depth in Chapter 14: Memory API, which we will get to later in the course.

    ```c
    void fn() {
        // this syntax is described in the malloc section below
        int *dynamic_var = (int *) malloc(sizeof(int));
        // ...

        // if not freed a memory leak occurs
        free(dynamic_var);
    }
    ```

### Data Structures in C

Data structures in C are defined like this:

```c
struct my_struct {
    // contents
    int a, b;
};

int main() {
    struct my_struct varName;
    varName.a = 1, varName.b = 2;
}
```

To avoid having to write `struct my_struct varName` the struct could be defined as a type using `typedef`:

```c
struct my_struct {
    int a, b;
};

// defines the type my_struct_t which can be used in the rest of the code
typedef struct my_struct my_struct_t;

int main() {
    my_struct_t varName;
    varName.a = 1, varName.b = 2;
}
```

The above is a bit verbose.
However, it is also possible to declare type of the struct as follows:

```c
// same effect as the previous example
typedef struct my_struct {
    int a, b;
} my_struct_t;

int main() {
    my_struct_t varName;
    varName.a = 1, varName.b = 2;
}
```

### Dynamic Memory Allocation with `malloc`

As explained earlier, `malloc` is used to allocate memory that will not be deallocated automatically.
`malloc` is included in `stdlib.h`.

#### Allocating Memory for an Integer

```c
// 1. create a pointer my_int pointing to an int
// 2. request the allocation of memory equal to the size of an int using malloc (unit: bytes)
// 3. typecast the void pointer returned from malloc to an int pointer using (int *)
int *my_int = (int *) malloc(sizeof(int));
```

#### Allocating Memory for a Data Structure

`malloc` can also be used to allocate memory for a struct.
The syntax is the same as above:

```c
#include <stdlib.h>

typedef struct my_struct {
    int a, b;
} my_struct_t;

int main() {
    // create a pointer my_object pointing to an instance of my_struct_t
    // and allocate enough memory for it
    my_struct_t *my_object = (my_struct_t *) malloc(sizeof(my_struct_t));
} // allocated memory is only freed automatically when the program finishes
```

Since `my_object` is a *pointer* to a `my_struct_t` object, the members of the struct must be accessed differently:

```c
// ...include statements and definitions...

int main() {
    my_struct_t *my_object = (my_struct_t *) malloc(sizeof(my_struct_t));

    // dereference the my_object pointer and set the a  
    // and b values of this instance of the my_struct_t object
    (*my_object).a = 1, (*my_object).b = 2;
}
```

`my_object->a` is a shorthand notation for accessing the members of a struct pointer:

```c
// ...include statements and definitions...

int main() {
    my_struct_t *my_object = (my_struct_t *) malloc(sizeof(my_struct_t));

    // using the shorthand notation (same result as above)
    my_object->a = 1, my_object->b = 2;
}
```

### Header Files and `#include` Statements

#### What are Header Files

Header files are files with the `.h` extension.
They contain things such as function prototypes, system-wide global variables, constants, and macro definitions.
Often they provide different functionality based on compiler options, such as which OS the code should be compiled for.
Header files can be shared between several files. There are two types of header files:

1. Files defined by the compiler (often called standard library)
   * These are usually located in `/usr/include` or `/usr/local/include` on UNIX systems
2. Files written by the programmer

#### The `#include` Statement

Header files are included in the C program using the `#include` statement.
There are two "variants" of the `#include` statement:

1. `#include <header_file.h>` includes header files defined by the compiler.
2. `#include "header_file.h"` includes header files written by the programmer, which are by default located in the same directory as the C program itself.

The included files are scanned in the "preprocessor state", i.e.
before reading the other contents of the C program.
The output of the header files is then included *at the point* of the `#include` statement in the `.c` file.

#### Macros

In C, you can define macros using `#define [macro] [macro_content]`, which is really just a fragment of code that is given a name.
Macros are useful for making parts of the code less confusing.

An example for a macro could be something like:

```c
// simple number macro
#define PI 3.14159
// using a function as a macro
#define CubeVol(n) n*n*n
// using previous macros to define another macro
#define SphereVol(r) (4/3)*PI*Cube(r)
```

### Bitwise Operators

The C language has bitwise operators which can be used for boolean algebra at the binary level.
Here are the available operators:

| Operator | Description                    | Example             |
|:--------:|--------------------------------|---------------------|
|    `&`   | bitwise AND                    | `0xf & 0x6 == 0x6`  |
|   `\|`   | bitwise OR                     | `0x9 \| 0x6 == 0xf` |
|    `^`   | bitwise XOR                    | `0xf ^ 0x6 == 0x9`  |
|   `<<`   | left shift                     | `0x2 << 2 == 0x8`   |
|   `>>`   | right shift                    | `0x8 >> 3 == 0x1`   |
|    `~`   | bitwise complement (inversion) | `~0xa == 0x5`       |

Here is a simple conversion table for those not familiar with hexadecimal and binary.

| Decimal     |   0  |   1  |   2  |   3  |   4  |   5  |   6  |   7  |   8  |   9  |  10  |  11  |  12  |  13  |  14  |  15  |   16  |
|-------------|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:-----:|
| Binary      | 0000 | 0001 | 0010 | 0011 | 0100 | 0101 | 0110 | 0111 | 1000 | 1001 | 1010 | 1011 | 1100 | 1101 | 1110 | 1111 | 10000 |
| Hexadecimal | 0x00 | 0x01 | 0x02 | 0x03 | 0x04 | 0x05 | 0x06 | 0x07 | 0x08 | 0x09 | 0x0A | 0x0B | 0x0C | 0x0D | 0x0E | 0x0F |  0x10 |

## Tasks

### Task 1: Multiple Choice Questions

Answer these multiple choice questions about [C programming](c_questions.md).

### Task 2: Implement a Hello World Program in C

Implement the hello world program in [`hello.c`](./hello/hello.c).
The following function body is given as:

```c
int main(int argc, char* argv[])
{
    // TODO: Implement program
    return 0;
}
```

Use the arguments to get the following functionality:

* If a user runs

  ```console
  ./hello
  ```

  the output should be

  ```console
  Hello, world!
  ```

* If the user supplies an argument, for example Tom

  ```console
  ./hello Tom
  ```

  the output should be

  ```console
  Hello, Tom!
  ```

* If the user supplies multiple arguments,

  ```console
  ./hello Tom Joe
  ```

  the output should be

  ```console
  Hello, Tom!
  Hello, Joe!
  ```

You may compile and run your code with the following commands (replacing `<args>` with some arguments, e.g. `Tom`):

```console
gcc ./hello/hello.c -o ./hello/hello
./hello/hello <args>
```

## References for Further Reading

* [C crash course](https://inst.eecs.berkeley.edu/~cs61c/sp10/tb/C_Crash_Course.pdf)
* [`typedef struct` vs `struct`](https://stackoverflow.com/questions/612328/difference-between-struct-and-typedef-struct-in-c)
* [Header files](https://www.tutorialspoint.com/cprogramming/c_header_files.htm)
* [Preprocessor options/macros](https://gcc.gnu.org/onlinedocs/gcc/Preprocessor-Options.html)
* [Memory Layout of C Program](https://www.geeksforgeeks.org/memory-layout-of-c-program/)
* [Static Variables in C](https://www.geeksforgeeks.org/static-variables-in-c/)
* [C Macros](https://www.programiz.com/c-programming/c-preprocessor-macros#macros)
* [`printf`](https://www.tutorialspoint.com/c_standard_library/c_function_printf.htm)
