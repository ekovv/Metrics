package agent

import (
	"flag"
	"os"
	"strconv"
)

//var (
//	FlagRunAddr           string
//	FlagIntReportInterval int
//	FlagIntPollInterval   int
//)
//
//func ParseFlagsAgent() {
//	flag.StringVar(&FlagRunAddr, "a", "localhost:8080", "address and port to run server")
//	flag.IntVar(&FlagIntReportInterval, "r", 10, "interval of send metrics")
//	flag.IntVar(&FlagIntPollInterval, "p", 2, "interval of update metrics")
//	flag.Parse()
//	envRunAddr := os.Getenv("ADDRESS")
//	envReport := os.Getenv("REPORT_INTERVAL")
//	envPoll := os.Getenv("POLL_INTERVAL")
//
//	if envRunAddr != "" {
//		FlagRunAddr = envRunAddr
//	}
//	if reportInterval, _ := strconv.Atoi(envReport); reportInterval != 0 {
//		FlagIntReportInterval = reportInterval
//	}
//
//	if pollInterval, _ := strconv.Atoi(envPoll); pollInterval != 0 {
//		FlagIntPollInterval = pollInterval
//	}
//}

type Config struct {
	Host           string
	ReportInterval int
	PollInterval   int
}

type F struct {
	host           *string
	reportInterval *int
	pollInterval   *int
}

var f F

func init() {
	f.host = flag.String("a", "localhost:8080", "address")
	f.reportInterval = flag.Int("r", 10, "interval of send metrics")
	f.pollInterval = flag.Int("p", 2, "interval of update metrics")
}

func New() Config {
	flag.Parse()

	envRunAddr := os.Getenv("ADDRESS")
	envReport := os.Getenv("REPORT_INTERVAL")
	envPoll := os.Getenv("POLL_INTERVAL")

	if envRunAddr != "" {
		f.host = &envRunAddr
	}
	if reportInterval, _ := strconv.Atoi(envReport); reportInterval != 0 {
		f.reportInterval = &reportInterval
	}

	if pollInterval, _ := strconv.Atoi(envPoll); pollInterval != 0 {
		f.pollInterval = &pollInterval
	}

	return Config{
		Host:           *f.host,
		ReportInterval: *f.reportInterval,
		PollInterval:   *f.pollInterval,
	}
}
