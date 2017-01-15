.PHONY:	build-dev

DEV_SERVER = vsdict
DISTRIBUTED_FILE = ev-dict-server
INIT_STARTUP_FILE = ev-dict-init
INIT_STARTUP_FOLDER = /etc/init.d
REMOTE_DIR = ~/ev-dict-server
DATABASE_FILE = dict_en_vi.sql

build-dev:
	gin -a 8000 -b ev-dict-server -i

deploy:
	env GOOS=linux GOARCH=amd64 go build -o ev-dict-server
	ssh $(DEV_SERVER) "mkdir -p $(REMOTE_DIR) || true"
	rsync -azvv --progress --update $(DATABASE_FILE) $(DEV_SERVER):$(REMOTE_DIR)
	rsync -azvv --progress --update $(INIT_STARTUP_FILE) $(DEV_SERVER):$(REMOTE_DIR)
	ssh $(DEV_SERVER) "sudo cp $(REMOTE_DIR)/$(INIT_STARTUP_FILE) $(INIT_STARTUP_FOLDER)"
	ssh $(DEV_SERVER) "sudo chmod u+x $(INIT_STARTUP_FOLDER)/$(INIT_STARTUP_FILE)"
	rsync -azvv --progress --update conf $(DEV_SERVER):$(REMOTE_DIR)
	rsync -azvv --progress  $(DISTRIBUTED_FILE) $(DEV_SERVER):$(REMOTE_DIR)
	ssh $(DEV_SERVER) "sudo service $(INIT_STARTUP_FILE) restart"
