
CUR_DIR=$(cd "$(dirname "${BASH_SOURCE-$0}")"; pwd)

go build -o ${CUR_DIR}/build/chaincode -v .