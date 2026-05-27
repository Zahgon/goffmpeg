package media

import (
	"io"
)

type File struct {
	aspect                string
	resolution            string
	videoBitRate          string
	videoBitRateTolerance int
	videoMaxBitRate       int
	videoMinBitrate       int
	videoCodec            string
	vframes               int
	frameRate             int
	audioRate             int
	maxKeyframe           int
	minKeyframe           int
	keyframeInterval      int
	audioCodec            string
	audioBitrate          string
	audioChannels         int
	audioVariableBitrate  bool
	bufferSize            int
	threadset             bool
	threads               int
	preset                string
	tune                  string
	audioProfile          string
	videoProfile          string
	target                string
	duration              string
	durationInput         string
	seekTime              string
	qscale                uint32
	crf                   uint32
	strict                int
	singleFile            int
	muxDelay              string
	seekUsingTsInput      bool
	seekTimeInput         string
	inputPath             string
	inputPipe             bool
	inputPipeReader       io.ReadCloser
	inputPipeWriter       io.Writer
	outputPipe            bool
	outputPipeReader      io.Reader
	outputPipeWriter      io.WriteCloser
	movFlags              string
	hideBanner            bool
	outputPath            string
	outputFormat          string
	copyTs                bool
	nativeFramerateInput  bool
	inputInitialOffset    string
	rtmpLive              string
	hlsPlaylistType       string
	hlsListSize           int
	hlsSegmentDuration    int
	hlsMasterPlaylistName string
	hlsSegmentFilename    string
	httpMethod            string
	httpKeepAlive         bool
	hwaccel               string
	streamIds             map[int]string
	metadata              Metadata
	videoFilter           string
	audioFilter           string
	skipVideo             bool
	skipAudio             bool
	compressionLevel      int
	mapMetadata           string
	tags                  map[string]string
	encryptionKey         string
	movflags              string
	bframe                int
	pixFmt                string
	rawInputArgs          []string
	rawOutputArgs         []string
}

