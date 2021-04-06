package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Tim0401/line-new-graduate/pkg/line"
)

const location = "Asia/Tokyo"

func init() {
	// timezone設定
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	// 乱数シード
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func main() {
	// アプリ実行
	apiInterface := line.InitializeApi()
	apiInterface.Serve()
}
