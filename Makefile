include $(GOROOT)/src/Make.inc

TARG=github.com/jacobsa/ogletest
GOFILES=\
	any_of.go \
	equals.go \
	matcher.go \
	not.go \

include $(GOROOT)/src/Make.pkg
