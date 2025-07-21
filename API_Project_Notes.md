# Project Learning Notes


## POST/GET refresher

GET Request

Its purpose is to request or "get" data from a server (requested by client)

The data you are request is usually sent in the URL itself being the route

Making the same GET request should always return the same result and should never change the data on the server


POST Request

To submit new data to the server to create a new resource. It is a write operation that changes the state of the server


The data is sent securely and invisibly in the body of the HTTP request, not the URL. This is better for sensitive information(passwords or credit card numbers)

Making the same POST request multiple times will create new records.


![[Screenshot 2025-07-21 at 12.34.33 AM.png]]


## GET/POST applied to this Project

`GET /authorizations`

This endpoint uses a GET Request.

When a client calls this endpoint, they are asking "Show me the list of authorizations you have stored". They are not trying to manipulate the list, only to retrieve and observe it. This is a classic read-only operation, making GET the perfect choice.



`POST /authorize`

This endpoint uses a POST request.

When a client calls this endpoint, they are sending a new transaction to be approved. They are asking the server to create a new authorization record. This action manipulates the server's state by adding a new item to your `authorization` slice (acts as our memory db for now). Because it involves creating new data and sending sensitive information (like a card number), POST is correct and the best method to use.


## `handleAuthorize(...)` Function

![[Screenshot 2025-07-21 at 1.14.41 AM.png]]


- `w http.ResponseWriter` (RW) is your tool for writing your response back to the client. you use it to set headers and write the response body
- `r *http.Request` ( * R) contains everything about the incoming request from the client, including the method (`POST`), the URL, headers, and the request body, and the request body (the JSON data)

![[Screenshot 2025-07-21 at 1.05.11 AM.png]]

This is a guard clause to protect the actual endpoint. This function is designed to create new data, so it should only accept `POST` requests.

- `r.Method` check the HTTP method used by the client.
- We have a w in the `http.Error` method because this is what we will use a Response Writer
- if anything other than `POST` , it sends an `http.error` function saying `405 Method Method not` allowed which is done by the `"Method not allowed"` string and the `http.StatusNotAllowed` method. 
- It then returns this `Error` object

`var newAuth Authorization Request`

Here we declare a new, empty struct variable named `newAuth` . Its type is `AuthorizationRequest` , the struct we defined earlier. So therefore predefined structs can be types. 

**Think of this as an empty for that we are about to fill in with the data from the Client Request. Basically filling in this new variable using the previous Struct format for formatting and the client for input**


![[Screenshot 2025-07-21 at 1.13.57 AM.png]]

This is a very critical part of the input process

1. `r.body` : This is the raw data sent by the Client
2. `json.NewDecoder(r.body)`: This creates a "decoder" object that knows how to read JSON Data
3. `.Decode(&newAuth)`: This is the action it tells the decoder to read the JSON from the body and try to map it onto the fields of the `newAuth` struct. The `&` is very important, it means we are giving the memory address of `newAuth` so it can modify and *point to it -acts as pointer* so it can modify the struct directly
- All of the above is assigned the variable `err` 

1. `if err := ....; err != nil`: this is the Standard Go error handling. If the client sends bad JSON, the `.Decode` step will fail and product an error. We catch that error and send bad a `400Bad request` using the `http.Error` function that has the arguments `w`,`err.Error()`(object), and `http.StatusBadRequest`(sends 400)!

![[Screenshot 2025-07-21 at 2.31.31 AM.png]]

- Remember that `nil` means nothing relate this to the `http.ListenAndServe` function as well. You want your function to equal  actually  `nil`  meaning it points to `nil`  or thing because if it did not then the variable would then it would point to an error object.

![[Screenshot 2025-07-21 at 1.37.28 AM.png]]

The Client only send the `CardNumber` & `Amount`. The server is responsible for generating the other data such as the `ID` & `timestamp`

- `uuid.New().String()` : We use the `uuid` library to generate a universally. **This ensures no two transactions ever have the same ID**

- `time.Now().Unix()`: We get the current server time and convert it to a Unix timestamp (the number of seconds since Jan 1, 1970). This is a standard way to store messages

- NOTE: we are accessing a Structs fields (structs don't have ) using the `.` notation for accessing in `newAuth.ID` & `newAuth.Timestamp`


![[Screenshot 2025-07-21 at 1.42.12 AM.png]]

Now that our `newAuth` object is complete ( it has the client's data through mapping and the server-generated ID and timestamp), we add it to our `authorizaiton`slice. The `append` function is Go ways to amakes slices dynamic. Our underlying in-memory "database" being authorizations (remember capacity rule) has been now updated.


![[Screenshot 2025-07-21 at 2.18.05 AM.png]]

`w.Header().set(...)`: We set the `Content-Type` header to `application/json` to tell the client that the body of our response will be JSON.
- First argument of `.Set(...)` is the what `attribute` you want to set and the second one is the `type` 

`w.WriteHeader(...)`: We set explicitly set the HTTP Status code to `201` created . This is better thatn the default `200 OK` becuase it tells the client that their request to create a new resource was successful. 

![[Screenshot 2025-07-21 at 2.22.11 AM.png]]
Finally we need to send the response. So we have to return the full object that was created with the `ID` ,`CardNumber` ,`Amount` & `Timestamp`. We use the `Encode(...)` function to encode our `newAuth`object into JSON and write it to the response.





