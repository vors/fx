# Developing UberFx

This doc is intended for contributors to UberFx (hopefully that's you!)

## Development Environment

* Go. Install on OS X with `brew install go`. Make sure `go version` returns at
  least `1.7` since we're going to be using 1.7+ features like subtests.

* [Overcommit](https://github.com/brigade/overcommit), a git hook manager.
  Install `overcommit` into your path with `sudo gem install overcommit`.
  Enable it on this repo with `overcommit --install`.
  We use Overcommit to enforce a variety of style, semantic, and legal things
  (e.g. licence headers on all source files).

## Checking out the code

Make sure the repository is cloned to the correct location:

```bash
go get go.uber.org/fx/...
cd $GOPATH/src/go.uber.org/fx
```

## Dependency management

Dependencies are tracked via `glide.yaml`. If you're not familiar with `glide`,
read the [docs](https://github.com/Masterminds/glide#usage).

## Licence headers

This project is Open Source Software, and requires a header at the beginning of
all source files. This is enforced by commit hooks and TravisCI.

To add licence headers, use
[uber-licence](https://github.com/uber/uber-licence):

```lang=bash
make add-uber-licence
```

## Commit Messages

Overcommit adds some requirements to your commit messages. At Uber, we follow the
[Chris Beams](http://chris.beams.io/posts/git-commit/) guide to writing git
commit messages. Read it, follow it, learn it, love it.

## FIXMEs

If you ever are in the middle of writing code and remember a thing you need to
do, leave a comment like:

```go
// FIXME(ai) make this actually work
```

Your initials in the parens are optional but a good practice. This is better
than a TODO because our CI checks for unresolved FIXMEs, so if you forget to fix
it, your code won't get merged.

## Testing

Run all the tests with coverage and race detector enabled:

```bash
make test RACE=-race
```

### Disabling the race detector

The race detector makes the tests run way slower. To disable it:

```bash
make test
```

TODO(ai) come up with something better for this.

### Viewing HTML coverage

```bash
make coverage.html && open coverage.html
```

You'll need to have [gocov-html](https://github.com/matm/gocov-html) installed:

```bash
go get -u gopkg.in/matm/v1/gocov-html
```
