# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

YACC=goyacc

TARG=golambda
GOFILES=\
	expr.go\
	grammar.go\
	parse.go\
	lex.go

include $(GOROOT)/src/Make.pkg

grammar.go: grammar.y
	$(YACC) -o $@ $<