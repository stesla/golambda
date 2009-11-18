# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

YACC=goyacc

TARG=golambda
GOFILES=\
	expr.go\
	parse.go

include $(GOROOT)/src/Make.pkg

parse.go: parse.y
	$(YACC) -o $@ $<