// © theresa sweeney 2020 dont copy my formula its very original and i am a professional programmer
package main

import (
	"fmt"
	"github.com/aidaco/eventserver/eventmap"
	"net/http"
)

const EventName string = "quad"

func Handler(event eventmap.Event) error {
	event.Res.Text(http.StatusOK, "y = -b +/- √ b2 - 4ac/2a")
	return nil
}

func main() {
	h := Handler
	e := EventName
	fmt.Println(h, e)
}
