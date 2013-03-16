/*
   This file will wrap the seelog APIs in the style of golang standard log,
   so that our existing codebase need only trival modification.

   General Usage:
     just replace your
         import log
     to
         import log "common/seelogWrapper"

   Comments:
     Generally speaking, you need not to change your style of calling log.*
     APIs, since seelogWrapper maintains the same interfaces of golang's log.

     In detail, as the default levels defined in seelog and golang are slightly
     different with each other. This wrapper map them as below:

     golang level                    seelog level
     --------------------------------------------
     Print         ===============>  Info
     Fatal         ===============>  Error
     Panic         ===============>  Critical

     Despite the level mapping,
       1. Semantics of golang's log remain the same. E.g., Fatal also means
     Print() then os.Exit().
       2. All seelog's APIs also be exposed. E.g., log.Tracef() is exposed even
     though there is no mapping in golang's log.

   Default Setting:
     The default logger is setup to write all levels of log to console in
     synchronized fashion. Its configuration can be changed with
     SetLoggerConfig function.
*/
package seelogWrapper

import (
  "fmt"
  log "seelog"
  "io"
  "os"
)

var seelogStaticFuncCallDepth int

// Init the default logger to console, and in sync mode

func init() {
  seelogStaticFuncCallDepth = log.GetStaticFuncCallDepth()
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 1)

  c := `<seelog type="sync">
          <outputs formatid="ccmp">
            <console />
          </outputs>
          <formats>
            <format id="ccmp" format="%Ns [%LEVEL] (%File:%Func) %Msg"/>
          </formats>
        </seelog>`
  SetLoggerConfig(c)
}

// belows are APIs originally provided by seelog

func UseLogger(logger log.LoggerInterface) error {
  return log.UseLogger(logger)
}

func ReplaceLogger(logger log.LoggerInterface) error {
  return log.ReplaceLogger(logger)
}

func Tracef(format string, params ...interface{}) {
  log.Tracef(format, params...)
}

func Debugf(format string, params ...interface{}) {
  log.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
  log.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) {
  log.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) {
  log.Errorf(format, params...)
}

func Criticalf(format string, params ...interface{}) {
  log.Criticalf(format, params...)
}

func Trace(v ...interface{}) {
  log.Trace(v...)
}

func Debug(v ...interface{}) {
  log.Debug(v...)
}

func Info(v ...interface{}) {
  log.Info(v...)
}

func Warn(v ...interface{}) {
  log.Warn(v...)
}

func Error(v ...interface{}) {
  log.Error(v...)
}

func Critical(v ...interface{}) {
  log.Error(v...)
}

func Flush() {
  log.Flush()
}

func LogLevelFromString(levelStr string) (level log.LogLevel, found bool) {
  return log.LogLevelFromString(levelStr)
}

func LoggerFromConfigAsBytes(data []byte) (log.LoggerInterface, error) {
  return log.LoggerFromConfigAsBytes(data)
}

func LoggerFromConfigAsFile(fileName string) (log.LoggerInterface, error) {
  return log.LoggerFromConfigAsFile(fileName)
}

func LoggerFromConfigAsString(data string) (log.LoggerInterface, error) {
  return log.LoggerFromConfigAsString(data)
}

func LoggerFromWriterWithMinLevel(output io.Writer,
  minLevel log.LogLevel) (log.LoggerInterface, error) {
  return log.LoggerFromWriterWithMinLevel(output, minLevel)
}

// belows are APIs needed by our codebase

// Fatal equals seelog.Error() then os.Exit(1)
func Fatal(v ...interface{}) {
  // +2 because Fatal -> Error -> seelog.Error, others are similar
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  log.Error(v...)
  os.Exit(1)
}

// Panic equals seelog.Critical() then panic()
func Panic(v ...interface{}) {
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  log.Critical(v...)
  panic("Panic in seelogWrapper, check last critical log for reason.!")
}

func Print(v ...interface{}) {
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  log.Info(v...)
}

func Println(v ...interface{}) {
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  log.Info(fmt.Sprintln(v...))
}

// Same side-effect as Fatal
func Fatalf(format string, v ...interface{}) {
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  s := fmt.Sprintf(format, v...)
  Fatal(s)
}

// Same side-effect as Panic
func Panicf(format string, v ...interface{}) {
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  s := fmt.Sprintf(format, v...)
  Panic(s)
}

func Printf(format string, v ...interface{}) {
  log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth + 2)
  defer log.SetStaticFuncCallDepth(seelogStaticFuncCallDepth)
  s := fmt.Sprintf(format, v...)
  Print(s)
}

// SetLoggerConfig sets logger with config.
//
// For details of how to write the seelog config,
// check https://github.com/cihub/seelog/wiki
func SetLoggerConfig(config string) {
  logger, _ := log.LoggerFromConfigAsBytes([]byte(config))
  err := log.ReplaceLogger(logger)
  if err != nil {
    Panicf("Can not replace default logger with new config: %v", config)
  }
}

