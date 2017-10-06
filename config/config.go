/**
 * This file is part of the plopix/remote-ssh-keys package.
 *
 * @author    Plopix - Sébastien Morel <morel.seb@gmail.com>
 * @copyright 2017 Sébastien Morel
 * @license   MIT
 *
 * For the full copyright and license information(MIT), please view the LICENSE
 * file that was distributed with this source code
 */
package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Providers per account
type Providers struct {
	Github []string `yaml:"github"`
	Plain  []string `yaml:"plain"`
	URL    []string `yaml:"url"`
}

// Configuration project
type Configuration struct {
	Account map[string]Providers `yaml:"accounts"`
}

// Load the configuration
func Load(filename string) Configuration {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		panic(err)
	}

	configurationContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var yamlConfig Configuration
	err = yaml.Unmarshal([]byte(configurationContent), &yamlConfig)
	if err != nil {
		panic(err)
	}
	return yamlConfig
}
