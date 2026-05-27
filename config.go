package goffmpeg

import (
	"context"
)

const (
	ffmpegCommand  = "ffmpeg"
	ffprobeCommand = "ffprobe"
)

type Configuration struct {
	ffprobeBinPath string
	ffmpegBinPath  string
}

func (cfg Configuration) FFmpegBinPath() string { _ = "STUB: not implemented"; return "" }

func (cfg Configuration) FFprobeBinPath() string { _ = "STUB: not implemented"; return "" }

func Configure(ctx context.Context) (Configuration, error) {
	_ = "STUB: not implemented"
	return *new(Configuration), nil
}

func normalizeBinPath(binPath string) string { _ = "STUB: not implemented"; return "" }

func lineSeparator() string { _ = "STUB: not implemented"; return "" }
