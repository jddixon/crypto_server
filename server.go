package crypto_server

// crypto_server/server.go

import (
	"crypto/rsa"
	xd "github.com/jddixon/xlOverlay_go/datakeyed"
	xh "github.com/jddixon/xlattice_go/httpd"
)

type CryptoServer struct {
	serverVersion string
	debugLog      string
	masterPubKey  *rsa.PublicKey
	BASE_DIR_NAME string // = "./"

	name2Hash    *xh.Name2Hash
	memCache     *xd.MemCache
	nodeDirName  string
	sitesDirName string
}
