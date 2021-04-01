package signals

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

// https://golang.org/pkg/syscall/

// Parse matches signal name (number) to internal type.
func Parse(sig string) (unix.Signal, error) { //nolint: cyclop // to avoid using big map
	switch strings.ToUpper(strings.TrimSpace(sig)) {
	case "1", "SIGHUP", "HUP":
		return unix.SIGHUP, nil
	case "2", "SIGINT", "INT":
		return unix.SIGINT, nil
	case "3", "SIGQUIT", "QUIT":
		return unix.SIGQUIT, nil
	case "4", "SIGILL", "ILL":
		return unix.SIGILL, nil
	case "5", "SIGTRAP", "TRAP":
		return unix.SIGTRAP, nil
	case "6", "SIGABRT", "SIGIOT", "ABRT", "IOT":
		return unix.SIGABRT, nil
	case "7", "SIGBUS", "BUS":
		return unix.SIGBUS, nil
	case "8", "SIGFPE", "FPE":
		return unix.SIGFPE, nil
	case "9", "SIGKILL", "KILL":
		return unix.SIGKILL, nil
	case "10", "SIGUSR1", "USR1":
		return unix.SIGUSR1, nil
	case "11", "SIGSEGV", "SEGV":
		return unix.SIGSEGV, nil
	case "12", "SIGUSR2", "USR2":
		return unix.SIGUSR2, nil
	case "13", "SIGPIPE", "PIPE":
		return unix.SIGPIPE, nil
	case "14", "SIGALRM", "ALRM":
		return unix.SIGALRM, nil
	case "15", "SIGTERM", "TERM":
		return unix.SIGTERM, nil
	case "16", "SIGSTKFLT", "STKFLT":
		return unix.SIGSTKFLT, nil
	case "17", "SIGCHLD", "SIGCLD", "CHLD", "CLD":
		return unix.SIGCHLD, nil
	case "18", "SIGCONT", "CONT":
		return unix.SIGCONT, nil
	case "19", "SIGSTOP", "STOP":
		return unix.SIGSTOP, nil
	case "20", "SIGTSTP", "TSTP":
		return unix.SIGTSTP, nil
	case "21", "SIGTTIN", "TTIN":
		return unix.SIGTTIN, nil
	case "22", "SIGTTOU", "TTOU":
		return unix.SIGTTOU, nil
	case "23", "SIGURG", "URG":
		return unix.SIGURG, nil
	case "24", "SIGXCPU", "XCPU":
		return unix.SIGXCPU, nil
	case "25", "SIGXFSZ", "XFSZ":
		return unix.SIGXFSZ, nil
	case "26", "SIGVTALRM", "VTALRM":
		return unix.SIGVTALRM, nil
	case "27", "SIGPROF", "PROF":
		return unix.SIGPROF, nil
	case "28", "SIGWINCH", "WINCH":
		return unix.SIGWINCH, nil
	case "29", "SIGIO", "SIGPOLL", "IO", "POLL":
		return unix.SIGIO, nil
	case "30", "SIGPWR", "PWR":
		return unix.SIGPWR, nil
	case "31", "SIGSYS", "SIGUNUSED", "SYS", "UNUSED":
		return unix.SIGSYS, nil
	}

	return 0, fmt.Errorf("unknown signal value: %s", sig)
}

// All is a convinience variable that contains all known signals exactly once.
var All = []os.Signal{ //nolint: gochecknoglobals // list of Linux signals
	unix.SIGHUP,    // "1", "SIGHUP"
	unix.SIGINT,    // "2", "SIGINT"
	unix.SIGQUIT,   // "3", "SIGQUIT"
	unix.SIGILL,    // "4", "SIGILL"
	unix.SIGTRAP,   // "5", "SIGTRAP"
	unix.SIGABRT,   // "6", "SIGABRT", "SIGIOT"
	unix.SIGBUS,    // "7", "SIGBUS"
	unix.SIGFPE,    // "8", "SIGFPE"
	unix.SIGKILL,   // "9", "SIGKILL"
	unix.SIGUSR1,   // "10", "SIGUSR1"
	unix.SIGSEGV,   // "11", "SIGSEGV"
	unix.SIGUSR2,   // "12", "SIGUSR2"
	unix.SIGPIPE,   // "13", "SIGPIPE"
	unix.SIGALRM,   // "14", "SIGALRM"
	unix.SIGTERM,   // "15", "SIGTERM"
	unix.SIGSTKFLT, // "16", "SIGSTKFLT"
	unix.SIGCHLD,   // "17", "SIGCHLD", "SIGCLD": only useful for zombie reaping
	unix.SIGCONT,   // "18", "SIGCONT"
	unix.SIGSTOP,   // "19", "SIGSTOP"
	unix.SIGTSTP,   // "20", "SIGTSTP"
	unix.SIGTTIN,   // "21", "SIGTTIN"
	unix.SIGTTOU,   // "22", "SIGTTOU"
	unix.SIGURG,    // "23", "SIGURG"
	unix.SIGXCPU,   // "24", "SIGXCPU"
	unix.SIGXFSZ,   // "25", "SIGXFSZ"
	unix.SIGVTALRM, // "26", "SIGVTALRM"
	unix.SIGPROF,   // "27", "SIGPROF"
	unix.SIGWINCH,  // "28", "SIGWINCH"
	unix.SIGIO,     // "29", "SIGIO", "SIGPOLL"
	unix.SIGPWR,    // "30", "SIGPWR"
	unix.SIGSYS,    // "31", "SIGSYS", "SIGUNUSED"
}
