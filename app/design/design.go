package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("club", func() {
	Title("The Club")
	Description("A club that serves tea and plays jazz. A Goa and Speakeasy example.")
	Version("1.0.0")
    Server("club", func() {
        Host("localhost", func() {
            URI("http://localhost:51000")
            URI("grpc://localhost:52000")
        })
    })
})

var _ = Service("order", func() {
	Description("A waiter that brings drinks.")
	Method("tea", func() {
		Description("Order a cup of tea.")
		Payload(func() {
			Field(1, "isGreen", Boolean, "Whether to have green tea instead of normal.")
			Field(2, "numberSugars", Int, "Number of spoons of sugar.")
			Field(3, "includeMilk", Boolean, "Whether to have milk.")
		})
		Result(String)
		HTTP(func() {
			GET("/tea")
		})
		GRPC(func() {
		})
	})
	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("band", func() {
	Description("A band that plays jazz.")
	Method("play", func() {
		Description("Choose your jazz style.")
		Payload(func() {
			Attribute("style", String, "Style of music to play", func() {
				Enum("Bebop", "Swing")
				Meta("rpc:tag", "1")
			})
			Required("style")
		})
		Result(Empty)
		HTTP(func() {
			POST("/play")
		})
		GRPC(func() {
		})
	})
	Files("/openapi.json", "./gen/http/openapi.json")
})

