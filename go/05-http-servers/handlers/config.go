package handlers

import (
	"github.com/andycostintoma/http-servers/db/generated"
	"sync/atomic"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	Db             *generated.Queries
	Platform       string
	JwtSecret      string
}
