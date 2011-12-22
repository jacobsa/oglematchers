include $(GOROOT)/src/Make.inc

TARG = github.com/jacobsa/oglematchers
GOFILES = \
	any_of.go \
	equals.go \
	error.go \
	greater_than.go \
	has_substr.go \
	less_or_equal.go \
	less_than.go \
	matcher.go \
	not.go \
	panics.go \
	transform_description.go \

include $(GOROOT)/src/Make.pkg
