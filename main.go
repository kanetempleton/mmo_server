// main.go

package main

import (
    "mmo_server/game/cmd"
)

func main() {
    launcher := cmd.NewLauncher()
    //shell := cmd.NewShell(launcher)

    // Your main logic here, for example:
    launcher.Start()
}