/*** SETTERS ***/
func (m *File) SetAudioFilter(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoFilter(v string) {
	_ = "STUB: not implemented"

	// Deprecated: Use SetVideoFilter instead.
	return
}

func (m *File) SetFilter(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetAspect(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetResolution(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoBitRate(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoBitRateTolerance(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoMaxBitrate(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoMinBitRate(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoCodec(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetVframes(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetFrameRate(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetAudioRate(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetAudioVariableBitrate() { _ = "STUB: not implemented"; return }

func (m *File) SetMaxKeyFrame(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetMinKeyFrame(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetKeyframeInterval(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetAudioCodec(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetAudioBitRate(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetAudioChannels(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetPixFmt(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetBufferSize(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetThreads(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetPreset(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetTune(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetAudioProfile(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetVideoProfile(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetDuration(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetDurationInput(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetSeekTime(v string) { _ = "STUB: not implemented"; return }

func (m *File) SetSeekTimeInput(v string) { _ = "STUB: not implemented"; return }

// Q Scale must be integer between 1 to 31 - https://trac.ffmpeg.org/wiki/Encode/MPEG-4
func (m *File) SetQScale(v uint32) { _ = "STUB: not implemented"; return }

func (m *File) SetCRF(v uint32) { _ = "STUB: not implemented"; return }

func (m *File) SetStrict(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetSingleFile(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetSeekUsingTsInput(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetCopyTs(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetInputPath(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetInputPipe(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetInputPipeReader(r io.ReadCloser) { _ = "STUB: not implemented"; return }

func (m *File) SetInputPipeWriter(w io.Writer) { _ = "STUB: not implemented"; return }

func (m *File) SetOutputPipe(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetOutputPipeReader(r io.Reader) { _ = "STUB: not implemented"; return }

func (m *File) SetOutputPipeWriter(w io.WriteCloser) { _ = "STUB: not implemented"; return }

func (m *File) SetMovFlags(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetHideBanner(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetMuxDelay(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetOutputPath(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetOutputFormat(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetNativeFramerateInput(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetRtmpLive(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetHlsListSize(val int) { _ = "STUB: not implemented"; return }

func (m *File) SetHlsSegmentDuration(val int) { _ = "STUB: not implemented"; return }

func (m *File) SetHlsPlaylistType(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetHlsMasterPlaylistName(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetHlsSegmentFilename(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetHttpMethod(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetHttpKeepAlive(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetHardwareAcceleration(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetInputInitialOffset(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetStreamIds(val map[int]string) { _ = "STUB: not implemented"; return }

func (m *File) SetSkipVideo(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetSkipAudio(val bool) { _ = "STUB: not implemented"; return }

func (m *File) SetMetadata(v Metadata) { _ = "STUB: not implemented"; return }

func (m *File) SetCompressionLevel(val int) { _ = "STUB: not implemented"; return }

func (m *File) SetMapMetadata(val string) { _ = "STUB: not implemented"; return }

func (m *File) SetTags(val map[string]string) { _ = "STUB: not implemented"; return }

func (m *File) SetBframe(v int) { _ = "STUB: not implemented"; return }

func (m *File) SetRawInputArgs(args []string) { _ = "STUB: not implemented"; return }

func (m *File) SetRawOutputArgs(args []string) {
	_ = "STUB: not implemented"
	return

	/*** GETTERS ***/
}

// Deprecated: Use VideoFilter instead.
func (m *File) Filter() string { _ = "STUB: not implemented"; return "" }

func (m *File) VideoFilter() string { _ = "STUB: not implemented"; return "" }

func (m *File) AudioFilter() string { _ = "STUB: not implemented"; return "" }

func (m *File) Aspect() string { _ = "STUB: not implemented"; return "" }

func (m *File) Resolution() string { _ = "STUB: not implemented"; return "" }

func (m *File) VideoBitrate() string { _ = "STUB: not implemented"; return "" }

func (m *File) VideoBitRateTolerance() int { _ = "STUB: not implemented"; return 0 }

func (m *File) VideoMaxBitRate() int { _ = "STUB: not implemented"; return 0 }

func (m *File) VideoMinBitRate() int { _ = "STUB: not implemented"; return 0 }

func (m *File) VideoCodec() string { _ = "STUB: not implemented"; return "" }

func (m *File) Vframes() int { _ = "STUB: not implemented"; return 0 }

func (m *File) FrameRate() int { _ = "STUB: not implemented"; return 0 }

func (m *File) GetPixFmt() string { _ = "STUB: not implemented"; return "" }

func (m *File) AudioRate() int { _ = "STUB: not implemented"; return 0 }

func (m *File) MaxKeyFrame() int { _ = "STUB: not implemented"; return 0 }

func (m *File) MinKeyFrame() int { _ = "STUB: not implemented"; return 0 }

func (m *File) KeyFrameInterval() int { _ = "STUB: not implemented"; return 0 }

func (m *File) AudioCodec() string { _ = "STUB: not implemented"; return "" }

func (m *File) AudioBitrate() string { _ = "STUB: not implemented"; return "" }

func (m *File) AudioChannels() int { _ = "STUB: not implemented"; return 0 }

func (m *File) BufferSize() int { _ = "STUB: not implemented"; return 0 }

func (m *File) Threads() int { _ = "STUB: not implemented"; return 0 }

func (m *File) Target() string { _ = "STUB: not implemented"; return "" }

func (m *File) Duration() string { _ = "STUB: not implemented"; return "" }

func (m *File) DurationInput() string { _ = "STUB: not implemented"; return "" }

func (m *File) SeekTime() string { _ = "STUB: not implemented"; return "" }

func (m *File) Preset() string { _ = "STUB: not implemented"; return "" }

func (m *File) AudioProfile() string { _ = "STUB: not implemented"; return "" }

func (m *File) VideoProfile() string { _ = "STUB: not implemented"; return "" }

func (m *File) Tune() string { _ = "STUB: not implemented"; return "" }

func (m *File) SeekTimeInput() string { _ = "STUB: not implemented"; return "" }

func (m *File) QScale() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *File) CRF() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *File) Strict() int { _ = "STUB: not implemented"; return 0 }

func (m *File) SingleFile() int { _ = "STUB: not implemented"; return 0 }

func (m *File) MuxDelay() string { _ = "STUB: not implemented"; return "" }

func (m *File) SeekUsingTsInput() bool { _ = "STUB: not implemented"; return false }

func (m *File) CopyTs() bool { _ = "STUB: not implemented"; return false }

func (m *File) InputPath() string { _ = "STUB: not implemented"; return "" }

func (m *File) InputPipe() bool { _ = "STUB: not implemented"; return false }

func (m *File) InputPipeReader() io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (m *File) InputPipeWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (m *File) OutputPipe() bool { _ = "STUB: not implemented"; return false }

func (m *File) OutputPipeReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (m *File) OutputPipeWriter() io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (m *File) MovFlags() string { _ = "STUB: not implemented"; return "" }

func (m *File) HideBanner() bool { _ = "STUB: not implemented"; return false }

func (m *File) OutputPath() string { _ = "STUB: not implemented"; return "" }

func (m *File) OutputFormat() string { _ = "STUB: not implemented"; return "" }

func (m *File) NativeFramerateInput() bool { _ = "STUB: not implemented"; return false }

func (m *File) RtmpLive() string { _ = "STUB: not implemented"; return "" }

func (m *File) HlsListSize() int { _ = "STUB: not implemented"; return 0 }

func (m *File) HlsSegmentDuration() int { _ = "STUB: not implemented"; return 0 }

func (m *File) HlsMasterPlaylistName() string { _ = "STUB: not implemented"; return "" }

func (m *File) HlsSegmentFilename() string { _ = "STUB: not implemented"; return "" }

func (m *File) HlsPlaylistType() string { _ = "STUB: not implemented"; return "" }

func (m *File) InputInitialOffset() string { _ = "STUB: not implemented"; return "" }

func (m *File) HttpMethod() string { _ = "STUB: not implemented"; return "" }

func (m *File) HttpKeepAlive() bool { _ = "STUB: not implemented"; return false }

func (m *File) HardwareAcceleration() string { _ = "STUB: not implemented"; return "" }

func (m *File) StreamIds() map[int]string { _ = "STUB: not implemented"; return nil }

func (m *File) SkipVideo() bool { _ = "STUB: not implemented"; return false }

func (m *File) SkipAudio() bool { _ = "STUB: not implemented"; return false }

func (m *File) Metadata() Metadata { _ = "STUB: not implemented"; return *new(Metadata) }

func (m *File) CompressionLevel() int { _ = "STUB: not implemented"; return 0 }

func (m *File) MapMetadata() string { _ = "STUB: not implemented"; return "" }

func (m *File) Tags() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *File) SetEncryptionKey(v string) { _ = "STUB: not implemented"; return }

func (m *File) EncryptionKey() string { _ = "STUB: not implemented"; return "" }

func (m *File) RawInputArgs() []string { _ = "STUB: not implemented"; return nil }

func (m *File) RawOutputArgs() []string {
	_ = "STUB: not implemented"
	return

	/** OPTS **/
	nil
}

func (m *File) ToStrCommand() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAudioFilter() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoFilter() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAspect() []string {
	_ = "STUB: not implemented"
	// Set aspect
	return nil
}

func (m *File) ObtainHardwareAcceleration() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainInputPath() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainInputPipe() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainOutputPipe() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainMovFlags() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHideBanner() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainNativeFramerateInput() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainOutputPath() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoCodec() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVframes() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainFrameRate() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAudioRate() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainResolution() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoBitRate() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAudioCodec() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAudioBitRate() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAudioChannels() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoMaxBitRate() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoMinBitRate() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainBufferSize() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoBitRateTolerance() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainThreads() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainTarget() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainDuration() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainDurationInput() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainKeyframeInterval() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainSeekTime() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainSeekTimeInput() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainPreset() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainTune() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainCRF() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainQScale() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainStrict() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainSingleFile() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainVideoProfile() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainAudioProfile() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainCopyTs() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainOutputFormat() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainMuxDelay() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainSeekUsingTsInput() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainRtmpLive() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHlsPlaylistType() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainInputInitialOffset() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHlsListSize() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHlsSegmentDuration() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHlsMasterPlaylistName() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHlsSegmentFilename() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHttpMethod() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainPixFmt() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainHttpKeepAlive() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainSkipVideo() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainSkipAudio() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainStreamIds() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainCompressionLevel() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainMapMetadata() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainEncryptionKey() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainBframe() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainTags() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainRawInputArgs() []string { _ = "STUB: not implemented"; return nil }

func (m *File) ObtainRawOutputArgs() []string { _ = "STUB: not implemented"; return nil }

func CheckFileType(streams []Streams) string { _ = "STUB: not implemented"; return "" }
