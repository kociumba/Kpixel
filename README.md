# Pixel sorting algorith in go

I got inspired by the youtube video by Acerola on pixel sorting.

So i made my own implementation in go, this is by no means feature complete.

## Usage

`-sort` defines the way we sort pixels options are:

- `column` (sorts pixels in respective columns, preserves vertical elements)
- `row` (sorts pixels in respective rows, preserves horizontal elements)

Output always goes into the folder of the original image with a .sorted extension to indicate that it has been sorted.

For now only sorting with hue values is supported.

I will add more options in the future when I have some time.