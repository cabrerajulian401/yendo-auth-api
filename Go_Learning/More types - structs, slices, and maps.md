#### Exact Source: Go Tour (official Website): https://go.dev/tour/welcome/1


# Pointers 

Go has pointers. 

A pointer holds the memory address of a value, pretty much acting as a separate reference to it holding its underlying value(memory address)


The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

`var p *int` : In this `*int` is a pointer to the `p` value

The `&` operator generates a pointer to its operand

Example: ```
```
i := 42
p = &i
```

In this `&` generates the pointer for `i` being `p`. 

### Use use `&` to point


The `*` operator denotes the pointer's underlying value.

Example:
```
fmt.Println(*p) // read i through the pointer p
*p = 21 // set i through the pointer p


```


Analyzing examples of pointers:
```
i, j := 42, 2701


p := &i  // point to the i
fmt.Println(*p) // read i through the pointer [prints 42]
*p = 21 // set i through the pointer
fmt.Println(*p) // seeing new value which is now 21

p = &j // points to j 
*p = *p / 37 // divide j through the pointer
fmt.Println(j) // see the new value of j
```

This entire pointer process is known as "dereferencing" or "indirecting". 

Unlike C, Go has not pointer arithmetic. 



## Structs

A `struct` is a collection of fields which are basically parameters that have both a name and a data type. Serving as pretty much a template in essence.

Think of structs as object having a **NAME** and a **TYPE**


Struct literal syntax: `type Struct_name struct`

Example: 
![[Screenshot 2025-07-19 at 11.29.00 PM.png]]


- You can see that the Struct holds the Fields or Arguments with their types also declared next to them
- Note that this struct in this example return in braces: `{1 2}


# Struct Fields 

Struct Fields are accessed using a dot

See below for example:

```
type vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

}
```

You can see in the example above the `struct` is assigned to the variable `v` and therefore it can use the dot `.` to access the `X` **field** (in the Struct) to assign it the value of 4

Side-Note: When accessing the field in the struct for value assignment we can just use the `=` and not `:=` since we do not need the `type`


# Pointers to structs

Structs can be accessed through a struct pointer.

To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`  (remember that `*` denotes pointers underlying value). However that is unneeded, we can just write `p.X` where `p` is the pointer and it is accessing the `X` field.

Example:
```
type vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9

	fmt.Println(v)

}
```


# Struct Literals


A **Struct literal** denotes a newly allocated struct value by listing the values of its fields. Basically is the instance: `Vertex{...}` - looks like a call

You can only declare a Subset of the fields using the following inside a struct literal: `NAME : value`

If you are using named fields, their order of placement is irrelevant.

The Special pointer prefix `&` returns a pointer to the Struct Value.

Example: 

```
v1 = Vertex{1, 2} // has type Vertex
v2 = Vertex{X: 1} // subset for X field
v3 = Vertex{}.    // both X:0 and Y:0 since it is not init
p = &Vertex{1, 2} // has type *Vertex



```

- Note that `p` will return as `&{1 2}` when printed out for the pointer operand 



## Arrays 

The type `[n]T` is an array of n value of `n` value of type `T`

The expression:
`var a [10]int` 

Note that `a[10]int` is an array literal

Declares a variable `a` as an array of ten integers. so in this case `[10]int` is a type in this.

**Therefore an array's length is part of its type, so arrays cannot be resized**. This seems limiting but we will find out later that we can work with this.

Example below:
```
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

Since `var a` is declared you can slice `a` to assign the value of each of the arrays. 

Also notice that you can specify which values you want to initialize in the array: `[6]int{2, 3, 5, 7, 11, 13}` - *This is also called an array literal*

# Slices


An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays. 

The type `[]T` is a slice with elements of type `T`. So basically `[]int` is the slice in the example below with `int` being the base type.


A Slice is formed by specifying two indices, a low and high bound, separated by a colon:

`a[low : high]`

**This selects a half-open rang which includes the first element, but excludes the last one.**


