
# goweb3manager


cd $GOPATH/src/github.com/ethereum/go-ethereum/

go build
go install
make
make devtools

sudo add-apt-repository ppa:ethereum/ethereum
 sudo apt-get update
 sudo apt-get install solc




abigen --abi=AccessControllerInterface.abi --pkg=AccessControllerInterface --out=AccessControllerInterface.go
abigen --abi=AggregatorInterface.abi --pkg=AggregatorInterface --out=AggregatorInterface.go
abigen --abi=AggregatorProxy.abi --pkg=AggregatorProxy --out=AggregatorProxy.go
abigen --abi=AggregatorV2V3Interface.abi --pkg=AggregatorV2V3Interface --out=AggregatorV2V3Interface.go
abigen --abi=AggregatorV3Interface.abi --pkg=AggregatorV3Interface --out=AggregatorV3Interface.go
abigen --abi=EACAggregatorProxy.abi --pkg=EACAggregatorProxy --out=EACAggregatorProxy.go
abigen --abi=Owned.abi --pkg=Owned --out=Owned.go

mkdir AccessControllerInterface
mkdir AggregatorInterface
mkdir AggregatorProxy
mkdir AggregatorV2V3Interface
mkdir AggregatorV3Interface
mkdir EACAggregatorProxy
mkdir Owned


solc AccessControllerInterface.sol --bin --abi -o . --overwrite
solc AggregatorInterface.sol --bin --abi -o . --overwrite
solc AggregatorProxy.sol --bin --abi -o . --overwrite
solc AggregatorV2V3Interface.sol --bin --abi -o . --overwrite
solc AggregatorV3Interface.sol --bin --abi -o . --overwrite
solc EACAggregatorProxy.sol --bin --abi -o . --overwrite
solc Owned.sol --bin --abi -o . --overwrite



solc ManagedWallet.sol --bin --abi --optimize -o ./
