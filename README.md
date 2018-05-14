
# cascadia

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/cascadia?status.svg)](http://godoc.org/github.com/suntong/cascadia)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/cascadia)](https://goreportcard.com/report/github.com/suntong/cascadia)
[![travis Status](https://travis-ci.org/suntong/cascadia.svg?branch=master)](https://travis-ci.org/suntong/cascadia)
[![Codeship Status](https://codeship.com/projects/7fbac590-a3dd-0134-4b89-26c19bdf8358/status?branch=master)](https://codeship.com/projects/190387)

## TOC
- [Usage](#usage)
  - [$ cascadia](#-cascadia)
- [Examples](#examples)
  - [None-block selection mode](#none-block-selection-mode)
    - [Multi-selection](#multi-selection)
  - [Block selection mode](#block-selection-mode)
    - [Twitter Search](#twitter-search)
  - [Reconstruct the separated pages](#reconstruct-the-separated-pages)
- [More On CSS Selector](#more-on-css-selector)

The [Go Cascadia package](https://github.com/andybalholm/cascadia) implements CSS selectors for html. This is the command line tool, started as a thin wrapper around that package, but growing into a better tool to test CSS selectors without writing Go code:

# Usage

#### $ cascadia
```sh
cascadia wrapper
Version 0.2.0 built on 2018-05-12

Command line interface to go cascadia CSS selectors package

Usage:
  cascadia -i in -c css -o [Options...]

Options:

  -h, --help            display help information
  -i, --in             *The html/xml file to read from (or stdin)
  -o, --out            *The output file (or stdout)
  -c, --css            *CSS selectors (can provide more if not using --piece)
  -p, --piece           sub CSS selectors within -css to split that block up into pieces
			format: PieceName=[RAW:]selector_string
			RAW: will return the selected as-is; else the text will be returned
  -d, --delimiter[=	]   delimiter for pieces csv output
  -w, --wrap-html       wrap up the output with html tags
  -b, --base            base href tag used in the wrapped up html
  -q, --quiet           be quiet
```

Its output has two modes, _none-block selection mode_ and _block selection mode_, depending on whether the `--piece` parameter is given on the command line or not.

- The none-block selection mode will output the selection as HTML source, while
- The block selection mode will output HTML text in a tsv/csv table form

For details about the concept of block and pieces, check out [andrew-d/goscrape](https://github.com/andrew-d/goscrape) (in fact, `cascadia` was initially developed just for it, so that I don't need to tweak Go code, build & run it just to test out the block and pieces selectors). Here is the exception:

- Inside each page, there's 1 or more *blocks* - some logical method of splitting up a page into subcomponents.
- Inside each block, you define some number of *pieces* of data that you wish
  to extract.  Each piece consists of a name, a selector, and what data to
  extract from the current block.

This all sounds rather complicated, but in practice it's quite simple. See the next section for details.


## Examples

### None-block selection mode

All the three `-i -o -c` options are required. By default it reads from `stdin` and output to `stdout`:

```sh
$ echo '<input type="radio" name="Sex" value="F" />' | tee /tmp/cascadia.xml | cascadia -i -o -c 'input[name=Sex][value=M]'
0 elements for 'input[name=Sex][value=M]':
```

Either the input or the output can be followed by a file name:


```sh
$ cascadia -i /tmp/cascadia.xml -o -c 'input[name=Sex][value=F]'
1 elements for 'input[name=Sex][value=F]':
<input type="radio" name="Sex" value="F"/>
```

#### Multi-selection

Of course, any number of selections are allowed (provided out of box from the CSS selection "`,`" syntax):

```sh
$ echo '<table border="0" cellpadding="0" cellspacing="0" style="table-layout: fixed; width: 100%; border: 0 dashed; border-color: #FFFFFF"><tr style="height:64px">aaa</tr></table>' | cascadia -i -o -c 'table[border="0"][cellpadding="0"][cellspacing="0"], tr[style=height\:64px]'
2 elements for 'table[border="0"][cellpadding="0"][cellspacing="0"], tr[style=height\:64px]':
<table border="0" cellpadding="0" cellspacing="0" style="table-layout: fixed; width: 100%; border: 0 dashed; border-color: #FFFFFF"><tbody><tr style="height:64px"></tr></tbody></table>
<tr style="height:64px"></tr>
```

Or, to make the multi-selection explicit on cli, emphasizing selecting being from different parts using different selectors, one can provide multiple `--css` on the command line. E.g.,

    cascadia -o -i http://www.iciba.com/conformity -c 'div.js-base-info > div > div > div.in-base-top.clearfix' -c 'div.js-base-info > div > div > ul' -c 'div.js-base-info > div > div > li' -c 'div.info-article.article-tab'

It'll construct the return from all four `-c` CSS selectors.

It has the same effect as using the "`,`" syntax, but

- The CSS selectors are provided explicitly with multiple `--css` parameters.
- The "`,`" syntax will return according to the order the selections occur in source, while
- The multiple `--css` will return according to the order the `--css` parameters.

### Block selection mode

First, as the none-block selection mode will output the selection as HTML _source_, so if you want HTML _text_ instead, then you can make use of the block selection mode. 

```sh
$ echo '<div class="container"><p align="justify"><b>Name: </b>John Doe</p></div>' | tee /tmp/cascadia.xml | cascadia -i -o -c 'div > p'
1 elements for 'div > p':
<p align="justify"><b>Name: </b>John Doe</p>

$ cat /tmp/cascadia.xml | cascadia -i -o -c 'div' --piece SelText='p'
SelText
Name: John Doe
```

Note that the block selection mode can output in HTML as well -- it just outputs (HTML) text by default:

```sh
$ cat /tmp/cascadia.xml | cascadia -i -o -c 'div' --piece SelText='RAW:p'
SelText 
<p align="justify"><b>Name: </b>John Doe</p>
```

The real power of _block selection mode_ resides in its capability of producing tsv/csv tables without any go programming:


```
$ curl --silent https://news.ycombinator.com | cascadia -i -o -c 'tr.athing' -p No=span.rank -p Title='td.title > a' -p Site=span.sitestr
No      Title   Site
1.      Onedrive is slow on Linux but fast with a ?Windows? user-agent (2016)   microsoft.com
2.      Starting today, users of Firefox can also enjoy Netflix on Linux        netflix.com
3.      Research Debt   distill.pub
...
27.     USPS Informed Delivery ? Digital Images of Front of Mailpieces  usps.com
28.     Performance bugs ? the dark matter of programming bugs  forwardscattering.org
29.     Most items of clothing have complicated international journeys  bbc.co.uk
30.     High-performance employees need quieter work spaces     qz.com
```

It's poor man's scrapper tool if text are the only thing needed. For scrapping beyond text, then just go one step further, to use  [andrew-d/goscrape](https://github.com/andrew-d/goscrape) (or my [goscrape](https://github.com/suntong/goscrape) instead, which has some enhancements to it).

Again, if text are the only thing needed, then `cascadia` might be already enough. Here is how to scrap Hacker News _across several pages_:

```
$ curl --silent https://news.ycombinator.com/news?p=[1-3] | cascadia -i -o -c 'tr.athing' -p No=span.rank -p Title='td.title > a' -p Site=span.sitestr
No      Title   Site
1.      Starting today, users of Firefox can also enjoy Netflix on Linux        netflix.com
2.      Onedrive is slow on Linux but fast with a ?Windows? user-agent (2016)   microsoft.com 
3.      Research Debt   distill.pub
...
27.     Yes I Still Want to Be Doing This at 56 (2012)  thecodist.com
28.     Performance bugs ? the dark matter of programming bugs  forwardscattering.org
29.     USPS Informed Delivery ? Digital Images of Front of Mailpieces  usps.com
30.     High-performance employees need quieter work spaces     qz.com
31.     Most items of clothing have complicated international journeys  bbc.co.uk
32.     Telstra?s Gigabit Class LTE Network     cellularinsights.com
...
58.     The New Laptop Ban Adds to Travelers' Lack of Privacy and Security      eff.org 
59.     QEMU: user-to-root privesc inside VM via bad translation caching        chromium.org
60.     Startups that debuted at Y Combinator W17 Demo Day 2    techcrunch.com
61.     The Cracking Monolith: Forces That Call for Microservices       semaphoreci.com 
62.     Amsterdam Airport Launches API Platform schiphol.nl
...
88.     Founder Stories: Leah Culver of Breaker (YC W17)        ycombinator.com 
89.     Find out what you, or someone on your team, did on the last working day github.com
90.     PSD2 ? a directive that will change banking in Europe   evry.com
```

By default it uses tab `\t` as fields delimiter, so the output is in `.tsv` format. To change to `.csv`, add `-d ,` to the command line.


#### Twitter Search

Block selection mode is poor man's web scrapping tool, and it is very simple to use. Here is another _practical_ example -- Twitter searching. We all know that you have to [pay for the Twitter Search API and it _only serves Tweets from the past week_](https://dev.twitter.com/rest/public/search). With `cascadia`, you can search the tweets for free, and get the latest content as well.

Here is how I watch for Toronto/GTA's Gas Price Alert, _without getting all other tweets_ from him:

```sh
$ cascadia -i 'https://twitter.com/search?q=%22Gas%20Price%20Alert%22%20%23GTA%20from%3AGasBuddyDan&src=typd' -o -c 'div.stream div.original-tweet div.content' --piece Time='small.time' --piece Tweet='div.js-tweet-text-container > p'
Time    Tweet

  Jul 31
        Gas Price Alert #Toronto #GTA #Hamilton #Ottawa #LdnOnt #Barrie #Kitchener #Niagara #Windsor N/C Tues and to a 2ct/l HIKE gor Wednesday

  Jul 6
        Gas Price Alert #Toronto #GTA #LdnOnt #Hamilton #Ottawa #Barrie #KW to see a 1 ct/l drop @ for Friday July 7

  May 30
        Gas Price Alert #Toronto #GTA #Ottawa #LdnOnt #Hamilton #KW #Barrie #Windsor prices won't change Wednesday but will DROP 1 ct/l Thursday

  May 15
        Gas Price Alert #Toronto #GTA #Barrie #Hamilton #LdnOnt #Ottawa #KW #Windsor NO CHANGE @  except gas bar shenanigans for Tues & Wednesday

  Mar 7
        Gas Price Alert #Toronto #GTHA #LdnOnt #Ottawa #Barrie #KW #Windsor to see a 1 cent a litre HIKE Wed March 8 (to 107.9 in the #GTA)

```


### Reconstruct the separated pages

Many web sites annoyingly separated one file into several small pieces so that they can show it to you in different web pages, with different ads. However, I'd like to view them in one page and no ads. Or, at least that is what I'd been hoping for all the time, but I didn't have an easy way of doing it until now, with `cascadia`.


With `cascadia` then no more programming is necessary. All we need to do now is to pass on some command line parameters, and the magic will happen. There are so many such sites that break thing into several small pieces, the following two are those I just did the other day.

The first one is separated across over 23 pages! Twenty-three! I would just give up if I don't have `cascadia`, but with it, it is so simple:

    curl --silent http://www.chinadmd.com/file/prrxtuivvxsxxwwaexuuwovp_[1-23].html | cascadia -i -o -c div.panel-body -p 'Book=div.tofu-txt' > /tmp/book.txt


The [first page is here](http://www.chinadmd.com/file/prrxtuivvxsxxwwaexuuwovp_1.html), and [all 23-pages are collected here](https://docs.google.com/document/d/1HkJ2oxvRSvoaNXl0n3t-uGhT5Dd08cvDbP9tB9Dmy8Q/preview). I collect them as plain text because the HTML were just wrapping around the plain text, thus no need HTML, plain text is good enough.

Collecting as HTML is no trouble either. Here is another example:

     curl --silent http://www.shangxueedu.com/shuxue/ksdg/20170113_162_[1-6].html | cascadia -i -o -c div.m-post -p 'Book=RAW:div.post-con' --wrap-html | tee /tmp/book.html

The [fifth page is here](http://www.shangxueedu.com/shuxue/ksdg/20170113_162_5.html), and [all pages are collected here](https://docs.google.com/document/d/1StFwP7kChHiGsL-hm3tnY29bsBRQWCU7xdhu2shsGcg/preview). Please check them out.

# More On CSS Selector

I'm not an expert on CSS Selector at all, but the following resources are what I found most helpful to me.

- [CSS Selectors Cheat Sheet](http://butlerccwebdev.net/support/css-selectors-cheatsheet.html) I think It's very good, because it's usage oriented and very practical, i.e., it arranges the Selectors according to their purposes. If that's too dry for you, check out
- [The 30 CSS Selectors You Must Memorize](http://code.tutsplus.com/tutorials/the-30-css-selectors-you-must-memorize--net-16048) It only lists those selectors that are important, but it gives concrete examples and explanations
- [CSS Selector Reference](http://www.w3schools.com/cssref/css_selectors.asp) from w3schools. This is the one I most often refer to, because the list is comprehensive, and there is also an online [CSS Selector Tester](http://www.w3schools.com/cssref/trysel.asp) that really helped me learn and understand
