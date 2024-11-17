package resolver

import (
	"path"

	"github.com/golang/glog"
	"google.golang.org/grpc/resolver"

	"github.com/joshqu1985/lego/transport/naming"
)

func newNacos(n naming.Naming) (resolver.Builder, error) {
	return &nacosBuilder{
		naming: n,
	}, nil
}

type nacosBuilder struct {
	naming naming.Naming
}

func (this *nacosBuilder) Scheme() string {
	return this.naming.Name()
}

func (this *nacosBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	_ resolver.BuildOptions) (resolver.Resolver, error) {

	service := this.naming.Service(path.Base(target.URL.Path))
	if _, err := service.Addrs(); err != nil {
		return nil, err
	}

	f := func() {
		values, err := service.Addrs()
		if err != nil {
			glog.Errorf("naming addrs err:%v", err)
			return
		}

		addrs := []resolver.Address{}
		for _, val := range values {
			addrs = append(addrs, resolver.Address{Addr: val})
		}

		if err := cc.UpdateState(resolver.State{
			Addresses: addrs,
		}); err != nil {
			glog.Errorf("update grpc ClientConn state err:%v", err)
		}
	}
	service.AddListener(f)
	f()

	return &nopResolver{cc: cc}, nil
}
