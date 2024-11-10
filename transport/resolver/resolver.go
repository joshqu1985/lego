package resolver

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"

	"lego/transport/naming"
)

func New(n naming.Naming) (resolver.Builder, error) {
	switch n.Name() {
	case "etcd":
		return newEtcd(n)
	}
	return nil, fmt.Errorf("unsupported naming: %s", n.Name())
}

func BuildTarget(n naming.Naming, target string) string {
	return fmt.Sprintf("%s://%s/%s", n.Name(),
		strings.Join(n.Endpoints(), ","), target)
}

type nopResolver struct {
	cc resolver.ClientConn
}

func (r *nopResolver) Close() {
}

func (r *nopResolver) ResolveNow(_ resolver.ResolveNowOptions) {
}
