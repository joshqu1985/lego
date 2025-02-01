package resolver

import (
	"path"
	"strings"

	"google.golang.org/grpc/resolver"

	"github.com/joshqu1985/lego/transport/naming"
)

func newPass(n naming.Naming) (resolver.Builder, error) {
	return &passBuilder{
		naming: n,
	}, nil
}

type passBuilder struct {
	naming naming.Naming
}

func (this *passBuilder) Scheme() string {
	return this.naming.Name()
}

func (this *passBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	_ resolver.BuildOptions) (resolver.Resolver, error) {
	endpoints := strings.FieldsFunc(path.Base(target.URL.Path), func(r rune) bool {
		return r == ','
	})

	addrs := make([]resolver.Address, 0, len(endpoints))
	for _, val := range endpoints {
		addrs = append(addrs, resolver.Address{Addr: val})
	}

	if err := cc.UpdateState(resolver.State{
		Addresses: addrs,
	}); err != nil {
		return nil, err
	}
	return &nopResolver{cc: cc}, nil
}
