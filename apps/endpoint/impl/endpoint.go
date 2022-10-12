package impl

import (
	"context"
	"github.com/zginkgo/ginkgo_keyauth/apps/endpoint"
)

func (s *service) RegistryEndpoint(ctx context.Context, req *endpoint.EndpiontSet) (
	*endpoint.RegistryResponse, error) {
	if err := s.save(ctx, req); err != nil {
		return nil, err
	}
	resp := endpoint.NewRegistryResponse()
	return resp, nil
}
