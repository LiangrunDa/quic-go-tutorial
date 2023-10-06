.PHONY:	test
test: setup-env.sh run-client.sh run-server.sh client server optimize VERSION
	mkdir -p testArtifact
	cp setup-env.sh run-client.sh run-server.sh client server optimize_client optimize_server VERSION testArtifact

client: */*.go */*/*.go
	./scripts/build_client.sh

server: */*.go */*/*.go
	./scripts/build_server.sh

optimize: */*.go */*/*.go
	./scripts/build_optimize.sh

VERSION: 
	git rev-parse HEAD > VERSION

clean:
	rm ./client ./server ./optimize_client ./optimize_server VERSION
	rm -r ./testArtifact
