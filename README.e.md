---
SingleSel1: echo '<input type="radio" name="Sex" value="F" />' | tee /tmp/cascadia.xml | cascadia -i -o -c 'input[name=Sex][value=F]'
SingleSel2: cascadia -i /tmp/cascadia.xml -o -c 'input[name=Sex][value=F]'
---

## {{toc 5}}
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## {{.Name}} - CSS selector CLI tool

The [Go Cascadia package](https://github.com/andybalholm/cascadia) implements CSS selectors for html. This is the command line tool, started as a thin wrapper around that package, but growing into a better tool to test CSS selectors without writing Go code:

## Usage

### $ {{exec "cascadia" | color "sh"}}

Its output has two modes, _none-block selection mode_ and _block selection mode_, depending on whether the `--piece` parameter is given on the command line or not.

For details about the concept of block and pieces, check out [andrew-d/goscrape](https://github.com/andrew-d/goscrape) (in fact, `cascadia` was initially developed just for it, so that I don't need to tweak Go code, build & run it just to test out the block and pieces selectors). Here is the exception:

- Inside each page, there's 1 or more *blocks* - some logical method of splitting up a page into subcomponents.
- Inside each block, you define some number of *pieces* of data that you wish
  to extract.  Each piece consists of a name, a selector, and what data to
  extract from the current block.

This all sounds rather complicated, but in practice it's quite simple. See the next section for details.

In summary,

- The none-block selection mode will output the selection as HTML source by default
  * but if `-t`, or `--text` cli option is provided, the none-block selection mode will [output as text](https://github.com/suntong/cascadia/issues/6#issuecomment-980757881) instead.
    - By default, such text output will get their leading and trailing white space trimmed.
    - However, if `-R`, or `--Raw` cli option is provided, no trimming will be done.
- The block selection mode will output HTML as text in a `tsv`/`csv` table form by default
  * if the `--piece` selection is prefixed with `RAW:`, then that specific block selection will output in HTML instead. See the following for details.

### Examples

All the three `-i -o -c` options are required. By default it reads from `stdin` and output to `stdout`:

```sh
$ {{shell .SingleSel1}}
```

Either the input or the output can be followed by a file name:


```sh
$ {{shell .SingleSel2}}
```


```sh
$ cascadia -i /tmp/cascadia.xml -c 'input[name=Sex][value=F]' -o /tmp/out.html
1 elements for 'input[name=Sex][value=F]':

$ cat /tmp/out.html
<input type="radio" name="Sex" value="F"/>
```

More other options can be applied too:

```sh
# using --wrap-html
$ cascadia -i /tmp/cascadia.xml -c 'input[name=Sex][value=F]' -o /tmp/out.html -w
1 elements for 'input[name=Sex][value=F]':

$ cat /tmp/out.html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<base href="">

</head>
<body>
<input type="radio" name="Sex" value="F"/>
</body>

# using --wrap-html with --style
$ cascadia -i /tmp/cascadia.xml -c 'input[name=Sex][value=F]' -o /tmp/out.html -w -y '<link rel="stylesheet" href="styles.css">'
1 elements for 'input[name=Sex][value=F]':

$ cat /tmp/out.html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<base href="">
<link rel="stylesheet" href="styles.css">
</head>
<body>
<input type="radio" name="Sex" value="F"/>
</body>
```

- For more on using the `--style` option, check out ["adding styles"](https://github.com/suntong/cascadia/wiki/Adding-styles).
- For more examples, check out the [wiki](https://github.com/suntong/cascadia/wiki/), which includes but not limits to, 

  * [None-block selection mode](https://github.com/suntong/cascadia/wiki#none-block-selection-mode)
    * [Multi-selection](https://github.com/suntong/cascadia/wiki#multi-selection)
  * [Block selection mode](https://github.com/suntong/cascadia/wiki#block-selection-mode)
    * [Block selection mode HTML output](https://github.com/suntong/cascadia/wiki#block-selection-mode-html-output)
    * [Block selection mode table output](https://github.com/suntong/cascadia/wiki#block-selection-mode-table-output)
    * [Attribute selection](https://github.com/suntong/cascadia/wiki#attribute-selection)
    * [Twitter Search](https://github.com/suntong/cascadia/wiki#twitter-search)
  * [Reconstruct the separated pages](https://github.com/suntong/cascadia/wiki#reconstruct-the-separated-pages)
  * [More On CSS Selector](https://github.com/suntong/cascadia/wiki#more-on-css-selector)

