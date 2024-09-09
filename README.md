# semversort

Sort a list of semver tags in the format `vX.Y.Z` (the leading `v` is optional). Useful for situations where lexicographic sorting is incorrect due to the absence of leading zeroes.

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

Ignore lines that don't match the semver format:

```
git tag | semversort -i
```

# example

```
echo '
v0.0.10
v0.0.9
' | semversort
```

Should output

```
v0.0.9
v0.0.10
```
