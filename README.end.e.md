
## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/{{.User}}/{{.Name}}/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `{{.Name}}_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf {{.Name}}_*_linux_amd64.tar.gz
sudo mv -v {{.Name}}_*_linux_amd64/{{.Name}} /usr/local/bin/
rmdir -v {{.Name}}_*_linux_amd64
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
apt-cache policy {{.Name}}

sudo apt-get install -y {{.Name}}
```

## Install Source

To install the source code instead:

```
go get -v -u {{.ProjectURL}}
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
