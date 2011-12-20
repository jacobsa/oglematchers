include $(GOROOT)/src/Make.inc

TARG = github.com/jacobsa/oglematchers
GOFILES = \
	any_of.go \
	equals.go \
	less_than.go \
	matcher.go \
	not.go \

include $(GOROOT)/src/Make.pkg
