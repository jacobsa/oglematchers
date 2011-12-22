include $(GOROOT)/src/Make.inc

TARG = github.com/jacobsa/oglematchers
GOFILES = \
	any_of.go \
	equals.go \
	error.go \
	has_substr.go \
	less_or_equal.go \
	less_than.go \
	matcher.go \
	not.go \
	panics.go \

include $(GOROOT)/src/Make.pkg
