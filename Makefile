include $(GOROOT)/src/Make.inc

TARG=github.com/jacobsa/ogletest
GOFILES=\
	equals.go \
	matcher.go \

include $(GOROOT)/src/Make.pkg
