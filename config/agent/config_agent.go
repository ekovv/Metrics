package agent

import (
	"flag"
	"os"
	"strconv"
)

var (
	FlagRunAddr           string
	FlagIntReportInterval int
	FlagIntPollInterval   int
)

func ParseFlagsAgent() {
	flag.StringVar(&FlagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&FlagIntReportInterval, "r", 10, "interval of send metrics")
	flag.IntVar(&FlagIntPollInterval, "p", 2, "interval of update metrics")
	flag.Parse()
	envRunAddr := os.Getenv("ADDRESS")
	envReport := os.Getenv("REPORT_INTERVAL")
	envPoll := os.Getenv("POLL_INTERVAL")

	if envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	if reportInterval, _ := strconv.Atoi(envReport); reportInterval != 0 {
		FlagIntReportInterval = reportInterval
	}

	if pollInterval, _ := strconv.Atoi(envPoll); pollInterval != 0 {
		FlagIntPollInterval = pollInterval
	}
}
