package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/xlab/suplog"
)

// readEnv is a special utility that reads `.env` file into actual environment variables
// of the current app, similar to `dotenv` Node package.
func readEnv() {
	if envdata, _ := ioutil.ReadFile(".env"); len(envdata) > 0 {
		s := bufio.NewScanner(bytes.NewReader(envdata))
		for s.Scan() {
			parts := strings.Split(s.Text(), "=")
			if len(parts) != 2 {
				continue
			}
			strValue := strings.Trim(parts[1], `"`)
			if err := os.Setenv(parts[0], strValue); err != nil {
				log.WithField("name", parts[0]).WithError(err).Warningln("failed to override ENV variable")
			}
		}
	}
}
