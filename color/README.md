## ascii-art-color

## Overview

This project extends the original ascii-art functionality by adding support for colorizing specific substrings in the input text. You can highlight parts of your ASCII art or text output by applying colors to chosen substrings using the command-line interface.

### Objectives

The program allows you to colorize substrings within a string. The coloring is controlled via the `--color` flag, and you can specify the color and the substring to be colored. The supported colors can be in various notations such as `color names`, `RGB`, or `ANSI escape codes`.

### Usage

```shell
$ go run . --color=<color> <substring to be colored> "your string here"
```
- `--color=<color>`: Specifies the color to apply to the substring. You can use color names (e.g., red, blue), RGB values (e.g., rgb(255,0,0)), or other supported color notations.
- `<substring to be colored>`: The specific substring within the input text that you want to colorize.
- `"your string here"`: The string in which the coloring will be applied.

## Examples

1. ## Colorizing a specific substring
```shell
$ go run . --color=red kit "a king kitten have kit"
```
This command will color the substring kit in both kitten and kit with the color red.

2. ## Colorizing the entire string
If no substring is provided, the entire string will be colored.
```shell
$ go run . --color=blue "a king kitten have kit"
```
This command will color the entire string with blue.

### Usage Messages

- ## Correct Usage
For the example above, the substring `kit` in the word `kitten` and the word `kit` at the end should be colored.
```shell
$ go run . --color=red kit "a king kitten have kit"
```

- ## Incorrect Usage

1. If the `--color` flag is missing or not formatted correctly, or if the command is used incorrectly, the program will display the following usage message:
```console
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

2. If the `"your string here"` have a unprintable ASCII character, the program will display the following error message:
```console
Character 'the unprintable character you typed' is not a printable ASCII character
```

3. If the `<substring to be colored>` have a unprintable ASCII character, or have a white spaces, the program will display the following error message:
```console
Character 'the unprintable character you typed' is not a printable ASCII character

[Check README file]
```

### Additional Features

- The program is designed to work with other optional projects related to ASCII art. It supports various [OPTION] flags and/or [BANNER] configurations if implemented.

- The functionality to run the program with a single [STRING] argument remains intact, allowing basic text output without color customization.

### Getting Started

To run the program:
1. Ensure you have Go installed on your machine.
2. Clone the repository or navigate to the project directory.
3. Use the go run . command with the appropriate flags and arguments as described above.

Feel free to explore different color notations and experiment with colorizing various parts of your text. Enjoy enhancing your ASCII art and text outputs with vibrant colors!