package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mdata := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mdata.UserAgent = userAgents[0]
		}
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mdata.UserAgent = userAgents[0]
		}
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mdata.ClientIP = clientIPs[0]
		}
	}
	if p, ok := peer.FromContext(ctx); ok {
		mdata.ClientIP = p.Addr.String()
	}

	return mdata
}
