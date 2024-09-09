package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type semver struct {
	Major int
	Minor int
	Patch int
	Orig  string
}

// Implement sort.Interface for []semver
type semverSlice []semver

func (s semverSlice) Len() int      { return len(s) }
func (s semverSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s semverSlice) Less(i, j int) bool { // Define how to compare semver structs
	if s[i].Major != s[j].Major {
		return s[i].Major < s[j].Major
	}
	if s[i].Minor != s[j].Minor {
		return s[i].Minor < s[j].Minor
	}
	return s[i].Patch < s[j].Patch
}

func parseVersion(version string) (semver, error) {
	trimmed := strings.TrimPrefix(version, "v")
	parts := strings.Split(trimmed, ".")
	if len(parts) != 3 {
		return semver{}, fmt.Errorf("invalid version format")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return semver{}, err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return semver{}, err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return semver{}, err
	}

	return semver{Major: major, Minor: minor, Patch: patch, Orig: version}, nil
}

func main() {
	reverse := flag.Bool("r", false, "sort in reverse order")
	ignoreInvalid := flag.Bool("i", false, "ignore lines that do not parse to a semver")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var versions semverSlice

	for scanner.Scan() {
		v, err := parseVersion(scanner.Text())
		if err != nil {
			if !*ignoreInvalid {
				fmt.Fprintf(os.Stderr, "Error parsing version '%s': %v\n", scanner.Text(), err)
			}
			continue
		}
		versions = append(versions, v)
	}

	if *reverse {
		sort.Sort(sort.Reverse(versions))
	} else {
		sort.Sort(versions)
	}

	for _, v := range versions {
		fmt.Println(v.Orig)
	}
}
