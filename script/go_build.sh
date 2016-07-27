#!/bin/bash
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  glide install
  go build -v -o ${distdir}/tugbot-parse
}

go_build
