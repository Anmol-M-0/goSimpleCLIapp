package scanners

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

//scan for a float64 from std input (console)
func ScanFloat() (f float64) {
	fmt.Scanf("%f", &f)
	sc.Scan()
	return f
}

//scan for a int from std input (console)
func ScanInt() (i int) {
	fmt.Scanf("%d", &i)
	sc.Scan()
	return i
}

// scan for a string from std input (console)
// it will by default trim all the empty spaces at the beginning and end of the
// input text
func ScanString() (str string) {
	sc.Scan()
	str = sc.Text()
	str = strings.TrimSpace(str)
	return str
}

//scan for residual text in line, equivalent to flushing the buffer
func Flush() {
	sc.Scan()
}

// returns flushed text. Call it after calling any of the following:
// Flush(), Scanf(), ScanString(), ScanInt(), ScanFloat(), or any function
// in this package that claims to clear the buffer
func FlushValue() string {
	return sc.Text()
}

// Scanf scans text read from standard input, storing successive
// space-separated values into successive arguments as determined by
// the format. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
// Newlines in the input must match newlines in the format.
// The one exception: the verb %c always scans the next rune in the
// input, even if it is a space (or tab etc.) or newline.
// This function additionally flushes the buffer so you do not have to deal with stray input
//
func Scanf(format string, a ...interface{}) (n int, err error) {
	// return Fscanf(os.Stdin, format, a...)
	n, err = fmt.Scanf(format, a...)
	Flush()
	return n, err
}

// scans the whole line and returns it as string
// if error, returns first non EOF error
func ScanLine() (string, error) {
	sc.Scan()
	return sc.Text(), sc.Err()
}
