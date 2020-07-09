package config

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	LogDir               string
	FileLog              string
	StaticLog            string
	DynamicLog           string
	ExecHome             string
	TimeLimitDynamic     int32
	StaticFinishedFName  string
	DynamicFinishedFName string
	StringsLimit         int32
	TargetLoader         string
	SysdigPluginDir      string
	LdDebugLog           string
	IsInplace            bool
	DecompressLimit      int32
	EnablePrefixRemove   bool
	TSharkUser           string
	EnableINetSim        bool
	INetSimDataDir       string
	INetSimCfgPath       string
	INetSimLogDir        string
	INetSimLogReportDir  string
	KernelLogPath        string
	NetEth0              string
	NetEth1              string
	TraceType            string
	TcpdumpLimit         int32
	SysdigLimit          int32
	TraceLimit           int32
	YaraRulesData        string
	FakeDnsIP            string
	EnableMemAnalysis    bool
	LimeSrcPath          string
	MemDumpPath          string
	UpdateVolProfilePath string
	VolProfileName       string
}

var (
	NoFileSpecified  = errors.New("no file specified")
	JsonDecodeFailed = errors.New("failed to decode json")
)

func LoadConfig(config string) (*Config, error) {
	if config == "" {
		return nil, NoFileSpecified
	}

	f, _ := os.Open(config)
	defer f.Close()

	decoder := json.NewDecoder(f)
	conf := &Config{
		LogDir:               "log",
		FileLog:              "system.log",
		StaticLog:            "output.static",
		DynamicLog:           "output.dynamic",
		ExecHome:             "workspace",
		TimeLimitDynamic:     10,
		StaticFinishedFName:  "TagFile_S.txt",
		DynamicFinishedFName: "MonitorComplete_TagFile.txt",
		StringsLimit:         10,
		LdDebugLog:           "ld_debug.log",
		IsInplace:            true,
		DecompressLimit:      128,
		EnablePrefixRemove:   false,
		EnableINetSim:        false,
		TraceType:            "auto",
		TcpdumpLimit:         500,
		SysdigLimit:          50000,
		TraceLimit:           50000,
		EnableMemAnalysis:    false,
	}

	if err := decoder.Decode(&conf); err != nil {
		return nil, JsonDecodeFailed
	}

	return conf, nil
}
