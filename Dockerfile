FROM ethereum/client-go:alltools-v1.10.23-amd64


RUN apk add cmake boost-dev git bash
RUN git clone --recursive https://github.com/ethereum/solidity.git
WORKDIR /solidity/
RUN git submodule update --init --recursive
RUN ./scripts/build.sh