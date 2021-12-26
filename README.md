# cascadia
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/cascadia?status.svg)](http://godoc.org/github.com/suntong/cascadia)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/cascadia)](https://goreportcard.com/report/github.com/suntong/cascadia)
[![Build Status](https://github.com/suntong/cascadia/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/suntong/cascadia/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)



## TOC
- [cascadia - CSS selector CLI tool](#cascadia---css-selector-cli-tool)
- [Usage](#usage)
  - [$ cascadia](#-cascadia)
  - [Examples](#examples)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## cascadia - CSS selector CLI tool

The [Go Cascadia package](https://github.com/andybalholm/cascadia) implements CSS selectors for html. This is the command line tool, started as a thin wrapper around that package, but growing into a better tool to test CSS selectors without writing Go code:

## Usage

### $ cascadia
```sh
cascadia wrapper
Version 1.2.6 built on 2021-11-28
Copyright (C) 2021, Tong Sun

Command line interface to go cascadia CSS selectors package

Usage:
  cascadia -i in -c css -o [Options...]

Options:

  -h, --help        display help information 
  -i, --in         *The html/xml file to read from (or stdin) 
  -o, --out        *The output file (or stdout) 
  -c, --css        *CSS selectors (can provide more if not using --piece) 
  -t, --text        Text output for none-block selection mode 
  -R, --Raw         Raw text output, no trimming of leading and trailing white space 
  -p, --piece       sub CSS selectors within -css to split that block up into pieces
			format: PieceName=[RAW:]selector_string
			RAW: will return the selected as-is; else the text will be returned 
  -d, --delimiter   delimiter for pieces csv output [=	]
  -w, --wrap-html   wrap up the output with html tags 
  -y, --style       style component within the wrapped html head 
  -b, --base        base href tag used in the wrapped up html 
  -q, --quiet       be quiet
```

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
$ echo '<input type="radio" name="Sex" value="F" />' | tee /tmp/cascadia.xml | cascadia -i -o -c 'input[name=Sex][value=F]'
1 elements for 'input[name=Sex][value=F]':
<input type="radio" name="Sex" value="F"/>
```

Either the input or the output can be followed by a file name:


```sh
$ cascadia -i /tmp/cascadia.xml -o -c 'input[name=Sex][value=F]'
1 elements for 'input[name=Sex][value=F]':
<input type="radio" name="Sex" value="F"/>
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
    * [Twitter Search](https://github.com/suntong/cascadia/wiki#twitter-search)
  * [Reconstruct the separated pages](https://github.com/suntong/cascadia/wiki#reconstruct-the-separated-pages)
  * [More On CSS Selector](https://github.com/suntong/cascadia/wiki#more-on-css-selector)

## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/suntong/cascadia/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `cascadia_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf cascadia_*_linux_amd64.tar.gz
sudo mv -v cascadia_*_linux_amd64/cascadia /usr/local/bin/
rmdir -v cascadia_*_linux_amd64
```


### Distro package

- Packages available for Linux distros are
  * [Alpine Linux](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-alpine)
  * [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb)
  * [RedHat](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-rpm)

The repo setup instruction url has been given above.
For example, for [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb) --

### Debian package


```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/suntong/repo/setup.deb.sh' \
  | sudo -E bash

# That's it. You then can do your normal operations, like

sudo apt-get update
apt-cache policy cascadia

sudo apt-get install -y cascadia
```

## Install Source

To install the source code instead:

```
go get -v -u github.com/suntong/cascadia
```

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe)  
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
the _one-stop wire-framing solution_ for Go cli based projects, from _init_ to _deploy_.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/suntong"><img src="https://avatars.githubusercontent.com/u/422244?v=4?s=100" width="100px;" alt=""/><br /><sub><b>suntong</b></sub></a><br /><a href="https://github.com/go-cc/cc2py2/commits?author=suntong" title="Code">💻</a> <a href="#ideas-suntong" title="Ideas, Planning, & Feedback">🤔</a> <a href="#design-suntong" title="Design">🎨</a> <a href="#data-suntong" title="Data">🔣</a> <a href="https://github.com/go-cc/cc2py2/commits?author=suntong" title="Tests">⚠️</a> <a href="https://github.com/go-cc/cc2py2/issues?q=author%3Asuntong" title="Bug reports">🐛</a> <a href="https://github.com/go-cc/cc2py2/commits?author=suntong" title="Documentation">📖</a> <a href="#blog-suntong" title="Blogposts">📝</a> <a href="#example-suntong" title="Examples">💡</a> <a href="#tutorial-suntong" title="Tutorials">✅</a> <a href="#tool-suntong" title="Tools">🔧</a> <a href="#platform-suntong" title="Packaging/porting to new platform">📦</a> <a href="https://github.com/go-cc/cc2py2/pulls?q=is%3Apr+reviewed-by%3Asuntong" title="Reviewed Pull Requests">👀</a> <a href="#question-suntong" title="Answering Questions">💬</a> <a href="#maintenance-suntong" title="Maintenance">🚧</a> <a href="#infra-suntong" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
