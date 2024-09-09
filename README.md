# semversort

Sort a list of semver tags in the format `vX.Y.Z`. Useful for situations where lexicographic sorting is incorrect due to the absence of leading zeroes.

# prerequisites

1. [go](https://go.dev)

# installation

```
go install github.com/jabalsad/semversort
```

# usage

E.g., on a git repository with semver tags:

```
git tag | semversort
```

Or, in reverse:

```
git tag | semversort -r
```