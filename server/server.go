package main

import
(
    "fmt"
    "github.com/vohumana/vohumana-gonetwork/NetworkStack"
    "log"
)

func OnServerData(c *NetworkStack.ServerClient, message []byte) {
    if (len(message) == 0) {
        return
    }

    // fmt.Println("OnData")
    // fmt.Printf("%v\n", c.Connection.LocalAddr())
    // fmt.Println(string(message))
}

func OnServerClose(c *NetworkStack.ServerClient, err error) {
    fmt.Println("Connection closed")
    fmt.Printf("%v", c.Connection.LocalAddr())
    fmt.Println(err)
}

func OnServerConnected(c *NetworkStack.ServerClient) {
    c.Connection.Write([]byte("Hello Client!"))
}

func SendToAllClients(server *NetworkStack.Server, data []byte) {
    for i := 0; i < len(server.ServerClients); i++ {
        err := server.ServerClients[i].Send(data)
        if (err != nil) {
            log.Println(err)
        }
    }
}

func main() {
    server := NetworkStack.NewServer(":9999", OnServerData, OnServerClose, OnServerConnected)
    go server.OpenAndListenForConnections()
    fmt.Println("Server is listening on port 9999")

    for {
        var command string
        var isExit bool
        fmt.Printf("Enter a commannd: ")
        fmt.Scanln(&command)

        switch (command) {
            case "exit":
                isExit = true
            break

            case "clients":
                fmt.Println("Connected clients")
                for i := 0; i < len(server.ServerClients); i++ {
                    fmt.Printf("%v\n", server.ServerClients[i].Connection.LocalAddr())
                }
            break

            case "init":
                SendToAllClients(server, []byte("Initalize"))
            break
        }

        if (isExit) {
            break
        }
    }
}