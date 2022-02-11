export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = ztserver ztclient

all: $(OBJ)

$(OBJ):
	go mod tidy && go build -gcflags "-N -l" -o $@ ./$(subst zt,,$@)

clean:
	rm -fr $(OBJ)

-include .deps
dep:
	echo -n "$(OBJ):" > .deps
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps