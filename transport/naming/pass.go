package naming

type (
	pass struct{}

	passService struct {
		key string
	}
)

func NewPass(conf *Config) Naming {
	return &pass{}
}

func (p *pass) Endpoints() []string {
	return make([]string, 0)
}

func (p *pass) Name() string {
	return "pass"
}

func (p *pass) Register(key, val string) error {
	return nil
}

func (p *pass) Deregister(key string) error {
	return nil
}

func (p *pass) Service(key string) RegService {
	return &passService{key: key}
}

// ----------------- passService -----------------

func (ps *passService) Name() string {
	return ps.key
}

func (ps *passService) Addrs() ([]string, error) {
	return []string{ps.key}, nil
}

func (ps *passService) AddListener(f func()) {
}
