//go:build gc && go1.25 && !go1.26
// +build gc,go1.25,!go1.26

package goid

type stack struct {
	lo uintptr
	hi uintptr
}
type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    uintptr
	ctxt uintptr
	lr   uintptr
	bp   uintptr
}
type g struct {
	stack        stack
	stackguard0  uintptr
	stackguard1  uintptr
	_panic       uintptr
	_defer       uintptr
	m            uintptr
	sched        gobuf
	syscallsp    uintptr
	syscallpc    uintptr
	syscallbp    uintptr
	stktopsp     uintptr
	param        uintptr
	atomicstatus uint32
	stackLock    uint32
	goid         int64 // Here it is!
}
type gsignalStack struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr
	stktopsp    uintptr
}
type p struct {
	id int32 // Here is pid
}
type m struct {
	g0      uintptr // goroutine with scheduling stack
	morebuf gobuf   // gobuf arg to morestack
	divmod  uint32  // div/mod denominator for arm - known to liblink
	// Fields not known to debuggers.
	procid     uint64       // for debuggers, but offset not hard-coded
	gsignal    uintptr      // signal-handling g
	goSigStack gsignalStack // Go-allocated signal handling stack
	sigmask    sigset       // storage for saved signal mask
	tls        [6]uintptr   // thread-local storage (for x86 extern register)
	mstartfn   func()
	curg       uintptr // current running goroutine
	caughtsig  uintptr // goroutine running during fatal signal
	p          *p      // attached p for executing go code (nil if not executing go code)
}
