NAME    := $(shell basename $(shell pwd))
VERSION ?= 0.0.0

all: test

clean: cleanbuild cleanpkg
cleanbuild:
	rm -fr ./build
cleanpkg:
	rm -fr ./pkg

show:
	echo "Building '${NAME}'"

test:
	go -v test ./...

build: build/${NAME}
build/%:
	GOOS=linux GOARCH=amd64 go build -v -ldflags="-X main.date=`date -I`" -o build/$*

pkg: show build pkg/${NAME}.deb
pkg/%.deb:
	mkdir -p ./pkg
# 	https://github.com/jordansissel/fpm/wiki
	fpm --verbose -s dir -t deb \
		--name ${NAME} \
		--package ./pkg/${NAME}.deb \
		--force \
		--deb-compression bzip2 \
		--url "${PKG_URL}" \
		--category ${PKG_CAT} \
		--description "${PKG_DESC}" \
		--maintainer "${PKG_MAINT}" \
		--vendor "${PKG_VEND}" \
		--license "${PKG_LICNS}" \
		--version ${VERSION} \
		--architecture ${PKG_ARCH} \
		--depends apt \
		./build/=/usr/bin/
