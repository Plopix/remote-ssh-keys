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
package logger

import (
	"io"
	"io/ioutil"
	"log"
)

var (
	// PROD mode for logger
	PROD = true
	// Trace Log
	Trace *log.Logger
	// Info Log
	Info *log.Logger
	// Warning Log
	Warning *log.Logger
	// Error Log
	Error *log.Logger
)

// Init the Logger
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	if PROD {
		Trace.SetFlags(0)
		Trace.SetOutput(ioutil.Discard)
		Info.SetFlags(0)
		Info.SetOutput(ioutil.Discard)
		Warning.SetFlags(0)
		Warning.SetOutput(ioutil.Discard)
		Error.SetFlags(0)
		Error.SetOutput(ioutil.Discard)
	}
}
