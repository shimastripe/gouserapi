package middleware

import (
	"errors"
	"strings"

	"github.com/blang/semver"
	"github.com/gin-gonic/gin"
)

func VersionInit(c *gin.Context) (semver.Version, error) {
	header := c.Request.Header["Accept"][0]
	header = strings.Join(strings.Fields(header), "")
	var ver semver.Version
	var err error

	// check accept-type
	if !strings.Contains(header, "application/vnd.wantedly+json") {
		return ver, errors.New("Incorrect Accept-type!")
	}

	// header version
	if strings.Contains(header, "version=") {
		h := strings.Split(header, "=")[1]
		h = setVersion(h)
		ver, err = semver.Parse(h)
		if err != nil {
			return ver, err
		}
	}

	// query v
	v := c.Query("v")
	if v != "" {
		v = setVersion(v)
		ver, err = semver.Parse(v)
		if err != nil {
			return ver, err
		}
	}
	ver_range := semver.MustParseRange("0.0.0")
	if ver_range(ver) {
		return ver, errors.New("specify the version!")
	}
	return ver, nil
}

func setVersion(input string) string {
	switch len(strings.Split(input, ".")) {
	case 2:
		input = input + ".0"
		return input
	case 1:
		input = input + ".0.0"
		return input
	}
	return input
}
