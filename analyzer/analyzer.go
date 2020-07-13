package analyzer

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Vancir/HaboGoHunter/pkg/config"
	"github.com/Vancir/HaboGoHunter/pkg/utils/fileutil"
	"github.com/Vancir/HaboGoHunter/pkg/utils/osutil"
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

var (
	RunCommandError = errors.New("Error happened when execute command")
)

type BaseAnalyzer struct {
	Target string
	Config config.Config
}

func (this BaseAnalyzer) PickIP(target string) string {
	return REGEX_IPADDR.FindString(target)
}

func (this BaseAnalyzer) IsPublicIP(target string) bool {
	return !REGEX_PRIVATE_IPADDR.MatchString(target)
}

type StaticAnalyzer struct {
	BaseAnalyzer
}

func (s StaticAnalyzer) IsUpxPacked() bool {
	output, err := osutil.RunCmd(5, ".", "/usr/bin/upx", "-q", "-t", s.Target)
	if err != nil {
		return false
	}
	if strings.Contains(output, "[OK]") {
		return true
	} else {
		return false
	}
}

func (s StaticAnalyzer) GetFileInfo() (string, error) {
	abspath, err := filepath.Abs(s.Target)
	if err != nil {
		return "", err
	}
	if isExist, err := fileutil.IsFileExist(abspath); !isExist {
		return "", err
	}
	output, err := osutil.RunCmd(5, ".", "/usr/bin/file", abspath)
	return strings.TrimSpace(strings.Split(output, ":")[1]), err
}
