package main

import (
	"github.com/code-serenade/conet-go/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":17111")
}
