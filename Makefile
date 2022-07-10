GO=go

pack:
	$(GO) build -x -v -o swagsync -race > build.log
	rm -rf ./out
	rm -f ./out.zip
	mkdir -p out
	mv ./swagsync ./out/
	cp ./config/demo.yaml ./out/swagsync.yaml
	zip -r swagsync.zip ./out
