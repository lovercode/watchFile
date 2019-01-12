package conf

type Conf struct {
	Files map[string]map[string][]string
}

var Config Conf
