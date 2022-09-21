package utils

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

func sanitizeUrl(url string) string {
	url = regexp.MustCompile("^.*@").ReplaceAllLiteralString(url, "")
	url = regexp.MustCompile("^(ssh|http|https)?://").ReplaceAllLiteralString(url, "")
	url = strings.TrimSuffix(url, ".git")
	url = strings.ReplaceAll(url, ":", "/")
	return url
}

func GenerateScmId(cloneUrl string) string {
	sanitizedUrl := sanitizeUrl(cloneUrl)
	hash := md5.Sum([]byte(sanitizedUrl))
	return hex.EncodeToString(hash[:])
}
