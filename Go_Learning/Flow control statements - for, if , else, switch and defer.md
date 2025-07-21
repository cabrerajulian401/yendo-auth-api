#### Exact Source: Go Tour (official Website): https://go.dev/tour/welcome/1


# For

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components seperated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`. 

Example:

```
func main() {
	sum:=0
	for i:=0; i < 10; i++{ 
	
		sum += 1
	
	}

	fmt.Println(sum)


}
```

Note: that you can use `:=` inside a function, so you do not have to use `var` declaration.

Other Note: Unlike C, Java, or JavaScript there are no parenthesis surrounding the the 3 components(initializer, condition, & post) for the `for` loop and the `{..}` braces  are always required


# For Continued

Know that the `init` and `post` statements in for loop are actually optional and not required argument/components. 

Example Below:
![[Screenshot 2025-07-19 at 8.47.29 PM.png]]
Note: The Semi Colons between the actual condition is still required.


## `for` can also be a `while` statement in Go

![[Screenshot 2025-07-19 at 8.49.49 PM.png]]

- You can see above that you drop the semi colons. 
- Note that the statement above also functions as a while statement
	- The same is being iterated as long as the Sum is less than 1000 just by using one argument/component in the for loop

# Forever 

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

![[Screenshot 2025-07-19 at 8.52.29 PM.png]]

# If with a short statement 

Like `for` , the `if` statement can start with components before getting to the actual body of the statement. These are also separated by semi-colons.
![[Screenshot 2025-07-19 at 8.54.22 PM.png]]

- Note that variables declared in this statement are only in scope until the end of the `if` body (very local)

-  Note that the `if` statement has two parts: 
	 - The initialization component
	 - The condition component

## If and Else

Variables declared inside an `if` short statement are also local to any `else` blocks.

If you look below variable `v` if local to both the if and else function which was declared in the short if statement.![[Screenshot 2025-07-19 at 8.57.48 PM.png]]



## Switch

A `switch` statement is a shorter way to write a sequence of `if-else` statements. It runs the first `case` whose value is equal to the condition expression. This is also the same as the `switch` statements in Java.

It is similar to Java, but Go only actaully runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automaticallyin GO since it exits the `switch` statements when it hits the first case.


Example Below:

![[Screenshot 2025-07-19 at 9.02.38 PM.png]]

Note that you can initialize variables that will be uses as conditions in your cases below. 

If none of the cases are used you can have the `defaut` case at the bottom.


## Switch Evaluation Order

Switch cases evaluate from top to bottom, stopping when a case succeeds.
![[Screenshot 2025-07-19 at 9.16.48 PM.png]]

It does not call `f` if `i==0`

# Switch with no condition

Switch without a condition is the same as `switch true`

This construct can be a clean way to write long if-then-else chains.


![[Screenshot 2025-07-19 at 9.18.39 PM.png]]


# Defer

A defer statement delays the execution of the statement/function until the surrounding function returns.

The deferred call's argument are evaluated immediately, but the actual function call is not executed until the surrounding function returns.

Example Below:

![[Screenshot 2025-07-19 at 9.21.02 PM.png]]

Notice that since the `defer` keyword is next to the statement `"world"` does not actually print out before `"hello"` is executed on the surrounding function.


# Stacking defers

Deferred function calls are pushed onto a stack, like recursion.

Deferred calls are executed last in first out order (up and down).

Example: 
![[Screenshot 2025-07-19 at 9.24.01 PM.png]]


Note that when a ran this, the `for` loop executed printing out starting from `9` and going down to `0`. This is because of the Deferred stack.






