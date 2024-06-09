package cors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ApplyCORSHeaders(ctx context.Context) {

	// CORS Headers
	// https://stackoverflow-com.translate.goog/questions/31379995/grpc-header-cookie-in-go?_x_tr_sl=en&_x_tr_tl=ru&_x_tr_hl=ru&_x_tr_pto=sc

	header := metadata.Pairs(
		"Access-Control-Allow-Origin", "*",
	)
	grpc.SendHeader(ctx, header)
}
