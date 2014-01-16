package crypto_server

// crypto_server/server.go

import (
	"crypto/rsa"
	xh "github.com/jddixon/xlattice_go/httpd"
	xd "github.com/jddixon/xlattice_go/overlay/datakeyed"
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
