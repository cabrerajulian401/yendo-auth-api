#### Exact Source: Go by Example: HTTP Servers https://gobyexample.com/http-servers





To write a basic HTTP server you need to use the net/http package.

`"net/http"`

A fundamental concept in `net/http` servers is *handlers* . A handler is an object implementing the `http.Handler` interface. A common way to write a handler is by using the `http.Handlerfunc` adapter on functions with the appropriate signature. Most of the time it has two arguement being `w http.RequestWriter` & `req &http.Request*`

Example:

`func hello(w http.ResonseWriter, req*http.Request)`

Functions serving as handlers take a `http.ResponseWriter ` and a `http.Request` as arguments . The response writer is used to fill in the HTTP response, note the `w` . 

In ours our simple response, our simple response if just `"hello\n"` 


-------------------------------------------------------------

The handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body. 
![[Screenshot 2025-07-20 at 6.45.14 PM.png]]




We then register our handler functions using the `http.HandleFunc` convenience function. It sets up the defaults router in the `net/http` package and takes a function as an argument, note we put `"/hello"` and `"/headers"` blow for the arguments being passed in the `w http.ResponseWriter`
- *Definition of a router: a device that connects two or more packet switched networks or subnetworks*


**Note that each Handler function call is a route** . In this case `"/hello"` & `"/headers"` are the name of the routes which serve as the first arguement and `hello` and `header` serve as the second arguement which are the actual functions

Finally, we can call the `ListenAndServe` with the *port* and a *handler*.

`nil` tells it to use the default router we have set up below.

See Example below:


![[Screenshot 2025-07-20 at 6.50.05 PM.png]]



## Another Example:

![[Screenshot 2025-07-20 at 11.37.57 PM.png]]

Function returns a value type `err` when it attempts to start a web server. Using the `http.ListenAndServe` function if it starts correctly it returns `nil` . If it fails it returns an error object. 

Next if the `err` is not equal to `nil` then it will log the error using the log command `Fatal` method.


# More on the `http.ListenAndServe` function

The first argument is the server address, and the second is the request handler.

---

### ## Function Signature

The actual function signature looks like this:

Go

```
func ListenAndServe(addr string, handler Handler) error
```

---

### ## Argument 1: Address (`addr string`)

You're correct about the first argument. It's a **string** that specifies the network address for the server to listen on.

- `":8080"` means listen on **port 8080** on all available network interfaces of the machine.
    
- You could also be more specific, like `"127.0.0.1:8080"`, to only listen for requests on the local machine.
    

---

### ## Argument 2: Handler (`handler Handler`)

This argument tells the server **how to handle incoming requests**. It acts as the main router or switchboard.

- In the example `http.ListenAndServe(":8080", nil)`, the second argument is **`nil`**.
    
- When you pass `nil`, Go uses a default router called `DefaultServeMux`. This is the same global router that your `http.HandleFunc("/authorizations", ...)` line registered its route with.
    
- So, passing `nil` tells the server to use the default set of rules you've already defined.
    

---

### ## The Return Value (`error`)

This is where success or failure is reported.

- After being called, `ListenAndServe` **returns** a value of type `error`.
    
- If the server starts and runs without any initial problems, the function blocks and never returns (until the server is shut down). If it shuts down cleanly, it returns `nil`.
    
- If the server **fails to start** (e.g., the port is already in use), the function immediately returns an `error` object. This returned error is what gets captured in the `err` variable in your `if` statement.
