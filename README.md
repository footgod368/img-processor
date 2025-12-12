# Introduction

This is a simple image processor cli.

Supported actions:
- invert
- mirror
- print

Supported image formats:
- jpeg
- png

# Prerequisites

- Go 1.18 or higher

# Usage

```bash
go build -o img-processor
./img-processor invert asset/test.jpg
```

See the help for more information:
```bash
./img-processor -h
```

Leverage the `-i` flag to print the output path, so that you can pipe the output to other commands.
```bash
./img-processor invert asset/test.jpg -i | xargs ./img-processor mirror
```


The output image will be saved in `out/inverted-test.jpg`

# Examples
## invert
![test](asset/test.jpg)
![inverted-test](out/inverted-test.jpg)
## mirror
![test](asset/test.jpg)
![mirrored-test](out/mirrored-test.jpg)
## pipe
![test](asset/test.jpg)
![mirrored-inverted-test](out/mirrored-inverted-test.jpg)