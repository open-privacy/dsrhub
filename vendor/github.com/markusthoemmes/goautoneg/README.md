# goautoneg

[![Go Report Card](https://goreportcard.com/badge/github.com/markusthoemmes/goautoneg)](https://goreportcard.com/report/github.com/markusthoemmes/goautoneg)
[![CircleCI](https://circleci.com/gh/markusthoemmes/goautoneg.svg?style=svg)](https://circleci.com/gh/markusthoemmes/goautoneg)

This is a complete rewrite of https://bitbucket.org/ww/goautoneg and aims to be a drop-in replacement of that module. The aforementioned repository seems unmaintained with the last commit having been done on **2012-07.07**.

The `Negotiate` function is not implemented currently because a quick search resulted in the `ParseAccept` function being used in the majority of cases.

Ultimately, we should end up with this implementation being internalized in Go's standard library. The [proposal for that has already been accepted](https://github.com/golang/go/issues/19307) and it's a matter of time for it to land. This is mostly a stop-gap until that happens.

## Why do this?

That's a very fair question!

1. The original version on bitbucket.org is hosted as a Mercurial repository. Nothing's wrong with that but `dep` has it's [issues with that](https://github.com/golang/dep/issues/1692) if `hg` (the Mercurial client) is not installed locally. Several high-profile repositories (like Kubernetes' API server) depend on it though, which makes it a pain to constantly point out to new contributors to install `hg`, as it seems so unnecessary. This is especially true if the dependency is pulled in transitively, so you might even be aware of the fact that you now have a Mercurial dependency in your project. An [inquiry to mirror the Mercurial repository to a git repository](https://bitbucket.org/ww/goautoneg/issues/3/mirror-as-a-git-repository) remains unanswered.
2. The original repository has no `LICENSE` file. The license is part of `README.txt` though so an automated license checker can be made happy with some `sed` magic. It ain't pretty though.
3. Bugs go unanswered.