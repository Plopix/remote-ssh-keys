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
package clients

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Client that knows how to get Keys for login
type Client struct{}

// GetKeys for an URL
func (client *Client) GetKeys(url string) []string {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		var empty []string
		return empty
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		var empty []string
		return empty
	}
	responseString := string(body)

	return strings.Split(responseString, "\n")
}
