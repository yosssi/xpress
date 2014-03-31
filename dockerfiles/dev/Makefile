build:
	docker build -t yosssi/xpress-dev:1.0.0 .
run:
	docker run -i -t -v $$HOME:/host -p 8080:8080 -p 9200:9200 yosssi/xpress-dev:1.0.0 /bin/bash
rm:
	docker ps -a | grep Exit | cut -d " " -f 1 | xargs docker rm
