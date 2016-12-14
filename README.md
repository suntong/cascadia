# cascadia

[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/cascadia)](https://goreportcard.com/report/github.com/suntong/cascadia)
[ ![Codeship Status](https://codeship.com/projects/7fbac590-a3dd-0134-4b89-26c19bdf8358/status?branch=master)](https://codeship.com/projects/190387)

The [Go Cascadia package](https://github.com/andybalholm/cascadia) implements CSS selectors for html. This is a thin wrapper around that package, a command line tool, to test CSS selectors without writing Go code:

```
$ cascadia 
cascadia wrapper
built on 2016-12-10

Command line interface to go cascadia CSS selectors package

Options:

  -h, --help   display help information
  -i, --in    *The html/xml file to read from (or stdin)
  -o, --out   *The output file (or stdout)
  -c, --css   *CSS selectors
```

## Examples

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

Of course, any number of selections allowed:

```sh
$ echo '<table border="0" cellpadding="0" cellspacing="0" style="table-layout: fixed; width: 100%; border: 0 dashed; border-color: #FFFFFF"><tr style="height:64px">aaa</tr></table>' | cascadia -i -o -c 'table[border="0"][cellpadding="0"][cellspacing="0"], tr[style=height\:64px]'
2 elements for 'table[border="0"][cellpadding="0"][cellspacing="0"], tr[style=height\:64px]':
<table border="0" cellpadding="0" cellspacing="0" style="table-layout: fixed; width: 100%; border: 0 dashed; border-color: #FFFFFF"><tbody><tr style="height:64px"></tr></tbody></table>
<tr style="height:64px"></tr>
```
