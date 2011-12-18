include $(GOROOT)/src/Make.inc

TARG=mypackage
GOFILES=\
	equals.go \
	matcher.go \

include $(GOROOT)/src/Make.pkg
