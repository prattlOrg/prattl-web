## Features

- Transcribe audio and output to stdout in json format (supports multiple file input)
- Helps with drafting captions for your editing workflow

### Prattl manages its own python distribution

When you run `prattl prepare` prattl installs a python distribution specific to your OS to your system, This means you don't need to manage Python dependencies or risk corrupting your existing environments.
`prattl clean` will completely remove this distribution.

### Transcribing

Under the hood, prattl is using [distil-whisper](https://huggingface.co/distil-whisper/distil-large-v3), which runs locally on your system. The better your GPU, the faster the transcription, if you do not have a GPU, it will use your CPU. To create a transcription, use the command:

```zsh
prattl transcribe <filepath/to/audio.mp3>
```

You can provide multiple file paths, and prattl will transcribe all of them as a single batch. This means the efficiency increases with more files!

Upon completion, the output will be a JSON object. For instance, if you run:

```zsh
prattl transcribe test1.mp3 test2.mp3 test3.mp3
```

the output will be:

```json
{
  "test1.mp3": "test1.mp3's transcription",
  "test2.mp3": "test2.mp3's transcription",
  "test3.mp3": "test3.mp3's transcription"
}
```

## Installation

### OS/CPU Architecture Support

- windows/386
- windows/amd64
- darwin/arm64
- darwin/amd64
- linux-gnu/arm64
- linux-gnu/amd64

> **_NOTE:_** CUDA architecture GPUs can take advantage of GPU acceleration for transcription

### Prerequisites

[ffmpeg](https://www.ffmpeg.org/) installed and included in `$PATH`

<!-- > **_NOTE:_** This step can be skipped if you are using a supported package manager to install prattl -->
<!--
### Install

Install with [Go](https://go.dev/):
`go install github.com/prattlOrg/prattl@latest`

Install on Windows with [Chocolatey](https://chocolatey.org/):

`choco install prattl`

Install on Mac/Linux with [Homebrew Taps](https://docs.brew.sh/Taps):

`brew tap prattlOrg/prattl`
`brew install prattl` -->

<!-- .deb, .rpm, .apk, .ipk -->

### Get Started

1. Install (or [build](https://github.com/prattlOrg/prattl/tree/main?tab=readme-ov-file#oscpu-architecture-support)) binary to chosen location
2. Include the path of the binary in `$PATH`
3. Run `prattl prepare` to install necessary python dependencies to `$HOME/.prattl` directory
4. Prattl is ready to be used!

## Usage

Help with any prattl command
`prattl help`

Remove the python distribution built at `$HOME/.prattl` by prattl
`prattl clean`

Prepare the python distribution at `$HOME/.prattl` required by prattl
`prattl prepare`

Transcribe provided audio by pathname, output transcription to stdout/json
`prattl transcribe <filepath/to/audio.mp3>`
