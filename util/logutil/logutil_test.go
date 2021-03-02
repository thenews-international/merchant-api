package logutil_test

import (
	"merchant/util/logutil"
)

func Example_logall() {
	logger, cleanup, err := logutil.NewLogger("*", "light-console", "stdout")
	if err != nil {
		panic(err)
	}
	defer cleanup()

	logger.Debug("top debug")
	logger.Info("top info")
	logger.Warn("top warn")
	logger.Error("top error")

	logger.Named("foo").Debug("foo debug")
	logger.Named("foo").Info("foo info")
	logger.Named("foo").Warn("foo warn")
	logger.Named("foo").Error("foo error")

	// Output:
	// DEBUG	bty               	logutil/logutil_test.go:14	top debug
	// INFO 	bty               	logutil/logutil_test.go:15	top info
	// WARN 	bty               	logutil/logutil_test.go:16	top warn
	// ERROR	bty               	logutil/logutil_test.go:17	top error
	// DEBUG	bty.foo           	logutil/logutil_test.go:19	foo debug
	// INFO 	bty.foo           	logutil/logutil_test.go:20	foo info
	// WARN 	bty.foo           	logutil/logutil_test.go:21	foo warn
	// ERROR	bty.foo           	logutil/logutil_test.go:22	foo error
}

func Example_logerrors() {
	logger, cleanup, err := logutil.NewLogger("*", "light-console", "stdout")
	if err != nil {
		panic(err)
	}
	defer cleanup()

	logger.Debug("top debug")
	logger.Info("top info")
	logger.Warn("top warn")
	logger.Error("top error")

	logger.Named("foo").Debug("foo debug")
	logger.Named("foo").Info("foo info")
	logger.Named("foo").Warn("foo warn")
	logger.Named("foo").Error("foo error")

	logger.Named("foo").Named("bar").Debug("foo.bar debug")
	logger.Named("foo").Named("bar").Info("foo.bar info")
	logger.Named("foo").Named("bar").Warn("foo.bar warn")
	logger.Named("foo").Named("bar").Error("foo.bar error")

	// Output:
	// DEBUG	bty               	logutil/logutil_test.go:42	top debug
	// INFO 	bty               	logutil/logutil_test.go:43	top info
	// WARN 	bty               	logutil/logutil_test.go:44	top warn
	// ERROR	bty               	logutil/logutil_test.go:45	top error
	// DEBUG	bty.foo           	logutil/logutil_test.go:47	foo debug
	// INFO 	bty.foo           	logutil/logutil_test.go:48	foo info
	// WARN 	bty.foo           	logutil/logutil_test.go:49	foo warn
	// ERROR	bty.foo           	logutil/logutil_test.go:50	foo error
	// DEBUG	bty.foo.bar       	logutil/logutil_test.go:52	foo.bar debug
	// INFO 	bty.foo.bar       	logutil/logutil_test.go:53	foo.bar info
	// WARN 	bty.foo.bar       	logutil/logutil_test.go:54	foo.bar warn
	// ERROR	bty.foo.bar       	logutil/logutil_test.go:55	foo.bar error

}
