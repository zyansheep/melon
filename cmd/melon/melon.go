package main

import (
    "fmt"
    "github.com/melondevs/melon/internal/util"
)

func main() {
    config := util.LoadConfig()
    fmt.Println(config.Address)
}