# Slices are like references to arrays


A slice does not store any data, it just describes a section of the underlying array.

Changing the elements of a slice modifies the corresponding elements of the underlying array. Note that in this the Slice is essentially acting as a a variables then modifying that new slice variables in another slice (changing it) modifies the underlying array.

Other slices that share the same underlying array will also see those changes.

You can pretty much leverage slices to modify arrays.


Example Code & Terminal Output:

![[Screenshot 2025-07-20 at 1.04.41 AM.png]]



# Slice Literals

*Side-Note: In Go, a literal is a notation in source code that directly represents a fixed value. It's a way to express data values without requiring any computation or evaluation.*


**A slice literal is like an array literal without the length. **
- This would essentially be the difference between a slice and an array, that an array literal would specify the length but the slice literal would not


### The Core Concept: A Slice _of_ a Type

**In Go, a slice is always a collection _of_ a single, specific type**. The syntax is always `[]TypeName`.

- `[]int`: A slice where every element is an `int`.

- `[]string`: A slice where every element is a `string`.

- `[]MyStruct`: A slice where every element is a `MyStruct`.

This clarification comes from the example below:
`var authorization []MyStruct`


You are simply telling Go, "I'm creating a list, and everything in this list will be a `MyStruct`."


This is an array literal:

`[3]bool{true, true, false}`


And this slice literal below creates the same array as above, **then builds a slice that references it *(just does not have the length):

`[]bool{true, true, false}`


In this slace the slice tahat would reference the two arrats below would be `r` and `q`


![[Screenshot 2025-07-20 at 1.09.33 AM.png]]


# Slice Defaults

When slicing you may omit the the high or low bounds to use defaults instead (same as python).

The default is index `0` for the low bound and the length of the slice for the high bound.

For the array:
`var a [10]int`

- Remember an array declaration is in the structure: `var name [len]type`

These slices of the array above are therefore also equivalent to the one above using the default:
`a[0:10]`
`a[:10]`
`a[0:]`
`a[:] // grabs total array using just ':'`


# Slice length and capacity

A slice has both a *length* & *capacity*

The length of a slice is the number of elements it contains. You can find the length of a slice using `len(s)` assuming that `s` is a slice. 


The capacity of a slice is the number of elements in the underlying array counting from the first element in the slice. You can find the capacity of a slice using `cap(s)` assuming that `s` is a slice.

You can extend a slices length by re-slicing it, just make sure that it has the adequate capacity. If you extend out of the slice range, you would get the following error `panic: runtime error: slice bounds out of range [:9] with capacity 6`

![[Screenshot 2025-07-20 at 2.36.11 PM.png]]


Note that you slice just by reassignment, so you can do `s = s[:0]` or `s = s[:9]

# Nil Slices


The zero value of a slice is `nil`

A nil slice has a length and capacity of 0 and has not underlying value.

![[Screenshot 2025-07-20 at 2.40.29 PM.png]]

- Note that the `s` slice literal is not assigned any elements so it's value is `nil` and it has no underlying array.


# Creating a slice with make

Slices can be created using the built-in `make` function; this is a how you create dynamically-seized arrays. 

- *Side-note*:  Note: that you can use `:=` inside a function, so you do not have to use `var` declaration. This is the main difference between `=` and `:=`, you do not have to use `var` before `:=`

The `make` function initializes an array and returns a slice that refers to to that array in the same statement:

`a := make([]int, 5) // len(a) = 5`

Note that the syntax structure for the `make` function is `make(array_type, length, capacity)

To specify a capacity according to the structure above we can add the third argument being capacity:

`b := make([]int, 0, 5)`

*Side-note:* When re-slicing you do can just use `=`  and not `:=`

`b = b[:cap(b)] // len(b)=5, cap(b)=5`

`b = b[1:] // len(b) = 4, cap(b)=4`

Slices are actually how we can make arrays dynamic and changing since we cannot manipulate the underlying itself directly after it is changed. You can also use the `append` function for slices; see in after 1 section below.

