package main

import (
    "Booksgen/cmd/api"
    st "Booksgen/cmd/standalone"
    u "Booksgen/internal/utils"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("%s\n", HELPMSG)
        return
    }

    if u.HasArg("-h") || u.HasArg("--help") {
        fmt.Printf("%s\n%s\n%s\n",
            HELPMSG, STANDHELPMSG, APIHELPMSG)
    }

    if u.HasArg("-st") || u.HasArg("--standalone") {
        if len(os.Args) < 3 {
            fmt.Printf("%s", STANDHELPMSG)
        } else {
            st.Run()
        }

    } else if u.HasArg("--api") {
        api.Run()
    }
}
