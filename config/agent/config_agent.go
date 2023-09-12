package agent

import "flag"

var (
	FlagRunAddr           string
	FlagIntReportInterval int
	FlagIntPollInterval   int
)

func ParseFlagsAgent() {
	flag.StringVar(&FlagRunAddr, "a", "8080", "address and port to run server")
	flag.IntVar(&FlagIntReportInterval, "r", 10, "interval of send metrics")
	flag.IntVar(&FlagIntPollInterval, "p", 2, "interval of update metrics")
	flag.Parse()
}
