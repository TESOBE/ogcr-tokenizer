// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ParcelNFTParcelData is an auto generated low-level Go binding around an user-defined struct.
type ParcelNFTParcelData struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}

// ParcelNFTMetaData contains all meta data concerning the ParcelNFT contract.
var ParcelNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParcel\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structParcelNFT.ParcelData\",\"components\":[{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_hash\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenIdByParcelId\",\"inputs\":[{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_hash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"parcels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_hash\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenIdByParcelId\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterUpdated\",\"inputs\":[{\"name\":\"oldMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParcelMinted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"parcel_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080346200041857601f1990601f6001600160401b0362001ff038819003838101861685018381118682101762000402578592829160405283396020948591810103126200041857516001600160a01b03808216949092909185900362000418576200006a6200041d565b93601185527011d95bd9dc985c1a1a58c814185c98d95b607a1b83860152620000926200041d565b600681526514105490d15360d21b8482015285519783891162000402576000988954976001988981811c91168015620003f7575b88821014620003e35790818684931162000390575b5087908683116001146200032f578c9262000323575b5050600019600383901b1c191690881b1789555b81519384116200030f5786548781811c9116801562000304575b86821014620002f057908184869594931162000299575b508592841160011462000238575088926200022c575b5050600019600383901b1c191690841b1783555b6006549260018060a01b0319933385821617600655604051933391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08880a36007558315620001eb57505081906008541617600855604051917f1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b738180a3611bb290816200043e8239f35b62461bcd60e51b82526004820152601660248201527f50617263656c4e46543a207a65726f206d696e74657200000000000000000000604482015260649150fd5b0151905038806200014c565b878a52858a20889590939291168a5b8782821062000282575050841162000268575b505050811b01835562000160565b015160001960f88460031b161c191690553880806200025a565b8385015186558a9790950194938401930162000247565b9091929350878a52858a208480870160051c820192888810620002e6575b9187968b92969594930160051c01915b828110620002d757505062000136565b8c81558796508a9101620002c7565b92508192620002b7565b634e487b7160e01b8a52602260045260248afd5b90607f16906200011f565b634e487b7160e01b89526041600452602489fd5b015190503880620000f1565b908c91858c95168380528a8420935b8b8282106200037957505084116200035f575b505050811b01895562000105565b015160001960f88460031b161c1916905538808062000351565b8385015186558e979095019493840193016200033e565b9091508b8052878c208680850160051c8201928a8610620003d9575b918c91869594930160051c01915b828110620003ca575050620000db565b8e81558594508c9101620003ba565b92508192620003ac565b634e487b7160e01b8c52602260045260248cfd5b90607f1690620000c6565b634e487b7160e01b600052604160045260246000fd5b600080fd5b60408051919082016001600160401b03811183821017620004025760405256fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146113125750806306fdde031461126f5780630754617214611246578063081812fc14611228578063095ea7b3146110ae57806323b872dd1461108557806342842e0e1461105d5780634a0706be1461095c5780634d0434f21461083557806361b8ce8c146108175780636352211e146107e75780636a94eb851461077a57806370a08231146106e3578063715018a6146106865780638da5cb5b1461065d5780638dfa037f146105bf57806395d89b41146104e0578063a0fa8ede14610463578063a22cb46514610391578063b88d4fde14610330578063c87b56dd146102bf578063e985e9c514610269578063f2fde38b146101a05763fca3b5aa1461012757600080fd5b3461019b57602036600319011261019b576101406113c5565b6101486115dc565b6001600160a01b039081169061015f821515611ae4565b816008549182167f1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73600080a36001600160a01b03191617600855005b600080fd5b3461019b57602036600319011261019b576101b96113c5565b6101c16115dc565b6001600160a01b0390811690811561021557600654826bffffffffffffffffffffffff60a01b821617600655167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461019b57604036600319011261019b576102826113c5565b61028a6113db565b9060018060a01b03809116600052600560205260406000209116600052602052602060ff604060002054166040519015158152f35b3461019b57602036600319011261019b576004356000908152600260205260409020546102f6906001600160a01b03161515611634565b60006040516103048161146e565b5261032c6040516103148161146e565b600081526040519182916020835260208301906113a0565b0390f35b3461019b57608036600319011261019b576103496113c5565b6103516113db565b606435916001600160401b03831161019b573660238401121561019b5761038561038f9336906024816004013591016114c5565b9160443591611746565b005b3461019b57604036600319011261019b576103aa6113c5565b6024359081151580920361019b576001600160a01b03169033821461041e57336000526005602052604060002082600052602052604060002060ff1981541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b60405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606490fd5b3461019b57602036600319011261019b5760043560005260096020526104c4604060002061032c61049382611536565b916104d26104af60026104a860018501611536565b9301611536565b916040519586956060875260608701906113a0565b9085820360208701526113a0565b9083820360408501526113a0565b3461019b57600036600319011261019b5760405160006001805490610504826114fc565b80855291818116908115610598575060011461053f575b61032c8461052b81860382611489565b6040519182916020835260208301906113a0565b600081815292507fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b82841061058057505050810160200161052b8261051b565b80546020858701810191909152909301928101610568565b60ff191660208087019190915292151560051b8501909201925061052b915083905061051b565b3461019b57602036600319011261019b576004356001600160401b03811161019b576105f16020913690600401611426565b91908260405193849283378101600a81520301902054801561061857602090604051908152f35b60405162461bcd60e51b815260206004820152601e60248201527f50617263656c4e46543a2070617263656c5f6964206e6f7420666f756e6400006044820152606490fd5b3461019b57600036600319011261019b576006546040516001600160a01b039091168152602090f35b3461019b57600036600319011261019b5761069f6115dc565b600680546001600160a01b031981169091556000906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461019b57602036600319011261019b576001600160a01b036107046113c5565b1680156107235760005260036020526020604060002054604051908152f35b60405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608490fd5b3461019b57602036600319011261019b576004356001600160401b03811161019b573660238201121561019b576107d460206107c1819336906024816004013591016114c5565b816040519382858094519384920161137d565b8101600a81520301902054604051908152f35b3461019b57602036600319011261019b576020610805600435611680565b6040516001600160a01b039091168152f35b3461019b57600036600319011261019b576020600754604051908152f35b3461019b5760208060031936011261019b576004359060606040805161085a81611453565b82815283810183905201526000828152600260205260409020546001600160a01b031615610918576060916000526009815261032c604060002091604051926108a284611453565b6108ab81611536565b84526109076108f16108d160026108c460018601611536565b9486890195865201611536565b9360408701948552604051978897828952519188015260808701906113a0565b915191601f1992838783030160408801526113a0565b9151908483030160608501526113a0565b6064906040519062461bcd60e51b82526004820152601f60248201527f50617263656c4e46543a20746f6b656e20646f6573206e6f74206578697374006044820152fd5b3461019b57608036600319011261019b576109756113c5565b6024356001600160401b03811161019b57610994903690600401611426565b6044356001600160401b03811161019b576109b3903690600401611426565b9290936064356001600160401b03811161019b576109d5903690600401611426565b6008546001600160a01b03979196929088163303611018576109fa8885161515611ae4565b8515610fd35760405186868237602081888101600a81520301902054610f8e57600754966000198814610f78576001880160075588851615610f3457600088815260026020526040902054610b0e93610b0291610a63906001600160a01b031615155b15611b30565b60008a815260026020526040902054610a86906001600160a01b03161515610a5d565b868b166000818152600360209081526040808320805460010190558d8352600290915280822080546001600160a01b031916841790555197918c917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a4610aee87611453565b610af9368b8b6114c5565b875236916114c5565b602085015236916114c5565b604082015284600052600960205260406000209080518051906001600160401b038211610d9a578190610b4185546114fc565b601f8111610ee4575b50602090601f8311600114610e7e57600092610e73575b50508160011b916000199060031b1c19161782555b60208101518051906001600160401b038211610d9a57610b9960018501546114fc565b601f8111610e2c575b50602090601f8311600114610dbb57604093929160009183610db0575b50508160011b916000199060031b1c19161760018401555b01518051906001600160401b038211610d9a57610bf760028401546114fc565b601f8111610d4f575b509285926020989288958a90601f8311600114610cb357827fe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364989995936040979593600293600092610ca8575b50508160011b916000199060031b1c1916179101555b868451868582378b81888101600a8152030190205583519485938b8552818c8601528585013760008389018501521695601f01601f19168101030190a3604051908152f35b015190508e80610c4d565b90600284016000528b6000209160005b601f1985168110610d3257508360409795936002936001937fe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e13649c9d9997601f19811610610d19575b505050811b01910155610c63565b015160001960f88460031b161c191690558e8080610d0b565b8282015184558c99508b9850600190930192918d01918d01610cc3565b600284016000526020600020601f840160051c81019160208510610d90575b601f0160051c01905b818110610d845750610c00565b60008155600101610d77565b9091508190610d6e565b634e487b7160e01b600052604160045260246000fd5b015190508a80610bbf565b90601f198316916001860160005260206000209260005b818110610e145750916001939185604097969410610dfb575b505050811b016001840155610bd7565b015160001960f88460031b161c191690558a8080610deb565b92936020600181928786015181550195019301610dd2565b600185016000526020600020601f840160051c810160208510610e6c575b601f830160051c82018110610e60575050610ba2565b60008155600101610e4a565b5080610e4a565b015190508980610b61565b6000868152602081209350601f198516905b818110610ecc5750908460019594939210610eb3575b505050811b018255610b76565b015160001960f88460031b161c19169055898080610ea6565b92936020600181928786015181550195019301610e90565b909150846000526020600020601f840160051c810160208510610f2d575b90849392915b601f830160051c82018110610f1e575050610b4a565b60008155859450600101610f08565b5080610f02565b606460405162461bcd60e51b815260206004820152602060248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152fd5b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b815260206004820152601960248201527f50617263656c4e46543a20616c7265616479206d696e746564000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601d60248201527f50617263656c4e46543a2070617263656c5f69642072657175697265640000006044820152606490fd5b60405162461bcd60e51b815260206004820152601f60248201527f50617263656c4e46543a2063616c6c6572206973206e6f74206d696e746572006044820152606490fd5b3461019b5761038f61106e366113f1565b906040519261107c8461146e565b60008452611746565b3461019b5761038f611096366113f1565b916110a96110a484336117e1565b6116e4565b6118a9565b3461019b57604036600319011261019b576110c76113c5565b602435906001600160a01b0380806110de85611680565b169216918083146111d9578033149081156111b4575b501561114957600083815260046020526040902080546001600160a01b0319168317905561112183611680565b167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600080a4005b60405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608490fd5b9050600052600560205260406000203360005260205260ff60406000205416846110f4565b60405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608490fd5b3461019b57602036600319011261019b5760206108056004356116a6565b3461019b57600036600319011261019b576008546040516001600160a01b039091168152602090f35b3461019b57600036600319011261019b5760405160008054611290816114fc565b8084529060019081811690811561059857506001146112b95761032c8461052b81860382611489565b600080805292507f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b8284106112fa57505050810160200161052b8261051b565b805460208587018101919091529093019281016112e2565b3461019b57602036600319011261019b576004359063ffffffff60e01b821680920361019b576020916380ac58cd60e01b811490811561136c575b811561135b575b5015158152f35b6301ffc9a760e01b14905083611354565b635b5e139f60e01b8114915061134d565b60005b8381106113905750506000910152565b8181015183820152602001611380565b906020916113b98151809281855285808601910161137d565b601f01601f1916010190565b600435906001600160a01b038216820361019b57565b602435906001600160a01b038216820361019b57565b606090600319011261019b576001600160a01b0390600435828116810361019b5791602435908116810361019b579060443590565b9181601f8401121561019b578235916001600160401b03831161019b576020838186019501011161019b57565b606081019081106001600160401b03821117610d9a57604052565b602081019081106001600160401b03821117610d9a57604052565b90601f801991011681019081106001600160401b03821117610d9a57604052565b6001600160401b038111610d9a57601f01601f191660200190565b9291926114d1826114aa565b916114df6040519384611489565b82948184528183011161019b578281602093846000960137010152565b90600182811c9216801561152c575b602083101461151657565b634e487b7160e01b600052602260045260246000fd5b91607f169161150b565b906040519182600082549261154a846114fc565b9081845260019485811690816000146115b95750600114611576575b505061157492500383611489565b565b9093915060005260209081600020936000915b8183106115a157505061157493508201013880611566565b85548884018501529485019487945091830191611589565b91505061157494506020925060ff191682840152151560051b8201013880611566565b6006546001600160a01b031633036115f057565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b1561163b57565b60405162461bcd60e51b815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606490fd5b6000908152600260205260409020546001600160a01b03166116a3811515611634565b90565b6000818152600260205260409020546116c9906001600160a01b03161515611634565b6000908152600460205260409020546001600160a01b031690565b156116eb57565b60405162461bcd60e51b815260206004820152602d60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201526c1c881bdc88185c1c1c9bdd9959609a1b6064820152608490fd5b9061176a93929161175a6110a484336117e1565b6117658383836118a9565b6119bf565b1561177157565b60405162461bcd60e51b81528061178a6004820161178e565b0390fd5b60809060208152603260208201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b60608201520190565b906001600160a01b0380806117f584611680565b16931691838314938415611828575b508315611812575b50505090565b61181e919293506116a6565b161438808061180c565b909350600052600560205260406000208260005260205260ff604060002054169238611804565b1561185657565b60405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608490fd5b906118d1916118b784611680565b6001600160a01b039391841692849290918316841461184f565b1691821561196e57816118ee916118e786611680565b161461184f565b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60008481526004602052604081206bffffffffffffffffffffffff60a01b9081815416905583825260036020526040822060001981540190558482526040822060018154019055858252600260205284604083209182541617905580a4565b60405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b9293600093909291803b15611ad957948491611a199660405180948193630a85bd0160e11b9788845233600485015260018060a01b0380921660248501526044840152608060648401528260209b8c9760848301906113a0565b0393165af1849181611a95575b50611a84575050503d600014611a7c573d611a40816114aa565b90611a4e6040519283611489565b81528091833d92013e5b80519182611a795760405162461bcd60e51b81528061178a6004820161178e565b01fd5b506060611a58565b6001600160e01b0319161492509050565b9091508581813d8311611ad2575b611aad8183611489565b81010312611ace57516001600160e01b031981168103611ace579038611a26565b8480fd5b503d611aa3565b505050915050600190565b15611aeb57565b60405162461bcd60e51b815260206004820152601760248201527f50617263656c4e46543a207a65726f20616464726573730000000000000000006044820152606490fd5b15611b3757565b60405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606490fdfea2646970667358221220f441e70355dac8704b447bfbbe2e0bb6e205efd3e78fad4ae931eceae25183cb64736f6c63430008130033",
}

// ParcelNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use ParcelNFTMetaData.ABI instead.
var ParcelNFTABI = ParcelNFTMetaData.ABI

// ParcelNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ParcelNFTMetaData.Bin instead.
var ParcelNFTBin = ParcelNFTMetaData.Bin

// DeployParcelNFT deploys a new Ethereum contract, binding an instance of ParcelNFT to it.
func DeployParcelNFT(auth *bind.TransactOpts, backend bind.ContractBackend, _minter common.Address) (common.Address, *types.Transaction, *ParcelNFT, error) {
	parsed, err := ParcelNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ParcelNFTBin), backend, _minter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ParcelNFT{ParcelNFTCaller: ParcelNFTCaller{contract: contract}, ParcelNFTTransactor: ParcelNFTTransactor{contract: contract}, ParcelNFTFilterer: ParcelNFTFilterer{contract: contract}}, nil
}

// ParcelNFT is an auto generated Go binding around an Ethereum contract.
type ParcelNFT struct {
	ParcelNFTCaller     // Read-only binding to the contract
	ParcelNFTTransactor // Write-only binding to the contract
	ParcelNFTFilterer   // Log filterer for contract events
}

// ParcelNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParcelNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParcelNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParcelNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParcelNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParcelNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParcelNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParcelNFTSession struct {
	Contract     *ParcelNFT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParcelNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParcelNFTCallerSession struct {
	Contract *ParcelNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ParcelNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParcelNFTTransactorSession struct {
	Contract     *ParcelNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ParcelNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParcelNFTRaw struct {
	Contract *ParcelNFT // Generic contract binding to access the raw methods on
}

// ParcelNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParcelNFTCallerRaw struct {
	Contract *ParcelNFTCaller // Generic read-only contract binding to access the raw methods on
}

// ParcelNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParcelNFTTransactorRaw struct {
	Contract *ParcelNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParcelNFT creates a new instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFT(address common.Address, backend bind.ContractBackend) (*ParcelNFT, error) {
	contract, err := bindParcelNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParcelNFT{ParcelNFTCaller: ParcelNFTCaller{contract: contract}, ParcelNFTTransactor: ParcelNFTTransactor{contract: contract}, ParcelNFTFilterer: ParcelNFTFilterer{contract: contract}}, nil
}

// NewParcelNFTCaller creates a new read-only instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFTCaller(address common.Address, caller bind.ContractCaller) (*ParcelNFTCaller, error) {
	contract, err := bindParcelNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTCaller{contract: contract}, nil
}

