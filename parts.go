package golog

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"runtime"
)

//#####color part
func logPart_ColorBegin(log *Logger) {

	if log.enableColor && log.currColor != NoColor {

		log.WriteRawString(logColorPrefix[log.currColor])
	}
}

func logPart_ColorEnd(log *Logger) {

	if log.enableColor && log.currColor != NoColor {

		log.WriteRawString(logColorSuffix)
	}
}

//#####file part
func LogPart_ShortFileName(log *Logger) {

	writeFilePart(log, true, false)
}

func LogPart_LongFileName(log *Logger) {

	writeFilePart(log, false, true)
}

func writeFilePart(log *Logger, shortFile, longFile bool) {
	if shortFile || longFile {

		var file string
		var line int

		if shortFile || longFile {
			// release lock while getting caller info - it'text expensive.

			var ok bool
			_, file, line, ok = runtime.Caller(4)
			if !ok {
				file = "???"
				line = 0
			}
		}

		if shortFile {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		log.WriteRawString(file)
		log.WriteRawByte(':')
		itoa(log, line, -1)
		log.WriteRawString(": ")
	}
}

//#####level part
func LogPart_Level(log *Logger) {
	log.WriteRawString(levelString[log.currLevel])
	log.WriteRawByte(' ')

}

func LogPart_Name(log *Logger) {

	if log.name != "" {
		log.WriteRawString(log.name)
		log.WriteRawByte(' ')
	}
}

//#####pid part
var Goid int

func LogPart_Pid(log *Logger) {

	itoa(log, os.Getpid(), -1)
	log.WriteRawString(" ")
}

func LogPart_Gid(log *Logger) {
	
	itoa(log, getGoId(), -1)
	log.WriteRawString(" ")
}

func getGoId() int {
	if Goid != 0 {
		return Goid
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()
	
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id,_ := strconv.Atoi(idField)
	Goid = id
	return Goid
}

//#####text part
func logPart_Text(log *Logger) {

	log.WriteRawString(log.currText)
}

func logPart_Line(log *Logger) {

	l := len(log.currText)

	if (l > 0 && log.currText[l-1] != '\n') || l == 0 {
		log.WriteRawByte('\n')
	}

}

//#####time part
func LogPart_Time(log *Logger) {

	writeTimePart(log, false)
}

func LogPart_TimeMS(log *Logger) {

	writeTimePart(log, true)
}

func writeTimePart(log *Logger, ms bool) {

	now := time.Now() // get this early.

	year, month, day := now.Date()

	itoa(log, year, 4)
	log.WriteRawByte('/')
	itoa(log, int(month), 2)
	log.WriteRawByte('/')
	itoa(log, day, 2)
	log.WriteRawByte(' ')

	hour, min, sec := now.Clock()
	itoa(log, hour, 2)
	log.WriteRawByte(':')
	itoa(log, min, 2)
	log.WriteRawByte(':')
	itoa(log, sec, 2)

	if ms {
		log.WriteRawByte('.')
		itoa(log, now.Nanosecond()/1e3, 6)
	}

	log.WriteRawByte(' ')

}
