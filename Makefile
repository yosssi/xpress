install:
	bower install
	sudo npm install
run:
	go run main.go xpress &
stop:
	ps aux | grep xpress | grep exe/main | grep -v grep | awk '{print $$2}' | xargs kill -9
rerun:
	make stop
	make run
