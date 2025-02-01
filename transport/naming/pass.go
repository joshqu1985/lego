package naming

func NewPass(conf *Config) Naming {
	return &pass{}
}

type pass struct {
}

func (this *pass) Endpoints() []string {
	return []string{}
}

func (this *pass) Name() string {
	return "pass"
}

func (this *pass) Register(key, val string) error {
	return nil
}
func (this *pass) Deregister(key string) error {
	return nil
}

func (this *pass) Service(key string) RegService {
	return &passService{key: key}
}

type passService struct {
	key string
}

func (this *passService) Name() string {
	return this.key
}

func (this *passService) Addrs() ([]string, error) {
	return []string{this.key}, nil
}

func (this *passService) AddListener(f func()) {
}
