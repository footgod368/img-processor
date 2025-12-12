# Introduction

This is a simple image processor cli.

Currently only support invert jpg image.

# Prerequisites

- Go 1.18 or higher

# Usage

```bash
go build -o img-processor
./img-processor invert asset/test.jpg
```

Leverage the `-i` flag to print the output path, so that you can pipe the output to other commands.
```bash
./img-processor invert asset/test.jpg -i | xargs ./img-processor mirror
```


The output image will be saved in `out/inverted-test.jpg`

![test](asset/test.jpg)
![inverted-test](out/inverted-test.jpg)