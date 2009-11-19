# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

PROGRAM=golambda

include $(GOROOT)/src/Make.$(GOARCH)

YACC=goyacc

TARG=golambda

GOLAMBDA_FILES=\
	expr.go\
	grammar.go\
	parse.go\
	lex.go

MAIN_FILES=\
	main.go

OFILES=\
	golambda.$O\
	main.$O

all: $(PROGRAM)

clean:
	rm -rf *.$O $(PROGRAM) grammar.go

golambda.$O: $(GOLAMBDA_FILES)
	$(GC) -o $@ $(GOLAMBDA_FILES)

main.$O: golambda.$O $(MAIN_FILES)
	$(GC) -o $@ $(MAIN_FILES)

$(PROGRAM): $(OFILES)
	$(LD) -o $@ $(OFILES)

grammar.go: grammar.y
	$(YACC) -o $@ $<