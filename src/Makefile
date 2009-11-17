PROGRAM=golambda
O=6 # x86_64
LD=6l
GC=6g

all: $(PROGRAM)

clean:
	rm -rf *.$O $(PROGRAM)

$(PROGRAM): $(OBJS) main.$O
	${LD} -o $@ main.$O $(OBJS)

.SUFFIXES: .go .6

.go.6:
	${GC} $<