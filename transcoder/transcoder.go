package transcoder

import (
	"io"
	"os/exec"

	"github.com/xfrr/goffmpeg"
	"github.com/xfrr/goffmpeg/media"
)

// Transcoder Main struct
type Transcoder struct {
	stdErrPipe         io.ReadCloser
	stdStdinPipe       io.WriteCloser
	process            *exec.Cmd
	mediafile          *media.File
	configuration      goffmpeg.Configuration
	whiteListProtocols []string
}

// SetProcessStderrPipe Set the STDERR pipe
func (t *Transcoder) SetProcessStderrPipe(v io.ReadCloser) {
	_ = "STUB: not implemented"

	// SetProcessStdinPipe Set the STDIN pipe
	return
}

func (t *Transcoder) SetProcessStdinPipe(v io.WriteCloser) {
	_ = "STUB: not implemented"

	// SetProcess Set the transcoding process
	return
}

func (t *Transcoder) SetProcess(cmd *exec.Cmd) {
	_ = "STUB: not implemented"

	// SetMediaFile Set the media file
	return
}

func (t *Transcoder) SetMediaFile(v *media.File) {
	_ = "STUB: not implemented"

	// SetConfiguration Set the transcoding configuration
	return
}

func (t *Transcoder) SetConfiguration(v goffmpeg.Configuration) { _ = "STUB: not implemented"; return }

func (t *Transcoder) SetWhiteListProtocols(availableProtocols []string) {
	_ = "STUB: not implemented"
	return
}

// Process Get transcoding process
func (t Transcoder) Process() *exec.Cmd {
	_ = "STUB: not implemented"

	// MediaFile Get the ttranscoding media file.
	return nil
}

func (t Transcoder) MediaFile() *media.File {
	_ = "STUB: not implemented"

	// FFmpegExec Get FFmpeg Bin path
	return nil
}

func (t Transcoder) FFmpegExec() string { _ = "STUB: not implemented"; return "" }

// FFprobeExec Get FFprobe Bin path
func (t Transcoder) FFprobeExec() string { _ = "STUB: not implemented"; return "" }

// GetCommand Build and get command
func (t Transcoder) GetCommand() []string { _ = "STUB: not implemented"; return nil }

// InitializeEmptyTranscoder initializes the fields necessary for a blank transcoder
func (t *Transcoder) InitializeEmptyTranscoder() error { _ = "STUB: not implemented"; return nil }

// Set new File

// Set transcoder configuration

// SetInputPath sets the input path for transcoding
func (t *Transcoder) SetInputPath(inputPath string) error { _ = "STUB: not implemented"; return nil }

// SetOutputPath sets the output path for transcoding
func (t *Transcoder) SetOutputPath(inputPath string) error { _ = "STUB: not implemented"; return nil }

// CreateInputPipe creates an input pipe for the transcoding process
func (t *Transcoder) CreateInputPipe() (*io.PipeWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateOutputPipe creates an output pipe for the transcoding process
func (t *Transcoder) CreateOutputPipe(containerFormat string) (*io.PipeReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize Init the transcoding process
func (t *Transcoder) Initialize(inputPath string, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Set new File

// Set transcoder configuration

// Run Starts the transcoding process
func (t *Transcoder) Run(progress bool) <-chan error { _ = "STUB: not implemented"; return nil }

// Set the stdinPipe in case we need to stop the transcoding

// If the user has requested progress, we send it to them on a Buffer

// If an input pipe has been set, we set it as stdin for the transcoding

// If an output pipe has been set, we set it as stdout for the transcoding

// Stop Ends the transcoding process
func (t *Transcoder) Stop() error { _ = "STUB: not implemented"; return nil }

// Output Returns the transcoding progress channel
func (t Transcoder) Output() <-chan Progress { _ = "STUB: not implemented"; return nil }

//windows \r\n
//so  first \r and then \n can remove unexpected line break

// We have a cr terminated line

// We have a full newline-terminated line.

//live stream check

// Progress calculation

func (t *Transcoder) closePipes() { _ = "STUB: not implemented"; return }
