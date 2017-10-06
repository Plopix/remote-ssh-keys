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

import "fmt"

// Github for Github API
type Github struct {
	Client
}

// GetKeysForLogin for a login
func (client *Github) GetKeysForLogin(login string) []string {
	var url = fmt.Sprintf("https://github.com/%s.keys", login)
	return client.GetKeys(url)
}
