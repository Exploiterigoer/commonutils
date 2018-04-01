// Package iniutil the package for analyzing configuration file with the ini suffix
package iniutil

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Config struct
type Config struct {
	filepath string                         //your ini file path directory+file
	conflist []map[string]map[string]string //configuration information slice
}

//SetConfig Creates an empty configuration file
func SetConfig(filepath string) *Config {
	c := new(Config)
	c.filepath = filepath

	return c
}

//GetValue To obtain corresponding value of the key values
func (c *Config) GetValue(section, name string) string {
	c.ReadList()
	conf := c.ReadList()
	for _, v := range conf {
		for key, value := range v {
			if key == section {
				return value[name]
			}
		}
	}
	return ""
}

//SetValue Set the corresponding value of the key value, if not add, if there is a key change
func (c *Config) SetValue(section, key, value string) bool {
	c.ReadList()
	data := c.conflist
	var ok bool
	var index = make(map[int]bool)
	var conf = make(map[string]map[string]string)
	for i, v := range data {
		_, ok = v[section]
		index[i] = ok
	}

	i, ok := func(m map[int]bool) (i int, v bool) {
		for i, v := range m {
			if v == true {
				return i, true
			}
		}
		return 0, false
	}(index)

	if ok {
		c.conflist[i][section][key] = value
		return true
	}

	conf[section] = make(map[string]string)
	conf[section][key] = value
	c.conflist = append(c.conflist, conf)
	return true
}

//DeleteValue Delete the corresponding key values
func (c *Config) DeleteValue(section, name string) bool {
	c.ReadList()
	data := c.conflist
	for i, v := range data {
		for key := range v {
			if key == section {
				delete(c.conflist[i][key], name)
				return true
			}
		}
	}
	return false
}

//ReadList List all the configuration file
func (c *Config) ReadList() []map[string]map[string]string {
	file, err := os.Open(c.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string

	var buf *bytes.Buffer
	x := make([]byte, 4096)
	n, _ := file.Read(x)

	buf = bytes.NewBuffer(x[:n])
	if x[0] == 239 {
		// If the configuration file is with utf8-bom format
		// it needs to cut the bom
		buf = bytes.NewBuffer(x[3:n])
	}

	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)

		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == ";": //doing nothing when a line staring with ";"
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			data[section][strings.TrimSpace(line[0:i])] = value
			if c.uniquappend(section) == true {
				c.conflist = append(c.conflist, data)
			}
		}
	}

	return c.conflist
}

// CheckErr dos something for the error
func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}

//uniquappend Ban repeated appended to the slice method
func (c *Config) uniquappend(conf string) bool {
	for _, v := range c.conflist {
		for k := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
