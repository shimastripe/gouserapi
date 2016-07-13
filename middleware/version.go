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
	ver := semver.MustParse("0.0.0")

	// check accept-type
	if !strings.Contains(header, "application/vnd.wantedly+json") {
		return ver, errors.New("Incorrect Accept-type!")
	}

	// header version
	if strings.Contains(header, "version=") {
		h := strings.Split(header, "=")[1]
		switch len(strings.Split(h, ".")) {
		case 2:
			h = h + ".0"
		case 1:
			h = h + ".0.0"
		}
		var err error
		ver, err = semver.Parse(h)
		if err != nil {
			return ver, err
		}
	}
	return ver, nil
}
