# phaseant/lzw-cli

This is a cli for Lempel–Ziv–Welch algorithm, made for my information coding class.

## Usage

`$ make dict`

Generate base dictionary for text

`$ make encode`

Generate an encoded text sequence. Dict have to be generated before.

`$ make decode`

Decode encoded text using base dictionary. All missing fragments are generated in process.
