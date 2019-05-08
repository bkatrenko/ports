This is simple app which contains two services for handle/retrieve some data.

Use docker-compose/Makefile to build and start all services.

Start:

[docker-compose up] --build or [make docker-up]

Testing:

make integration-test


NOTE:
!!!!!!!!!!!!WARN!!!!!!!!!!!!!!!!!!!!!!!!

Add .json file with data to root of the repo to make it all working (this repo is NOT contain this file according to test task).

!!!!!!!!!!!!WARN!!!!!!!!!!!!!!!!!!!!!!!!

This task was take about 5 hours for done. Some test has limitations for assertions count - because I need more time to add it (it's simple but can really take some time).



