//+build wireinject

package line

import (
	"github.com/google/wire"
)

func InitializeApi() ApiInterface {
	wire.Build(
		NewApi,
	)
	return nil
}
