package resolver

import (
	"path"

	"google.golang.org/grpc/resolver"

	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/transport/naming"
)

type nacosBuilder struct {
	naming naming.Naming
}

func newNacos(n naming.Naming) (resolver.Builder, error) {
	return &nacosBuilder{
		naming: n,
	}, nil
}

func (nb *nacosBuilder) Scheme() string {
	return nb.naming.Name()
}

func (nb *nacosBuilder) Build(target resolver.Target, cc resolver.ClientConn, //nolint:gocritic // target is heavy
	_ resolver.BuildOptions,
) (resolver.Resolver, error) {
	service := nb.naming.Service(path.Base(target.URL.Path))
	if _, err := service.Addrs(); err != nil {
		return nil, err
	}

	f := func() {
		values, err := service.Addrs()
		if err != nil {
			logs.Errorf("nacos naming addrs err:%v", err)

			return
		}

		addrs := make([]resolver.Address, 0)
		for _, val := range values {
			addrs = append(addrs, resolver.Address{Addr: val})
		}

		if xerr := cc.UpdateState(resolver.State{Addresses: addrs}); xerr != nil {
			logs.Errorf("nacos update grpc ClientConn state err:%v", xerr)
		}
	}

	service.AddListener(f)
	f()

	return &nopResolver{cc: cc}, nil
}
