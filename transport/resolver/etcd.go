package resolver

import (
	"log"
	"strings"

	"google.golang.org/grpc/resolver"

	"lego/transport/naming"
)

func newEtcd(n naming.Naming) (resolver.Builder, error) {
	return &etcdBuilder{
		naming: n,
	}, nil
}

func (this *etcdBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	_ resolver.BuildOptions) (resolver.Resolver, error) {

	service := this.naming.Service(strings.Trim(target.URL.Path, "/"))
	_, err := service.Addrs()
	if err != nil {
		return nil, err
	}

	f := func() {
		values, err := service.Addrs()
		if err != nil {
			log.Println("naming addrs err:", err)
			return
		}

		addrs := []resolver.Address{}
		for _, val := range values {
			addrs = append(addrs, resolver.Address{Addr: val})
		}

		if err := cc.UpdateState(resolver.State{
			Addresses: addrs,
		}); err != nil {
			log.Println("update grpc ClientConn state err:", err)
		}
	}
	service.AddListener(f)
	f()

	return &nopResolver{cc: cc}, nil
}

func (this *etcdBuilder) Scheme() string {
	return this.naming.Name()
}

type etcdBuilder struct {
	naming naming.Naming
}
