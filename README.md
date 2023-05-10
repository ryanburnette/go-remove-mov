# [remove-mov](https://github.com/ryanburnette/go-remove-mov)

`remove-mov` is a utility designed to clean up image directories downloaded from
Synology by identifying and removing redundant .mov files associated with live
images, while preserving standalone .mov video files. It scans for .heic and
.mov files with matching names and deletes the .mov file, ensuring that only the
necessary files are retained.

## Usage

```shell
remove-mov [directory] [-d]
```

- `[directory]` optional, assumes cwd
- `-d` dry run
