# SLF extractor

[![CI](https://github.com/memob0x/slf-extractor/workflows/Go/badge.svg)](https://github.com/memob0x/slf-extractor/actions?query=workflow%3AGo) [![Go Report Card](https://goreportcard.com/badge/github.com/memob0x/slf-extractor)](https://goreportcard.com/report/github.com/memob0x/slf-extractor) [![GoDoc](https://godoc.org/github.com/memob0x/slf-extractor?status.svg)](http://godoc.org/github.com/memob0x/slf-extractor) [![version](https://img.shields.io/github/v/tag/memob0x/slf-extractor?label=version)]()

In [Jagged Alliance 2](https://it.wikipedia.org/wiki/Jagged_Alliance_2), and perhaps in any [Sir-Tech](https://en.wikipedia.org/wiki/Sir-Tech) game installation folder, the resources files (textures, sounds...) are located in a folder called "Data", but they're all bundled in a proprietary file-format with "slf" extension, thus the game real assets are inacessible; this application aims to browse and export those assets.

## Usage

### GUI
[Download](https://github.com/memob0x/slf-extractor/releases) the latest version of the GUI application for your desktop environment and launch it with double click.

![gui](https://github.com/memob0x/slf-extractor/blob/master/assets/screenshot.png?raw=true)

### CLI

[Download](https://github.com/memob0x/slf-extractor/releases) the latest version of the CLI application for your desktop environment and launch it with your preferred terminal emulator.

```console
./slf-extractor-cli-linux-adm64 ./file.slf ./dir
```

### Extra

An early version written in node (command-line only) is also [available](https://github.com/memob0x/slf-extractor/releases/tag/v0.1.0).

## Dev

To build the cli application

```console
./build.bash
```

To build the gui application

[Docker](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository) is needed in order to use gui build tools.

```console
sudo /etc/init.d/docker start
```

```console
go install github.com/fyne-io/fyne-cross@latest
```

```console
./build.bash --with-gui
```

To test the source utilities

```console
go test -v ./utils
```

## Credits

Thanks to [Ja2-stracciatella](https://github.com/ja2-stracciatella).
