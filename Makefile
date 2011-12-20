include $(GOROOT)/src/Make.inc

TARG=github.com/jacobsa/ogletest
GOFILES=\
	any_of.go \
	equals.go \
	expect_that.go \
	less_than.go \
	matcher.go \
	not.go \
	register_test_suite.go \
	run_tests.go \

include $(GOROOT)/src/Make.pkg
