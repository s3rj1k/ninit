package signals

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

// https://golang.org/pkg/syscall/

// Parse matches signal name (number) to internal type.
func Parse(sig string) (syscall.Signal, error) {
	switch strings.ToUpper(strings.TrimSpace(sig)) {
	case "1", "SIGHUP", "HUP":
		return syscall.SIGHUP, nil
	case "2", "SIGINT", "INT":
		return syscall.SIGINT, nil
	case "3", "SIGQUIT", "QUIT":
		return syscall.SIGQUIT, nil
	case "4", "SIGILL", "ILL":
		return syscall.SIGILL, nil
	case "5", "SIGTRAP", "TRAP":
		return syscall.SIGTRAP, nil
	case "6", "SIGABRT", "SIGIOT", "ABRT", "IOT":
		return syscall.SIGABRT, nil
	case "7", "SIGBUS", "BUS":
		return syscall.SIGBUS, nil
	case "8", "SIGFPE", "FPE":
		return syscall.SIGFPE, nil
	case "9", "SIGKILL", "KILL":
		return syscall.SIGKILL, nil
	case "10", "SIGUSR1", "USR1":
		return syscall.SIGUSR1, nil
	case "11", "SIGSEGV", "SEGV":
		return syscall.SIGSEGV, nil
	case "12", "SIGUSR2", "USR2":
		return syscall.SIGUSR2, nil
	case "13", "SIGPIPE", "PIPE":
		return syscall.SIGPIPE, nil
	case "14", "SIGALRM", "ALRM":
		return syscall.SIGALRM, nil
	case "15", "SIGTERM", "TERM":
		return syscall.SIGTERM, nil
	case "16", "SIGSTKFLT", "STKFLT":
		return syscall.SIGSTKFLT, nil
	case "17", "SIGCHLD", "SIGCLD", "CHLD", "CLD":
		return syscall.SIGCHLD, nil
	case "18", "SIGCONT", "CONT":
		return syscall.SIGCONT, nil
	case "19", "SIGSTOP", "STOP":
		return syscall.SIGSTOP, nil
	case "20", "SIGTSTP", "TSTP":
		return syscall.SIGTSTP, nil
	case "21", "SIGTTIN", "TTIN":
		return syscall.SIGTTIN, nil
	case "22", "SIGTTOU", "TTOU":
		return syscall.SIGTTOU, nil
	case "23", "SIGURG", "URG":
		return syscall.SIGURG, nil
	case "24", "SIGXCPU", "XCPU":
		return syscall.SIGXCPU, nil
	case "25", "SIGXFSZ", "XFSZ":
		return syscall.SIGXFSZ, nil
	case "26", "SIGVTALRM", "VTALRM":
		return syscall.SIGVTALRM, nil
	case "27", "SIGPROF", "PROF":
		return syscall.SIGPROF, nil
	case "28", "SIGWINCH", "WINCH":
		return syscall.SIGWINCH, nil
	case "29", "SIGIO", "SIGPOLL", "IO", "POLL":
		return syscall.SIGIO, nil
	case "30", "SIGPWR", "PWR":
		return syscall.SIGPWR, nil
	case "31", "SIGSYS", "SIGUNUSED", "SYS", "UNUSED":
		return syscall.SIGSYS, nil
	}

	return 0, fmt.Errorf("unknown signal value: %s", sig)
}

// All is a convinience variable that contains all known signals exactly once.
var All = []os.Signal{
	syscall.SIGHUP,    // "1", "SIGHUP"
	syscall.SIGINT,    // "2", "SIGINT"
	syscall.SIGQUIT,   // "3", "SIGQUIT"
	syscall.SIGILL,    // "4", "SIGILL"
	syscall.SIGTRAP,   // "5", "SIGTRAP"
	syscall.SIGABRT,   // "6", "SIGABRT", "SIGIOT"
	syscall.SIGBUS,    // "7", "SIGBUS"
	syscall.SIGFPE,    // "8", "SIGFPE"
	syscall.SIGKILL,   // "9", "SIGKILL"
	syscall.SIGUSR1,   // "10", "SIGUSR1"
	syscall.SIGSEGV,   // "11", "SIGSEGV"
	syscall.SIGUSR2,   // "12", "SIGUSR2"
	syscall.SIGPIPE,   // "13", "SIGPIPE"
	syscall.SIGALRM,   // "14", "SIGALRM"
	syscall.SIGTERM,   // "15", "SIGTERM"
	syscall.SIGSTKFLT, // "16", "SIGSTKFLT"
	syscall.SIGCHLD,   // "17", "SIGCHLD", "SIGCLD": only useful for zombie reaping
	syscall.SIGCONT,   // "18", "SIGCONT"
	syscall.SIGSTOP,   // "19", "SIGSTOP"
	syscall.SIGTSTP,   // "20", "SIGTSTP"
	syscall.SIGTTIN,   // "21", "SIGTTIN"
	syscall.SIGTTOU,   // "22", "SIGTTOU"
	syscall.SIGURG,    // "23", "SIGURG"
	syscall.SIGXCPU,   // "24", "SIGXCPU"
	syscall.SIGXFSZ,   // "25", "SIGXFSZ"
	syscall.SIGVTALRM, // "26", "SIGVTALRM"
	syscall.SIGPROF,   // "27", "SIGPROF"
	syscall.SIGWINCH,  // "28", "SIGWINCH"
	syscall.SIGIO,     // "29", "SIGIO", "SIGPOLL"
	syscall.SIGPWR,    // "30", "SIGPWR"
	syscall.SIGSYS,    // "31", "SIGSYS", "SIGUNUSED"
}