// NewParcelNFTTransactor creates a new write-only instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*ParcelNFTTransactor, error) {
	contract, err := bindParcelNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTTransactor{contract: contract}, nil
}

// NewParcelNFTFilterer creates a new log filterer instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*ParcelNFTFilterer, error) {
	contract, err := bindParcelNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTFilterer{contract: contract}, nil
}

// bindParcelNFT binds a generic wrapper to an already deployed contract.
func bindParcelNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParcelNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParcelNFT *ParcelNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParcelNFT.Contract.ParcelNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParcelNFT *ParcelNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParcelNFT.Contract.ParcelNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParcelNFT *ParcelNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParcelNFT.Contract.ParcelNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParcelNFT *ParcelNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParcelNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParcelNFT *ParcelNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParcelNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParcelNFT *ParcelNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParcelNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ParcelNFT.Contract.BalanceOf(&_ParcelNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ParcelNFT.Contract.BalanceOf(&_ParcelNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.GetApproved(&_ParcelNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.GetApproved(&_ParcelNFT.CallOpts, tokenId)
}

// GetParcel is a free data retrieval call binding the contract method 0x4d0434f2.
//
// Solidity: function getParcel(uint256 tokenId) view returns((string,string,string))
func (_ParcelNFT *ParcelNFTCaller) GetParcel(opts *bind.CallOpts, tokenId *big.Int) (ParcelNFTParcelData, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "getParcel", tokenId)

	if err != nil {
		return *new(ParcelNFTParcelData), err
	}

	out0 := *abi.ConvertType(out[0], new(ParcelNFTParcelData)).(*ParcelNFTParcelData)

	return out0, err

}

// GetParcel is a free data retrieval call binding the contract method 0x4d0434f2.
//
// Solidity: function getParcel(uint256 tokenId) view returns((string,string,string))
func (_ParcelNFT *ParcelNFTSession) GetParcel(tokenId *big.Int) (ParcelNFTParcelData, error) {
	return _ParcelNFT.Contract.GetParcel(&_ParcelNFT.CallOpts, tokenId)
}

// GetParcel is a free data retrieval call binding the contract method 0x4d0434f2.
//
// Solidity: function getParcel(uint256 tokenId) view returns((string,string,string))
func (_ParcelNFT *ParcelNFTCallerSession) GetParcel(tokenId *big.Int) (ParcelNFTParcelData, error) {
	return _ParcelNFT.Contract.GetParcel(&_ParcelNFT.CallOpts, tokenId)
}

// GetTokenIdByParcelId is a free data retrieval call binding the contract method 0x8dfa037f.
//
// Solidity: function getTokenIdByParcelId(string parcel_id) view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) GetTokenIdByParcelId(opts *bind.CallOpts, parcel_id string) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "getTokenIdByParcelId", parcel_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenIdByParcelId is a free data retrieval call binding the contract method 0x8dfa037f.
//
// Solidity: function getTokenIdByParcelId(string parcel_id) view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) GetTokenIdByParcelId(parcel_id string) (*big.Int, error) {
	return _ParcelNFT.Contract.GetTokenIdByParcelId(&_ParcelNFT.CallOpts, parcel_id)
}

