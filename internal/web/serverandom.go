package web

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const (
	// total random bytes per write to the output
	CHUNKSZ = 128
	// max bytes, should be a configurable prop
	MAXSZ = 100 * 1024 * 1024 * 1024
)

func genRandomFile(c *gin.Context) {

	dataSz := c.Param("datasz")

	dataSzInt, err := strconv.ParseUint(dataSz, 10, 64)

	if err != nil {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(http.StatusBadRequest, "Unable to parse data size (%s)", dataSz)
		return
	} else if dataSzInt > MAXSZ {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(http.StatusForbidden, "Rejected request for %d bytes, max size is %d", dataSzInt, MAXSZ)
		return
	}

	bufSz := CHUNKSZ
	buf := make([]byte, bufSz)
	counter := uint64(0)
	readCount := 0
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/octet-string")
	for counter < dataSzInt {
		readCount, err = rand.Read(buf)
		if err != nil {
			errStr := fmt.Sprintf("error while generating random string: %s", err)
			log.Error().Msg(errStr)
			c.String(http.StatusInternalServerError, errStr)
			return
		}
		if readCount == bufSz {
			if counter+uint64(bufSz) < dataSzInt {
				c.Writer.Write(buf)
			} else {
				c.Writer.Write(buf[:dataSzInt-counter])
			}
			counter += uint64(readCount)

		} else {
			c.String(http.StatusInternalServerError, "entropy exhausted")
			return
		}
	}
	//c.String(http.StatusOK, fmt.Sprintf("size=%d bytes", dataSzInt))
}

func setupPingPong(c *gin.Context) {

	// Ping test
	c.String(http.StatusOK, "pong")

}
