# <p align="center">Welcome to</p>

<p align="center">
    <img src="Kpixel.svg" alt="Kpixel" title="Kpixel logo">
</p>

## Overview

I got inspired by the YouTube video by Acerola on pixel sorting.

So I made my own implementation in go, this is by no means feature complete.

Kpixel is a part of my K suite of tools that so far includes:
- [Ktool](https://github.com/kociumba/ktool)
- [Ksorter](https://github.com/kociumba/ksorter)
- [Kinjector](https://github.com/kociumba/Kinjector)
- [Kpixel](https://github.com/kociumba/kpixel) - this repo

## Installation

Right now the only way to install is to compile it yourself which only requires [go](https://go.dev/dl/).

I will create a scoop manifest for it in the future.

## Usage

Kpixel is a CLI tool for `-sort column` and `-sort row` the `-method` matters, for `-sort random` you can define the `-chunk` in pixels e.g. `-sort random -chunk 100`.

`-sort` defines the way we sort pixels, options are:

> - `column` (sorts pixels in respective columns, preserves vertical elements)
> - `row` (sorts pixels in respective rows, preserves horizontal elements)
> - `random` (randomly sorts pixels in chunks)

`-chunk` number of chunks to divide the image in to when using random sort defaults to 10 (only relevant if using random sort)

> [!NOTE]
> if the chunk size is bigger than the width of the image in pixels the whole image gets randomised and essentially becomes noise

`-method` defines the value that is used to sort the pixels, options are:

> - `hue` (very noisy)
> - `luminosity` (smooth and looks good)
> - `saturation` (kinda buggy, needs more testing)
> - `red` (looks good depending on the image)
> - `green` (looks good depending on the image)
> - `blue` (looks good depending on the image)

> [!IMPORTANT]
> always pass the path to the image you want to sort as the last argument
>
> if you don't pass the image path, a file picker will open prompting you to pick an image 

Output always goes into the folder of the original image with a .sorted extension to indicate that it has been sorted.