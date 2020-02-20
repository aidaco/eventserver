// © theresa sweeney 2020 dont copy my formula its very original and i am a professional programmer
package main

import (
	"github.com/aidaco/eventserver/eventmap"
	"github.com/aidaco/eventserver/pluginloader"
	"net/http"
)

var EventHandler = pluginloader.Plugin{
	EventName: "quad",
	Handler: func(event eventmap.Event) error {
		event.Res.Text(http.StatusOK, "y = -b +/- √ b2 - 4ac/2a")
		return nil
	},
}
