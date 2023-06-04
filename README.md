# [synology-image-cleanup](https://github.com/ryanburnette/synology-image-cleanup)

`synology-image-cleanup` is a utility designed to clean up image directories
downloaded from Synology Photos by identifying and removing redundant .mov
files associated with live images, while preserving standalone .mov video
files. It scans for .heic and .mov files with matching names and deletes the
.mov file, ensuring that only the necessary files are retained.

## Usage

```shell
synology-image-cleanup [directory] [-d]
```

-   `[directory]` directory, default: cwd
-   `-d` dry run
