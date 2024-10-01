## Features

- Transcribe audio and output to stdout/json (supports multiple file input)
- Helps with drafting captions for your editing workflow

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

### Get Started

1. [Download](https://github.com/prattlOrg/prattl/releases) (or [build](https://github.com/prattlOrg/prattl)) binary to chosen location
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
