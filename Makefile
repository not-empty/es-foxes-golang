.PHONY: build
build:
	rm -f es-foxes \
	&& docker build --target base -t gobuild . \
	&& docker run --rm --name esfoxes-build -v ${PWD}:/home/app -w /home/app gobuild go build -o es-foxes ./