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

const (
	EXEC_DIR     = "."
	EXEC_TIMEOUT = 5
	BIN_UPX      = "/usr/bin/upx"
	BIN_LDD      = "/usr/bin/ldd"
	BIN_FILE     = "/usr/bin/file"
	BIN_STRINGS  = "/usr/bin/strings"
	BIN_EXIFTOOL = "/usr/bin/vendor_perl/exiftool"
)

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
	output, err := osutil.RunCmd(EXEC_TIMEOUT, EXEC_DIR, BIN_UPX, "-q", "-t", s.Target)
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
	output, err := osutil.RunCmd(EXEC_TIMEOUT, EXEC_DIR, BIN_FILE, abspath)
	return strings.TrimSpace(strings.Split(output, ":")[1]), err
}

func (s StaticAnalyzer) GetExifInfo() (map[string]string, error) {
	abspath, err := filepath.Abs(s.Target)
	if err != nil {
		return nil, err
	}
	output, err := osutil.RunCmd(EXEC_TIMEOUT, EXEC_DIR, BIN_EXIFTOOL, abspath)
	result := make(map[string]string)
	for _, line := range strings.Split(output, "\n") {
		words := strings.SplitN(line, ":", 2)
		if len(words) != 2 {
			continue
		}
		key := strings.TrimSpace(words[0])
		value := strings.TrimSpace(words[1])
		result[key] = value
	}

	return result, nil
}

func (s StaticAnalyzer) GetAsciiStr() ([]string, error) {
	abspath, err := filepath.Abs(s.Target)
	if err != nil {
		return nil, err
	}
	output, err := osutil.RunCmd(EXEC_TIMEOUT, EXEC_DIR, BIN_STRINGS, "-a", abspath)
	return strings.Split(output, "\n"), err
}

func (s StaticAnalyzer) GetUnicode() ([]string, error) {
	abspath, err := filepath.Abs(s.Target)
	if err != nil {
		return nil, err
	}
	output, err := osutil.RunCmd(EXEC_TIMEOUT, EXEC_DIR, BIN_STRINGS, "-a", "-el", abspath)
	return strings.Split(output, "\n"), err
}

type LibraryItem struct {
	name string
	path string
}

func (s StaticAnalyzer) GetLibraryDepends() ([]LibraryItem, error) {
	var depends []LibraryItem
	abspath, err := filepath.Abs(s.Target)
	if err != nil {
		return depends, err
	}
	output, err := osutil.RunCmd(EXEC_TIMEOUT, EXEC_DIR, BIN_LDD, abspath)

	for _, line := range strings.Split(output, "\n") {
		if line == "" {
			continue
		}

		libItem := LibraryItem{}
		if strings.Contains(line, "=>") {
			words := strings.Split(line, "=>")
			libItem.name = strings.TrimSpace(words[0])
			libItem.path = strings.TrimSpace(
				words[1][:strings.Index(words[1], "(0x")])
		} else {
			libItem.name = strings.TrimSpace(
				line[:strings.Index(line, "(0x")])
			libItem.path = ""
		}
		depends = append(depends, libItem)
	}
	return depends, nil
}
