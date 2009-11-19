PROGRAM=golambda
TESTPROG=test_$(PROGRAM)

# Define all the commands we're going to use based on the arch.
include $(GOROOT)/src/Make.$(GOARCH)
YACC=goyacc

GOLAMBDA_FILES=\
	abstraction.go\
	application.go\
	expr.go\
	grammar.go\
	group.go\
	lex.go\
	parse.go\
	variable.go

TEST_FILES=$(wildcard *_test.go)

all: $(PROGRAM)

test: $(TESTPROG)
	./$(TESTPROG)

clean:
	rm -rf *.$O $(PROGRAM) $(TESTPROG) grammar.go

golambda.$O: $(GOLAMBDA_FILES)
	$(GC) -o $@ $^

main.$O: golambda.$O
test_main.$O: test_golambda.$O

$(PROGRAM): main.$O
	$(LD) -o $@ $^

test_golambda.$O: $(GOLAMBDA_FILES) $(TEST_FILES)
	$(GC) -o $@ $^

$(TESTPROG): test_main.$O
	$(LD) -o $@ $^

grammar.go: grammar.y
	$(YACC) -o $@ $<

.SUFFIXES: .$O .go

.go.$O:
	$(GC) -o $@ $<