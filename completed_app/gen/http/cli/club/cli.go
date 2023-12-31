// Code generated by goa v3.13.2, DO NOT EDIT.
//
// club HTTP client CLI support package
//
// Command:
// $ goa gen app/design

package cli

import (
	bandc "app/gen/http/band/client"
	orderc "app/gen/http/order/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `band play
order tea
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` band play --body '{
      "style": "Bebop"
   }'` + "\n" +
		os.Args[0] + ` order tea --body '{
      "includeMilk": false,
      "isGreen": false,
      "numberSugars": 375920453996173358
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		bandFlags = flag.NewFlagSet("band", flag.ContinueOnError)

		bandPlayFlags    = flag.NewFlagSet("play", flag.ExitOnError)
		bandPlayBodyFlag = bandPlayFlags.String("body", "REQUIRED", "")

		orderFlags = flag.NewFlagSet("order", flag.ContinueOnError)

		orderTeaFlags    = flag.NewFlagSet("tea", flag.ExitOnError)
		orderTeaBodyFlag = orderTeaFlags.String("body", "REQUIRED", "")
	)
	bandFlags.Usage = bandUsage
	bandPlayFlags.Usage = bandPlayUsage

	orderFlags.Usage = orderUsage
	orderTeaFlags.Usage = orderTeaUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "band":
			svcf = bandFlags
		case "order":
			svcf = orderFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "band":
			switch epn {
			case "play":
				epf = bandPlayFlags

			}

		case "order":
			switch epn {
			case "tea":
				epf = orderTeaFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "band":
			c := bandc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "play":
				endpoint = c.Play()
				data, err = bandc.BuildPlayPayload(*bandPlayBodyFlag)
			}
		case "order":
			c := orderc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "tea":
				endpoint = c.Tea()
				data, err = orderc.BuildTeaPayload(*orderTeaBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// bandUsage displays the usage of the band command and its subcommands.
func bandUsage() {
	fmt.Fprintf(os.Stderr, `A band that plays jazz.
Usage:
    %[1]s [globalflags] band COMMAND [flags]

COMMAND:
    play: Choose your jazz style.

Additional help:
    %[1]s band COMMAND --help
`, os.Args[0])
}
func bandPlayUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] band play -body JSON

Choose your jazz style.
    -body JSON: 

Example:
    %[1]s band play --body '{
      "style": "Bebop"
   }'
`, os.Args[0])
}

// orderUsage displays the usage of the order command and its subcommands.
func orderUsage() {
	fmt.Fprintf(os.Stderr, `A waiter that brings drinks.
Usage:
    %[1]s [globalflags] order COMMAND [flags]

COMMAND:
    tea: Order a cup of tea.

Additional help:
    %[1]s order COMMAND --help
`, os.Args[0])
}
func orderTeaUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] order tea -body JSON

Order a cup of tea.
    -body JSON: 

Example:
    %[1]s order tea --body '{
      "includeMilk": false,
      "isGreen": false,
      "numberSugars": 375920453996173358
   }'
`, os.Args[0])
}
