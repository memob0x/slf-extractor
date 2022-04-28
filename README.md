# SLF extractor

In [Jagged Alliance 2](https://it.wikipedia.org/wiki/Jagged_Alliance_2), and perhaps in any [Sir-Tech](https://en.wikipedia.org/wiki/Sir-Tech) game installation folder, the resources files (textures, sounds...) are located in a folder called "Data", but they're all bundled in a proprietary file-format with "slf" extension, thus the game real assets are inacessible; this application aims to browse and export those assets.

## Usage

### CLI

[Download](https://github.com/memob0x/slf-extractor/releases) the latest version of the standalone CLI and launch it.

```console
./slf-exporter-linux-adm64 ./FILE.slf ./dir
```

### Extra

An early version written in node (command-line only) is also [available](https://github.com/memob0x/slf-extractor/releases/tag/v0.1.0).

## Dev

To build the cli application

```console
./build.bash
```

To build the gui application

```console
./build.bash --with-gui
```

To test the source utilities

```console
go test -v ./utils
```

## Credits

Thanks to [Ja2-stracciatella](https://github.com/ja2-stracciatella).
