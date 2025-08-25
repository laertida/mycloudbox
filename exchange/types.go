package exchange

type BucketConfig struct {
	Bucket Bucket `yaml:"bucket"`
}

type Bucket struct {
	Name   string `yaml:"name"`
	Region string `yaml:"region"`
}

type Endpoint struct {
	Schema     string
	Path       string
	Properties []Property
	Exchange   Exchange
}

type Message struct {
	Headers     []Header
	Attachments map[string]string
	Body        os.File
}

type Header struct {
	Name  string
	Value string
}

type Property struct {
	Key   string
	Value string
}

type Exchange struct {
	in         Message
	out        Message
	Properties map[string]string
}

func NewEndpoit(
	Schema string,
	Path string,
	Properties []Property,
	Exchange Exchange) Endpoint {

	return Endpoint{
		Schema:     Schema,
		Path:       Path,
		Properties: Properties,
		Exchange:   Exchange,
	}
}
