package log

import (
	"context"

)

// Infof logs with severity INFO,
// if logging has been enabled for the source file where the call is
// performed at the provided verbosity level, via the vmodule setting.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
// INFO is used for informational messages, when no action
// is required as a result.
func Infof(ctx context.Context, format string, args ...interface{}) {

}

// Info logs with severity INFO.
// It extracts log tags from the context and logs them along with the given
// message.
//
// INFO is used for informational messages, when no action
// is required as a result.
func Info(ctx context.Context, msg string) {
}


// Warningf logs with severity WARNING,
// if logging has been enabled for the source file where the call is
// performed at the provided verbosity level, via the vmodule setting.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
// WARNING is used for situations which may require special handling,
// while normal operation is expected to resume automatically.
func Warningf(ctx context.Context, format string, args ...interface{}) {

}

// Warning logs with severity WARNING.
// It extracts log tags from the context and logs them along with the given
// message.
//
// WARNING is used for situations which may require special handling,
// while normal operation is expected to resume automatically.
func Warning(ctx context.Context, msg string) {

}


// Errorf logs with severity ERROR,
// if logging has been enabled for the source file where the call is
// performed at the provided verbosity level, via the vmodule setting.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
// ERROR is used for situations that require special handling,
// when normal operation could not proceed as expected.
// Other operations can continue mostly unaffected.
func Errorf(ctx context.Context, format string, args ...interface{}) {

}

// Error logs with severity ERROR.
// It extracts log tags from the context and logs them along with the given
// message.
//
// ERROR is used for situations that require special handling,
// when normal operation could not proceed as expected.
// Other operations can continue mostly unaffected.
func Error(ctx context.Context, msg string) {

}


// Fatalf logs with severity FATAL,
// if logging has been enabled for the source file where the call is
// performed at the provided verbosity level, via the vmodule setting.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
// FATAL is used for situations that require an immedate, hard
// server shutdown. A report is also sent to telemetry if telemetry
// is enabled.
func Fatalf(ctx context.Context, format string, args ...interface{}) {

}

// Fatal logs with severity FATAL.
// It extracts log tags from the context and logs them along with the given
// message.
//
// FATAL is used for situations that require an immedate, hard
// server shutdown. A report is also sent to telemetry if telemetry
// is enabled.
func Fatal(ctx context.Context, msg string) {
}

// FatalfDepth logs with severity FATAL,
// offsetting the caller's stack frame by 'depth'.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
// FATAL is used for situations that require an immedate, hard
// server shutdown. A report is also sent to telemetry if telemetry
// is enabled.
func FatalfDepth(ctx context.Context, depth int, format string, args ...interface{}) {
}
