build:
	CGO_ENABLED=0 go build ${LDFLAGS} -o ap github.com/ravisantoshgudimetla/aggregation-parser
