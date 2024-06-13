# fmove

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![Lines of code][lines-svg]][lines-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

✨ **`xuender/fmove` is a file move tool. It can move files and directories.**

## 🚀 Install

```shell
go install github.com/xuender/fmove@latest
```

## 💡 Usage

```shell
mkdir demo
cd demo
echo -e "[[target]]\npath = \"Images\"\ncondition = \"type==image\"\nsubDir = \"yy\"" > fm.toml
cp ../*.jpg .

fmove
```

### fm.toml

Move jpeg files to directory target/24/06.

```toml
[[target]]
path = "target"
condition = "type==image && sub=='jpeg'"
subDir = "yy+'/'+MM"
```

### Meta

| Keyword   | Description                                      | Example |
| --------- | ------------------------------------------------ | ------- |
| type      | File type, such as image, video, audio, archive. | image   |
| subtype   | File suffix, such as jpeg, mp4, mp3, zip.        | jpeg    |
| sub       | alias for subtype                                | jpeg    |
| extension | File extension, such as .jpg, .mp4, .mp3, .zip   | .jpg    |
| ext       | alias for extension                              | .jpg    |
| size      | File size, unit is byte.                         | 1024    |
| width     | Image or Video width, unit is pixel.             | 1920    |
| height    | Image or Video height, unit is pixel.            | 1080    |
| duration  | Video or Audio duration, unit is second.         | 60      |
| dur       | alias for duration                               | 60      |
| channels  | Audio or Video channels, 1 or 2.                 | 2       |
| chan      | alias for channels                               | 2       |


### Datetime

| Keyword | Description                                     | Example |
| ------- | ----------------------------------------------- | ------- |
| yyyy    | Full four-digit year                            | 2024    |
| yy      | Abbreviated two-digit year                      | 24      |
| MM      | Month, two digits with leading zeros            | 06      |
| M       | Month, one digit                                | 6       |
| dd      | Day of the month, two digits with leading zeros | 13      |
| d       | Day of the month, one digit                     | 13      |
| HH      | Hour in 24-hour, two digits with leading zeros  | 16      |
| H       | Hour in 24-hour, one digit                      | 16      |
| hh      | Hour in 12-hour, two digits with leading zeros  | 04      |
| h       | Hour in 12-hour, one digit                      | 4       |
| mm      | Minute, two digits with leading zeros           | 30      |
| m       | Minute, one digit                               | 30      |
| ss      | Second, two digits with leading zeros           | 45      |
| s       | Second, one digit                               | 45      |

## 👤 Contributors

![Contributors][contributors-svg]

## 📝 License

© ender, 2024~time.Now

[MIT LICENSE][license-url]

[action-url]: https://github.com/xuender/fmove/actions
[action-svg]: https://github.com/xuender/fmove/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/github.com/xuender/fmove
[goreport-svg]: https://goreportcard.com/badge/github.com/xuender/fmove

[godoc-url]: https://godoc.org/github.com/xuender/fmove
[godoc-svg]: https://godoc.org/github.com/xuender/fmove?status.svg

[license-url]: https://github.com/xuender/fmove/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg

[contributors-svg]: https://contrib.rocks/image?repo=xuender/fmove

[lines-svg]: https://sloc.xyz/github/xuender/fmove
[lines-url]: https://github.com/boyter/scc