# Slices of slices

Slices can contain any type, including other slices.

Note that slices are a `type`.

See example below:

![[Screenshot 2025-07-20 at 2.57.51 PM.png]]

So the first slice of `board` is pretty much the first slice`[](1)` and each one of the first order slices has a second order slice `[](2)` so that there are 9 total slice elements in total. This creates a 3x3 tic tac toe board.

# Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in `append` function. 

`func appends(s []T vs... T)[]T`
- note that `T` stands for Type

The first parameter of `append` is *a slice of type `T`*  (the slice you want to append to) and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slices plus the provided values.


If the backing(underlying capacity) of the array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will now point to the newly allocated array (Remember pointing is underlying). 

Appending also works on `nil` slices

`var s []int` <- this is initializing a slice literal, there are no elements in braces to the right of it so it is a `nil` slice

`s = append(s,0)`


Example below:

![[Screenshot 2025-07-20 at 3.07.57 PM.png]]


# Range 

The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. the First is the index, and the second is a copy of that element at that index.

Example Below:

![[Screenshot 2025-07-20 at 3.53.48 PM.png]]
You can see that the `i` and `v` are returning the index and actual element. 

The certain snippet is `for i, v := range pow {...}
- `i` and `v` are local variables
	- Both are set equal to `range pow` being `range <slice>`

What is returned:

![[Screenshot 2025-07-20 at 3.55.12 PM.png]]


# Range Continued

You can skip the index or value by assigning to `_`

For example:
`for i, _ := range pow` <- skips over element value
`for _, value: = range pow` <- skips index value

If you only want the index, you can omit the second value for a `range` function:
`for i := range pow`

^ That would only return the index variable, which you can still manipulate.




# Maps

A map maps(places) keys to value
- *note that this is the case with python too with dictionaries*
- Note that there is a `map` function

The zero value of a map is `nil` . A `nil` map has no keys, nor van keys be added.

The `make` function returns a map of the given type, initialized and ready for use.

Example below:

![[Screenshot 2025-07-20 at 4.04.20 PM.png]]


# Map literals

Map literals are like struct literals, but unlike structs the keys are required in the body.

You need to specify the type of the map right after the literal

ex:

`map[type][var/struct]`

Example: 

```
var m = map[string]Vertex{

key: value,

key: value

}





```



![[Screenshot 2025-07-20 at 4.06.17 PM.png]]

Notice in here that you can use a string on a map. to use those field to then insert keys and values into the actual map body. 

# Map Literals continued

If the top-level type is just a type is just a type name, you can omit it from the elements of the literal.


# Mutating Maps

Insert or update an element in map `m`:

`m[key] = element`




Retrieve a specific element from the map: 

`element = m[key]`

Delete an Element using `delete()` function:

`delete(m, key)`

Test that a key is present with a two-value assignment

`element, ok = m[key]`

- If `key` is in `m` ,`ok` is `true` . If not, `ok` is `false`
- if `key` is not in the map, then `element` is zero value for the map's element type


Note: if `elem` or `ok` have not yet been declared you could use a short declaration form: 

`elem, ok := m[key]
`
![[Screenshot 2025-07-20 at 4.22.02 PM.png]]




# Function Values 

**Functions are values tool. They can be passed around just like other values, such as in the parameters of other functions.**

**Function values may be used as function arguments and return values.**

### Example Below:

![[Screenshot 2025-07-20 at 4.24.24 PM.png]]


# Function Closures

Go functions may be closures. A closure is a function value that references variables from outside its body. **The function may access and assign to the referenced variables; in this sense the function is "bound"or "controlled to/by the variable**


For example, the `adder` function returns a closure which is the `func(x int) int)` . Each closure is bound to its own `sum` variable, in this case the `func(x int)` closure is bound to the `sum` variable that is defined in the `adder()` function.

See below:

![[Screenshot 2025-07-20 at 4.31.52 PM.png]]














