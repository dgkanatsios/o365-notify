VERSION = 0.0.1
REGISTRY ?= docker.io
REPO ?= dgkanatsios
EXE_NAME = o365-notify

.PHONY: build
build:
	go build -o o365-notify *.go

.PHONY: builddocker
builddocker: build
	docker build -t $(REGISTRY)/$(REPO)/$(EXE_NAME):$(VERSION) .

.PHONY: pushdocker 
pushdocker:
	docker push $(REGISTRY)/$(REPO)/$(EXE_NAME):$(VERSION)
	docker tag $(REGISTRY)/$(REPO)/$(EXE_NAME):$(VERSION) $(REGISTRY)/$(REPO)/$(EXE_NAME):latest

.PHONY: clean	
clean:
	rm $(EXE_NAME)
