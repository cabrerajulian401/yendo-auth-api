
#### Exact Source: Go Tour (official Website): https://go.dev/tour/welcome/1


# Notes  

# A Tour of Go

# Packages, variables, and functions

# Packages


Every Go program is made up of packages 

Programs start running in package `main` . Which is the certain package for the main file

Example:

```
package main

import (
	"fmt"
	"math/rand"

)




```

This program is using the packages with import paths `"fmt"` & `"math/rand"

The Package name is the same as the last element of the import path. For Example the `math/rand`package has file that begin with the statement `package rand`
- Above `"fmt"` and `"math/rand"` packages are imported, making them import 

**You Import Packages**


## Imports 

Code used above and below groups import into a "factored " import statement.

Note that in Go if you are going to import a package you have to use it.

Also not that you use parenthesis to write a package block.

```
package main

import (
	"fmt"
	"math/rand"

)
```


You can also write multiple import statement like:
`import "fmt"`
`import "math"`

However it is a good style to use the factored import statement

## Exported Names 

In Go, a name(variable) is exported if it begins with a capital letter. For example `Pi` is an exported name from the `math` package.

`pi` does not start with a capital letter so it is not exported.


When importing a package, you can only refer to its exported names. Any "unexpected" names are not accessible from outside the package.

Note that you changed: `math.pi` to `math.Pi` to therefore change it to a name you can export it out of the`math` package.

## Functions

A function can take zero or more arguments.

**Note that type of the variable comes after the variable name**

for example `x int` or `y int` 

### Example of a function:
![[Screenshot 2025-07-19 at 2.12.10 PM.png]]

- Note that you have to declare the type of the function after you declare its names and pass in the argument. Example is ``func add(x int, y int) int {...}


## Function Continued 

When two or more consecutive function parameters share a type, you cam omit the type from all but the last one.

For Example you can shorten:

`x int, y int`

to:

`x,y int`


## Multiple Results

A function can return any number of results. 

The `swap` function returns 2 strings. Therefore returning multiple results. 

See Example Below:

![[Screenshot 2025-07-19 at 2.17.09 PM.png]]



## Named return values 

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These Names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statement should only be used in short function as example show here. They can harm readability in longer functions.

Example of a naked return statement is below:![[Screenshot 2025-07-19 at 2.23.51 PM.png]]


## Variables 

The `var` statement declares a list of variables; as in function argument lists, the type is last.

Structure:

`var var_name = reference


A `var` statement can be at package or function level. We See both in this example.
![[Screenshot 2025-07-19 at 2.25.16 PM.png]]

- In this example `var` is being used to declare a list of variables being `i` (printed 0, 0 indexed), `c`, `python`, `java`
	- However they are not initialized

*Side-note: `fmt` Package is for input/output logging*


## Variables with initializers 

A var declaration can include initializers, **one per variable

**If an initializer is present, the type can be omitted**; the variable will take the type of the initializer


```
package main

import "fmt"

var i, j int = 1, 2

fmt.Println(....)
```


- So in this know that the two variables are seperated by common and the 2 values they are assigned to are also seperated by commas
- Know that also the type comes after both of the variables - You dont even need type their in this example


Here is another one:
`` var c, python, java = true, false, "no!"

In this:
- `c` is taking the value of a` true`(bool type)
- `python` is taking the value of `false` (bool type)
- `java` is taking the value of `"no!"` string type


## Short Variable Declarations

Inside a function, the short`:=` assignment statement can be used in place of a `var` declaration with implicit type( referencing the type of assigner). 

Remember that use use the normal `=` operator for normal `var` assignments.

Outside a function, every statement begins with a keyword (`var, func` and so on) and so the `:=` construct is not available. 

``` package main
	import "fmt"
 
	func main(){ # Func declaration

	# Assignments below
	var i, j int = 1, 2 
	k := 3 
	c, python, java := true, false, "no"

	fmt.Println(i, j, k, c, python, java)
		
	}
```


## Basic Types 


Go's Basic types are the following:

bool, string, int, byte, rune, float32/float64, complex 64 complex 128


Variable Declarations can also be factored into blocks:
``` 
var ( 

bagel, donut int = 1, 2

python, julian string = "sand","paper





)




```

- Remember that blocks are made with Parenthesis

## Zero Values 

Variables **without an explicit initial value** are given their zero value. 
- *Meaning they are not assigned*

Zero Value is :
- `0` for numeric types
- `false` for boolean type
- `""` (the empty string) for string (could be good for initializes empty string values )

**Example of non-assigned string values below:**
![[Screenshot 2025-07-19 at 2.53.29 PM.png]]


## Type Conversions

The expression `T(v)` converts the **value** `v` to the **type** `T`. 

Some numeric conversions are: 
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply (assuming inside function):

```
i := 42
f := float 64(i)
u := uint(f)
```

Unlike in C, Go assignment between items of different type requires an explicit conversion. Try removing the `float64` or `uint` conversions in the example and see what happens.

Therefore you need to specify what you are converting to, using the `Type` for specification in the `Type(Value` formula.

## Type Inference

When declaring a variable without using explicit type on the left side. The type of the Assignee is inferred by the value on the right hand side.

For example: 

`var i int`
`j := i // j is an int`

*Side-note: `//` are used as comments*

However when the right hand side contains an untyped numeric constant (number is not a declared variable, and is a numeric), the new variable may be an `int`, `float64`, or `complex128`  depending on the precision of the constant.

Example:

```
i := 42           // int
f := 3.142        // float64
g := .0867 + 0.5i // complex128
```


## Constants

Constants are declared like variables but with the `const` keyword.

Constants an be character, string, boolean, or numeric values.

Constants cannot be declared using the `:=` syntax, only `=`

## Structure

`const constant_name = reference


### Example: 

```
package main

import "fmt"

const Pi = 3.14 # Example of Constant




```

- Note that when I trued myself to assign a seperate var of `Truth` which was already a `const` declared earlier. **The Go Build failed to compile.**


## Numeric Constants 

Numeric Constants are high-precision values

An untyped constant takes the type needed by its context

