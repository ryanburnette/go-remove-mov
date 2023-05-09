# [remove-mov](https://github.com/ryanburnette/go-remove-mov)

Synology downloads of image directories provide a side-by-side .heic and .mov
file for live images. Videos are a single .mov file. This program looks for
side-by-side .heic and .mov files and deletes the associated .mov while leaving
standalone .mov files alone.

## Usage

The `[directory]` is assumed to be the current directory unless one is provided.

The `-d` flag gives only a dry run.

```shell
remove-mov [directory] [-d]
```
