/*
	gogger, is a Simple go logger with integration for sentry.
	Copyright (C) 2018 Asustin Studios <contact@asustin.net>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program. If not, see <http://www.gnu.org/licenses/>.

	For more information visit https://github.com/AsustinStudios/gogger
	or send an e-mail to contact@asustin.net
*/

package gogger // import "asustin.net/gogger"

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/getsentry/raven-go"
)

type Logger struct {
	logger *log.Logger
	Sentry *raven.Client
	tags   map[string]string
}

func New(out io.Writer, prefix string, flag int, sentryUrl string) *Logger {
	tags := map[string]string{
		"message":  "Authenticating",
		"category": "auth",
		"level":    "info",
	}

	client, err := raven.NewWithTags(sentryUrl, tags)
	if err != nil {
		fmt.Errorf("<%s>", err)
	}

	client.SetDefaultLoggerName("This is the logger name")
	client.SetEnvironment("This is the envirnoment")
	client.SetIncludePaths([]string{"log.go"})
	client.SetRelease("v1.5.69")

	return &Logger{log.New(out, prefix, flag), client, tags}
}

func (l *Logger) Debug(msg string, err error, environment map[string]string) {
	var formatted string
	env, err := json.MarshalIndent(environment, "", "    ")
	if err != nil {
		formatted = fmt.Sprintf("%#v", environment)
	} else {
		formatted = string(env)
	}
	l.logger.Printf("DEBUG: %s. %s\n%s\n", msg, err, formatted)
}

func (l *Logger) Info(msg string) {
	l.logger.Printf(" INFO: %s.\n", msg)
}

func (l *Logger) Warn(err error) {
	client.SetTagsContext(tags)
	client.SetUserContext(&raven.User{"1", "topo", "topo@asustin.net", "164.1.1.88"})
	ClearContext()

	l.Sentry.CaptureMessage(err, l.tags)
	l.logger.Printf(" WARN: %s. %s\n", err.Error())
}

func (l *Logger) Error(err error) {
	l.Sentry.CaptureError(err, l.tags)
	l.logger.Fatalf("ERROR: %s. %s\n", err.Error())
}

func (l *Logger) FatalError(err error) {
	l.Sentry.CaptureErrorAndWait(err, l.tags)
	l.logger.Fatalf("ERROR: %s. %s\n", err.Error())
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.logger.Fatalln(v)
}

func (l *Logger) Flags() int {
	return l.logger.Flags()
}

func (l *Logger) Output(calldepth int, s string) error {
	return l.logger.Output(calldepth, s)
}

func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panic(v)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v)
}

func (l *Logger) Panicln(v ...interface{}) {
	l.logger.Panicln(v)
}

func (l *Logger) Prefix() string {
	return l.logger.Prefix()
}

func (l *Logger) Print(v ...interface{}) {
	l.logger.Print(v)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v)
}

func (l *Logger) Println(v ...interface{}) {
	l.logger.Println(v)
}

func (l *Logger) SetFlags(flag int) {
	l.logger.SetFlags(flag)
}

func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

func (l *Logger) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}
