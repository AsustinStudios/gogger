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

type SentryUser raven.User

type Data struct {
	User *SentryUser
	Tags map[string]string
}

type Logger struct {
	Sentry      *raven.Client
	logger      *log.Logger
	debugLogger *log.Logger
	tags        map[string]string
}

func New(out io.Writer, prefix string, flag int, sentryUrl string) *Logger {
	client, err := raven.New(sentryUrl)
	if err != nil {
		log.Fatal(err)
	}
	client.SetDefaultLoggerName("cloud.asustin.net")
	return &Logger{
		Sentry:      client,
		logger:      log.New(out, prefix, flag),
		debugLogger: log.New(out, prefix, log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Debug(msg string, err error, environment map[string]string) {
	var e, formatted string
	if environment != nil {
		env, err := json.MarshalIndent(environment, "", "    ")
		if err != nil {
			formatted = fmt.Sprintf("%#v", environment)
		} else {
			formatted = string(env)
		}
	}
	if err != nil {
		e = err.Error()
	}
	l.debugLogger.Printf("DEBUG: %s. %s\n%s\n", msg, e, formatted)
}

func (l *Logger) Info(msg string) {
	l.logger.Printf("INFO: %s.\n", msg)
}

func (l *Logger) Warn(err error, data *Data) {
	if data != nil {
		if data.User != nil {
			user := raven.User(*data.User)
			l.Sentry.SetUserContext(&user)
		}
		if data.Tags != nil {
			l.Sentry.SetTagsContext(data.Tags)
		}
		defer l.Sentry.ClearContext()
	}
	l.Sentry.CaptureError(err, l.tags)
	l.logger.Printf("WARN: %s\n", err.Error())
}

func (l *Logger) Error(err error, data *Data) {
	if data != nil {
		if data.User != nil {
			user := raven.User(*data.User)
			l.Sentry.SetUserContext(&user)
		}
		if data.Tags != nil {
			l.Sentry.SetTagsContext(data.Tags)
		}
		defer l.Sentry.ClearContext()
	}
	l.Sentry.CaptureError(err, l.tags)
	l.logger.Fatalf("ERROR: %s\n", err.Error())
}

func (l *Logger) FatalError(err error, data *Data) {
	if data != nil {
		if data.User != nil {
			user := raven.User(*data.User)
			l.Sentry.SetUserContext(&user)
		}
		if data.Tags != nil {
			l.Sentry.SetTagsContext(data.Tags)
		}
		defer l.Sentry.ClearContext()
	}
	l.Sentry.CaptureErrorAndWait(err, l.tags)
	l.logger.Fatalf("ERROR: %s\n", err.Error())
}

func (l *Logger) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v)
	l.logger.Printf("INFO: %s", msg)
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
