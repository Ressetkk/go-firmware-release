# go-firmware-release

It's a drop-in, multi OS, Go based replacement of `firmware-release.exe` binary
provided by Anycubic to prepare Marlin firmware that will erase EEPROM on your Vyper.
This program will prepare the `firmware.bin` the same way as the original binary and
will work on basically any OS, even in container environment.

I've created this because the original binary is ONLY for Windows, and
there is no source code around there. Also finding original exec is a hell of a work,
since I couldn't find it anywhere on Anycubic site and Vyper Repository.

You can also use it to create a binary that won't erase EEPROM. Use flag `-erase-eeprom=false`.
The input binary will just be copied and renamed.

If needed you can override name and extension using flags `-custom-name` and `-extension`.

You can easily use this tool in CI.

## Requirements
* Go 1.17+

## Usage
```
$ ./firmware-release -help
Usage of firmware-release:
  -custom-name string
        Set custom output file name.
  -erase-eeprom
        Create file that erases EEPROM. (default true)
  -extension string
        Define to which the file should be saved. (default "bin")  
```

## Example
```
$ ./firmware-release ./firmware.bin
firmware-release started at 2022-02-15 00:17:40.7464875 +0000 UTC
Input file: .\firmware.bin
Output file: main_board_20220215_erase_eeprom.bin
Checksum: e1eee506bca3c6ac193c3ebfbe430423
Success! 
```