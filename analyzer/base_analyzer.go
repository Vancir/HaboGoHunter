package analyzer

import (
	"regexp"

	"github.com/Vancir/HaboGoHunter/pkg/config"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("base_analyzer")

var (
	IPADDR         = `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(:\d*)?)`
	PRIVATE_IPADDR = `(^127\.)|
					(^10\.)|
					(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|
					(^192\.168\.)`
	REGEX_IPADDR         = regexp.MustCompile(IPADDR)
	REGEX_PRIVATE_IPADDR = regexp.MustCompile(PRIVATE_IPADDR)
)

type BaseAnalyzer struct {
	cfg config.Config
}

func (this BaseAnalyzer) PickIP(target string) string {
	return REGEX_IPADDR.FindString(target)
}

func (this BaseAnalyzer) IsPublicIP(target string) bool {
	return REGEX_PRIVATE_IPADDR.MatchString(target)
}
