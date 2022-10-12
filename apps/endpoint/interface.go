package endpoint

const (
	AppName = "endpoint"
)

type Service interface {
	RPCServer
}

func (s *EndpiontSet) ToDocs() (docs []interface{}) {
	for i := range s.Endpoints {
		docs = append(docs, s.Endpoints[i])
	}
	return
}

func NewRegistryResponse() *RegistryResponse {
	return &RegistryResponse{}
}
