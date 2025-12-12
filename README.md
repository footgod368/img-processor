# Introduction

This is a simple image processor cli.

Currently only support invert jpg image.

# Prerequisites

- Go 1.18 or higher

# Run

```bash
go build -o img-processor
./img-processor invert asset/test.jpg
```

The output image will be saved in `out/inverted-test.jpg`

![test](asset/test.jpg)
![inverted-test](asset/inverted-test.jpg)