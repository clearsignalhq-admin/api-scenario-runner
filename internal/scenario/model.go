package scenario

type Scenario struct {
	Name  string            `yaml:"scenario"`
	Vars  map[string]string `yaml:"vars"`
	Steps []Step            `yaml:"steps"`
	Rules []Rule            `yaml:"rules"`
}

type Step struct {
	Name    string   `yaml:"name"`
	Repeat  int      `yaml:"repeat"`
	Request *Request `yaml:"request"`
}

type Request struct {
	Method string                 `yaml:"method"`
	URL    string                 `yaml:"url"`
	Body   map[string]any         `yaml:"body"`
	Header map[string]string      `yaml:"headers"`
}

type Rule struct {
	Name   string         `yaml:"name"`
	If     string         `yaml:"if"`
	Expect map[string]any `yaml:"expect"`
}
