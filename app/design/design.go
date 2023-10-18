package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("club", func() {
	Title("The Speakeasy Club")
	Version("1.0.0")
	Description("A club that serves drinks and plays jazz. A Goa and Speakeasy example.")
	Contact(func() {
		Name("Speakeasy Support")
		Email("None. Please use Slack.")
		URL("https://speakeasy-dev.slack.com/join/shared_invite/zt-1cwb3flxz-lS5SyZxAsF_3NOq5xc8Cjw")
	})
	Docs(func() {
		Description("The Speakeasy Club documentation")
		URL("https://www.speakeasyapi.dev/docs")
	})
	License(func() {
		Name("Apache 2.0")
		URL("https://www.apache.org/licenses/LICENSE-2.0.html")
	})
	TermsOfService("https://www.speakeasyapi.dev/docs/terms-of-service")
	Server("club", func() {
		Description("club server hosts the band and order services.")
		Services("band", "order")
		Host("dev", func() {
			Description("The development host. Safe to use for testing.")
			URI("http://{machine}:51000") // use the machine variable below
			URI("grpc://{machine}:52000")
			Variable("machine", String, "Machine IP Address", func() {
				Default("localhost")
			})
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

