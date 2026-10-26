# ASCII-ART

## Description
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big :).

## Project Structure
```bash
.
|____README.md
|____shadow.txt
|____main_test.go
|____standard.txt
|____thinkertoy.txt
|____tranformToAscii.go
|____main.go
|____go.mod
```

## Requirement
Requires **Go 1.24.4** or newer

## Installation
```bash
git clone <ascii-art repo>
cd ascii-art
```
## Features
- Converts text to ASCII art
- Supports multiple lines with `\n`
- Supports 3 banners: `standard`, `shadow`, `thinkertoy`
- This project handle an input with numbers, letters,  spaces, special characters and \n.

## Usage
```bash
$ go run . "" | cat -e
$ go run . "\n" | cat -e
$
$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
$ go run . "hello" | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$ go run . "HeLlO" | cat -e
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
$ go run . "Hello There" | cat -e
 _    _          _   _                 _______   _                           $
| |  | |        | | | |               |__   __| | |                          $
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| $
                                                                             $
                                                                             $
$ go run . "1Hello 2There" | cat -e
     _    _          _   _                         _______   _                           $
 _  | |  | |        | | | |                ____   |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $
                                                                                         $
                                                                                         $
$ go run . "{Hello There}" | cat -e
   __  _    _          _   _                 _______   _                           __    $
  / / | |  | |        | | | |               |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                              /_/   $
                                                                                         $
$ go run . "Hello\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
$ go run . "Hello\n\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
$
```

## Banner
By default, the program uses the `standard.txt` banner. Two more banners are included:

- [standard.txt](https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/standard.txt)
- [shadow.txt](https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/shadow.txt)
- [thinkertoy.txt](https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/thinkertoy.txt)

To use another banner, change the file name in `main.go`:

```go
fileData, err := os.ReadFile("shadow.txt")
```

## Limitations

- Only printable ASCII characters (32 to 126) are supported.
- The banner file must keep the same format: 8 art rows per character plus 1 separator line.
## Author
**Abdelkhaleq Elalaoui**
- **Github:** @abdeleaiso

## Licence
This project is licenced under the MIT.
See the [<a href="https://github.com/abdeleaiso/ascii-art/blob/main/LICENSE">LICENCE</a>](LICENCE) file for details