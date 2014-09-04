package servant

import (
	"strconv"
	"strings"
)

func ToEng(v string) string {
	switch {
	case "黒" == v:
		return "black"
	case "青" == v:
		return "blue"
	case "無" == v:
		return "colorless"
	case "赤" == v:
		return "red"
	case "緑" == v:
		return "green"
	case "白" == v:
		return "white"
	}
	return ""
}

func ToInt(v string) int {
	if v == "-" {
		return -1
	}
	i, _ := strconv.Atoi(v)
	return i
}

func TrimHyphen(v string) string {
	if v == "-" {
		return ""
	}
	return v
}

func TrimLinefeed(v string) string {
	return strings.TrimRight(strings.TrimLeft(strings.TrimSpace(v), "\n"), "\n")
}

func Contains(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}
