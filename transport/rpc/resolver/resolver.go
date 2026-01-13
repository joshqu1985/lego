package resolver

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"

	"github.com/joshqu1985/lego/transport/naming"
)

type nopResolver struct {
	cc resolver.ClientConn
}

func New(n naming.Naming) (resolver.Builder, error) {
	switch n.Name() {
	case "etcd":
		return newEtcd(n)
	case "nacos":
		return newNacos(n)
	case "pass":
		return newPass(n)
	}

	return nil, fmt.Errorf("unsupported naming: %s", n.Name())
}

func BuildTarget(n naming.Naming, target string) string {
	return fmt.Sprintf("%s://%s/%s", n.Name(),
		strings.Join(n.Endpoints(), ","), target)
}

func (r *nopResolver) Close() {
}

func (r *nopResolver) ResolveNow(_ resolver.ResolveNowOptions) {
}
