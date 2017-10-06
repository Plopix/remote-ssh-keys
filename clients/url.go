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
	"fmt"

	"github.com/plopix/remote-ssh-keys/logger"
)

// URL for key
type URL struct {
	Client
}

// GetKeysForLogin for a login
func (client *URL) GetKeysForLogin(urlText, login string) []string {
	var url = fmt.Sprintf(urlText, login)
	logger.Trace.Println(url)
	return client.GetKeys(url)
}
