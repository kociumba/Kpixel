# Pixel sorting algorith in go

I got inspired by the youtube video by Acerola on pixel sorting.

So i made my own implementation in go, this is by no means feature complete.

## Usage

`-sort` defines the way we sort pixels, options are:

> - `column` (sorts pixels in respective columns, preserves vertical elements)
> - `row` (sorts pixels in respective rows, preserves horizontal elements)
> - `random` (randomly sorts pixels in chunks)

`-chunk` number of chunks to devide the image in to when using random sort defaults to 10 (only relevant if using random sort)

> [!NOTE]
> if the chunk size is bigger than the width of the image in pixels the whole image gets randomised and essentially becomes noise

`-method` defines the value that is used to sort the pixels, options are:

> - `hue` (sorts pixels by their hue value)
> - `red` (sorts pixels by their red value)
> - `green` (sorts pixels by their green value)
> - `blue` (sorts pixels by their blue value)

> [!IMPORTANT]
>always pass the path to the image you want to sort as the last argument

Output always goes into the folder of the original image with a .sorted extension to indicate that it has been sorted.