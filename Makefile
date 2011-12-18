include $(GOROOT)/src/Make.inc

TARG=mypackage
GOFILES=\
	equality_matchers.go \
	matcher.go \

include $(GOROOT)/src/Make.pkg