// GetTokenIdByParcelId is a free data retrieval call binding the contract method 0x8dfa037f.
//
// Solidity: function getTokenIdByParcelId(string parcel_id) view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) GetTokenIdByParcelId(parcel_id string) (*big.Int, error) {
	return _ParcelNFT.Contract.GetTokenIdByParcelId(&_ParcelNFT.CallOpts, parcel_id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParcelNFT *ParcelNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParcelNFT *ParcelNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ParcelNFT.Contract.IsApprovedForAll(&_ParcelNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParcelNFT *ParcelNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ParcelNFT.Contract.IsApprovedForAll(&_ParcelNFT.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ParcelNFT *ParcelNFTCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ParcelNFT *ParcelNFTSession) Minter() (common.Address, error) {
	return _ParcelNFT.Contract.Minter(&_ParcelNFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) Minter() (common.Address, error) {
	return _ParcelNFT.Contract.Minter(&_ParcelNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParcelNFT *ParcelNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParcelNFT *ParcelNFTSession) Name() (string, error) {
	return _ParcelNFT.Contract.Name(&_ParcelNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParcelNFT *ParcelNFTCallerSession) Name() (string, error) {
	return _ParcelNFT.Contract.Name(&_ParcelNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) NextId() (*big.Int, error) {
	return _ParcelNFT.Contract.NextId(&_ParcelNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) NextId() (*big.Int, error) {
	return _ParcelNFT.Contract.NextId(&_ParcelNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParcelNFT *ParcelNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParcelNFT *ParcelNFTSession) Owner() (common.Address, error) {
	return _ParcelNFT.Contract.Owner(&_ParcelNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) Owner() (common.Address, error) {
	return _ParcelNFT.Contract.Owner(&_ParcelNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.OwnerOf(&_ParcelNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.OwnerOf(&_ParcelNFT.CallOpts, tokenId)
}

// Parcels is a free data retrieval call binding the contract method 0xa0fa8ede.
//
// Solidity: function parcels(uint256 ) view returns(string parcel_id, string parcel_uri, string parcel_hash)
func (_ParcelNFT *ParcelNFTCaller) Parcels(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "parcels", arg0)

	outstruct := new(struct {
		ParcelId   string
		ParcelUri  string
		ParcelHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParcelId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ParcelUri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ParcelHash = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Parcels is a free data retrieval call binding the contract method 0xa0fa8ede.
//
// Solidity: function parcels(uint256 ) view returns(string parcel_id, string parcel_uri, string parcel_hash)
func (_ParcelNFT *ParcelNFTSession) Parcels(arg0 *big.Int) (struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}, error) {
	return _ParcelNFT.Contract.Parcels(&_ParcelNFT.CallOpts, arg0)
}

// Parcels is a free data retrieval call binding the contract method 0xa0fa8ede.
//
// Solidity: function parcels(uint256 ) view returns(string parcel_id, string parcel_uri, string parcel_hash)
func (_ParcelNFT *ParcelNFTCallerSession) Parcels(arg0 *big.Int) (struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}, error) {
	return _ParcelNFT.Contract.Parcels(&_ParcelNFT.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParcelNFT *ParcelNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParcelNFT *ParcelNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParcelNFT.Contract.SupportsInterface(&_ParcelNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParcelNFT *ParcelNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParcelNFT.Contract.SupportsInterface(&_ParcelNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParcelNFT *ParcelNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParcelNFT *ParcelNFTSession) Symbol() (string, error) {
	return _ParcelNFT.Contract.Symbol(&_ParcelNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParcelNFT *ParcelNFTCallerSession) Symbol() (string, error) {
	return _ParcelNFT.Contract.Symbol(&_ParcelNFT.CallOpts)
}

// TokenIdByParcelId is a free data retrieval call binding the contract method 0x6a94eb85.
//
// Solidity: function tokenIdByParcelId(string ) view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) TokenIdByParcelId(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "tokenIdByParcelId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdByParcelId is a free data retrieval call binding the contract method 0x6a94eb85.
//
// Solidity: function tokenIdByParcelId(string ) view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) TokenIdByParcelId(arg0 string) (*big.Int, error) {
	return _ParcelNFT.Contract.TokenIdByParcelId(&_ParcelNFT.CallOpts, arg0)
}

// TokenIdByParcelId is a free data retrieval call binding the contract method 0x6a94eb85.
//
// Solidity: function tokenIdByParcelId(string ) view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) TokenIdByParcelId(arg0 string) (*big.Int, error) {
	return _ParcelNFT.Contract.TokenIdByParcelId(&_ParcelNFT.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParcelNFT *ParcelNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParcelNFT *ParcelNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ParcelNFT.Contract.TokenURI(&_ParcelNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParcelNFT *ParcelNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ParcelNFT.Contract.TokenURI(&_ParcelNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Approve(&_ParcelNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Approve(&_ParcelNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x4a0706be.
//
// Solidity: function mint(address to, string parcel_id, string parcel_uri, string parcel_hash) returns(uint256 tokenId)
func (_ParcelNFT *ParcelNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, parcel_id string, parcel_uri string, parcel_hash string) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "mint", to, parcel_id, parcel_uri, parcel_hash)
}

// Mint is a paid mutator transaction binding the contract method 0x4a0706be.
//
// Solidity: function mint(address to, string parcel_id, string parcel_uri, string parcel_hash) returns(uint256 tokenId)
func (_ParcelNFT *ParcelNFTSession) Mint(to common.Address, parcel_id string, parcel_uri string, parcel_hash string) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Mint(&_ParcelNFT.TransactOpts, to, parcel_id, parcel_uri, parcel_hash)
}

// Mint is a paid mutator transaction binding the contract method 0x4a0706be.
//
// Solidity: function mint(address to, string parcel_id, string parcel_uri, string parcel_hash) returns(uint256 tokenId)
func (_ParcelNFT *ParcelNFTTransactorSession) Mint(to common.Address, parcel_id string, parcel_uri string, parcel_hash string) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Mint(&_ParcelNFT.TransactOpts, to, parcel_id, parcel_uri, parcel_hash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParcelNFT *ParcelNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParcelNFT *ParcelNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParcelNFT.Contract.RenounceOwnership(&_ParcelNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParcelNFT *ParcelNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParcelNFT.Contract.RenounceOwnership(&_ParcelNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParcelNFT *ParcelNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParcelNFT *ParcelNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom0(&_ParcelNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom0(&_ParcelNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParcelNFT *ParcelNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParcelNFT *ParcelNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetApprovalForAll(&_ParcelNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetApprovalForAll(&_ParcelNFT.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ParcelNFT *ParcelNFTTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ParcelNFT *ParcelNFTSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetMinter(&_ParcelNFT.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetMinter(&_ParcelNFT.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParcelNFT *ParcelNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParcelNFT *ParcelNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferOwnership(&_ParcelNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferOwnership(&_ParcelNFT.TransactOpts, newOwner)
}

// ParcelNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ParcelNFT contract.
type ParcelNFTApprovalIterator struct {
	Event *ParcelNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTApproval represents a Approval event raised by the ParcelNFT contract.
type ParcelNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ParcelNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTApprovalIterator{contract: _ParcelNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ParcelNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTApproval)
				if err := _ParcelNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) ParseApproval(log types.Log) (*ParcelNFTApproval, error) {
	event := new(ParcelNFTApproval)
	if err := _ParcelNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ParcelNFT contract.
type ParcelNFTApprovalForAllIterator struct {
	Event *ParcelNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTApprovalForAll represents a ApprovalForAll event raised by the ParcelNFT contract.
type ParcelNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParcelNFT *ParcelNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ParcelNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTApprovalForAllIterator{contract: _ParcelNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParcelNFT *ParcelNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ParcelNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTApprovalForAll)
				if err := _ParcelNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParcelNFT *ParcelNFTFilterer) ParseApprovalForAll(log types.Log) (*ParcelNFTApprovalForAll, error) {
	event := new(ParcelNFTApprovalForAll)
	if err := _ParcelNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the ParcelNFT contract.
type ParcelNFTMinterUpdatedIterator struct {
	Event *ParcelNFTMinterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTMinterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTMinterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTMinterUpdated represents a MinterUpdated event raised by the ParcelNFT contract.
type ParcelNFTMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ParcelNFT *ParcelNFTFilterer) FilterMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*ParcelNFTMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTMinterUpdatedIterator{contract: _ParcelNFT.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ParcelNFT *ParcelNFTFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *ParcelNFTMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTMinterUpdated)
				if err := _ParcelNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterUpdated is a log parse operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ParcelNFT *ParcelNFTFilterer) ParseMinterUpdated(log types.Log) (*ParcelNFTMinterUpdated, error) {
	event := new(ParcelNFTMinterUpdated)
	if err := _ParcelNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParcelNFT contract.
type ParcelNFTOwnershipTransferredIterator struct {
	Event *ParcelNFTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTOwnershipTransferred represents a OwnershipTransferred event raised by the ParcelNFT contract.
type ParcelNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParcelNFT *ParcelNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParcelNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTOwnershipTransferredIterator{contract: _ParcelNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParcelNFT *ParcelNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParcelNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTOwnershipTransferred)
				if err := _ParcelNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParcelNFT *ParcelNFTFilterer) ParseOwnershipTransferred(log types.Log) (*ParcelNFTOwnershipTransferred, error) {
	event := new(ParcelNFTOwnershipTransferred)
	if err := _ParcelNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTParcelMintedIterator is returned from FilterParcelMinted and is used to iterate over the raw logs and unpacked data for ParcelMinted events raised by the ParcelNFT contract.
type ParcelNFTParcelMintedIterator struct {
	Event *ParcelNFTParcelMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTParcelMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTParcelMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTParcelMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTParcelMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTParcelMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTParcelMinted represents a ParcelMinted event raised by the ParcelNFT contract.
type ParcelNFTParcelMinted struct {
	TokenId  *big.Int
	To       common.Address
	ParcelId string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterParcelMinted is a free log retrieval operation binding the contract event 0xe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364.
//
// Solidity: event ParcelMinted(uint256 indexed tokenId, address indexed to, string parcel_id)
func (_ParcelNFT *ParcelNFTFilterer) FilterParcelMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address) (*ParcelNFTParcelMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "ParcelMinted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTParcelMintedIterator{contract: _ParcelNFT.contract, event: "ParcelMinted", logs: logs, sub: sub}, nil
}

// WatchParcelMinted is a free log subscription operation binding the contract event 0xe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364.
//
// Solidity: event ParcelMinted(uint256 indexed tokenId, address indexed to, string parcel_id)
func (_ParcelNFT *ParcelNFTFilterer) WatchParcelMinted(opts *bind.WatchOpts, sink chan<- *ParcelNFTParcelMinted, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "ParcelMinted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTParcelMinted)
				if err := _ParcelNFT.contract.UnpackLog(event, "ParcelMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParcelMinted is a log parse operation binding the contract event 0xe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364.
//
// Solidity: event ParcelMinted(uint256 indexed tokenId, address indexed to, string parcel_id)
func (_ParcelNFT *ParcelNFTFilterer) ParseParcelMinted(log types.Log) (*ParcelNFTParcelMinted, error) {
	event := new(ParcelNFTParcelMinted)
	if err := _ParcelNFT.contract.UnpackLog(event, "ParcelMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ParcelNFT contract.
type ParcelNFTTransferIterator struct {
	Event *ParcelNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTTransfer represents a Transfer event raised by the ParcelNFT contract.
type ParcelNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ParcelNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTTransferIterator{contract: _ParcelNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ParcelNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTTransfer)
				if err := _ParcelNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) ParseTransfer(log types.Log) (*ParcelNFTTransfer, error) {
	event := new(ParcelNFTTransfer)
	if err := _ParcelNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
