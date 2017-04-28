go:
				glide up
				go build

all:
				go
clean:
				rm -rf snap-plugin-processor-metric-repeat-filter

test:
				go test -v $$(glide novendor)
