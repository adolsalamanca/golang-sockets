# Golang sockets demo

Small demo that shows how to perform communication through TCP sockets between a server and client in Go.

## Getting Started



### Prerequisites

You just need have go installed in your local machine, this is my version:  *go1.14.4 linux/amd64*


## Running the code

To run this example we need to open two terminals, one for the server and the other one for the client.

First of all, we start the server:

```ssh 
go run server/main.go
```

Then, we start the client, you can change the ports, if you leave the code unchanged you can run like this:

```ssh 
go run client/main.go localhost:1200 100
```

Where localhost:1200 is the address:port of the server and 100 are the number of TCP sockets that will be open in paralell using separated go routines once the program is started.

Now you will be able to see messages coming from one part to the other in both terminals.


## Author

* **Adolfo Rodriguez** - *Golang sockets* - [adolsalamanca](https://github.com/adolsalamanca)

## Docs

-  https://en.wikipedia.org/wiki/Network_socket
-  https://medium.com/swlh/understanding-socket-connections-in-computer-networking-bac304812b5c

## License

This project is licensed under the MIT License


