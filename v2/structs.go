package v2

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type tonCommon struct {
	Type  string `json:"@type"`
	Extra string `json:"@extra"`
}

type SecureBytes []byte
type SecureString string
type Bytes []byte
type TvmStackEntry interface{}
type SmcMethodId int32
type TvmNumber string
type GenericAccountState string

// JSONInt64 alias for int64, in order to deal with json big number problem
type JSONInt64 int64

// MarshalJSON marshals to json
func (jsonInt *JSONInt64) MarshalJSON() ([]byte, error) {
	intStr := strconv.FormatInt(int64(*jsonInt), 10)
	return []byte(intStr), nil
}

// UnmarshalJSON unmarshals from json
func (jsonInt *JSONInt64) UnmarshalJSON(b []byte) error {
	intStr := string(b)
	intStr = strings.Replace(intStr, "\"", "", 2)
	jsonBigInt, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return err
	}
	*jsonInt = JSONInt64(jsonBigInt)
	return nil
}

// TonMessage is the interface for all messages send and received to/from tonlib
type TonMessage interface {
	MessageType() string
}

// LogStreamEnum Alias for abstract LogStream 'Sub-Classes', used as constant-enum here
type LogStreamEnum string

// LogStream enums
const (
	LogStreamDefaultType LogStreamEnum = "logStreamDefault"
	LogStreamFileType    LogStreamEnum = "logStreamFile"
	LogStreamEmptyType   LogStreamEnum = "logStreamEmpty"
) // LogStream Describes a stream to which tonlib internal log is written
type LogStream interface {
	GetLogStreamEnum() LogStreamEnum
}

// Double
type Double struct {
	tonCommon
}

// MessageType return the string telegram-type of Double
func (double *Double) MessageType() string {
	return "double"
}

// NewDouble creates a new Double
//
func NewDouble() *Double {
	doubleTemp := Double{
		tonCommon: tonCommon{Type: "double"},
	}

	return &doubleTemp
}

// String
type String struct {
	tonCommon
}

// MessageType return the string telegram-type of String
func (string *String) MessageType() string {
	return "string"
}

// NewString creates a new String
//
func NewString() *String {
	stringTemp := String{
		tonCommon: tonCommon{Type: "string"},
	}

	return &stringTemp
}

// Int32
type Int32 struct {
	tonCommon
}

// MessageType return the string telegram-type of Int32
func (int32 *Int32) MessageType() string {
	return "int32"
}

// NewInt32 creates a new Int32
//
func NewInt32() *Int32 {
	int32Temp := Int32{
		tonCommon: tonCommon{Type: "int32"},
	}

	return &int32Temp
}

// Int53
type Int53 struct {
	tonCommon
}

// MessageType return the string telegram-type of Int53
func (int53 *Int53) MessageType() string {
	return "int53"
}

// NewInt53 creates a new Int53
//
func NewInt53() *Int53 {
	int53Temp := Int53{
		tonCommon: tonCommon{Type: "int53"},
	}

	return &int53Temp
}

// Int64
type Int64 struct {
	tonCommon
}

// MessageType return the string telegram-type of Int64
func (int64 *Int64) MessageType() string {
	return "int64"
}

// NewInt64 creates a new Int64
//
func NewInt64() *Int64 {
	int64Temp := Int64{
		tonCommon: tonCommon{Type: "int64"},
	}

	return &int64Temp
}

// BoolFalse
type BoolFalse struct {
	tonCommon
}

// MessageType return the string telegram-type of BoolFalse
func (boolFalse *BoolFalse) MessageType() string {
	return "boolFalse"
}

// NewBoolFalse creates a new BoolFalse
//
func NewBoolFalse() *BoolFalse {
	boolFalseTemp := BoolFalse{
		tonCommon: tonCommon{Type: "boolFalse"},
	}

	return &boolFalseTemp
}

// BoolTrue
type BoolTrue struct {
	tonCommon
}

// MessageType return the string telegram-type of BoolTrue
func (boolTrue *BoolTrue) MessageType() string {
	return "boolTrue"
}

// NewBoolTrue creates a new BoolTrue
//
func NewBoolTrue() *BoolTrue {
	boolTrueTemp := BoolTrue{
		tonCommon: tonCommon{Type: "boolTrue"},
	}

	return &boolTrueTemp
}

// Error
type Error struct {
	tonCommon
	Code    int32  `json:"code"`    //
	Message string `json:"message"` //
}

// MessageType return the string telegram-type of Error
func (error *Error) MessageType() string {
	return "error"
}

// NewError creates a new Error
//
// @param code
// @param message
func NewError(code int32, message string) *Error {
	errorTemp := Error{
		tonCommon: tonCommon{Type: "error"},
		Code:      code,
		Message:   message,
	}

	return &errorTemp
}

// Ok
type Ok struct {
	tonCommon
}

// MessageType return the string telegram-type of Ok
func (ok *Ok) MessageType() string {
	return "ok"
}

// NewOk creates a new Ok
//
func NewOk() *Ok {
	okTemp := Ok{
		tonCommon: tonCommon{Type: "ok"},
	}

	return &okTemp
}

// KeyStoreTypeDirectory
type KeyStoreTypeDirectory struct {
	tonCommon
	Directory string `json:"directory"` //
}

// MessageType return the string telegram-type of KeyStoreTypeDirectory
func (keyStoreTypeDirectory *KeyStoreTypeDirectory) MessageType() string {
	return "keyStoreTypeDirectory"
}

// NewKeyStoreTypeDirectory creates a new KeyStoreTypeDirectory
//
// @param directory
func NewKeyStoreTypeDirectory(directory string) *KeyStoreTypeDirectory {
	keyStoreTypeDirectoryTemp := KeyStoreTypeDirectory{
		tonCommon: tonCommon{Type: "keyStoreTypeDirectory"},
		Directory: directory,
	}

	return &keyStoreTypeDirectoryTemp
}

// KeyStoreTypeInMemory
type KeyStoreTypeInMemory struct {
	tonCommon
}

// MessageType return the string telegram-type of KeyStoreTypeInMemory
func (keyStoreTypeInMemory *KeyStoreTypeInMemory) MessageType() string {
	return "keyStoreTypeInMemory"
}

// NewKeyStoreTypeInMemory creates a new KeyStoreTypeInMemory
//
func NewKeyStoreTypeInMemory() *KeyStoreTypeInMemory {
	keyStoreTypeInMemoryTemp := KeyStoreTypeInMemory{
		tonCommon: tonCommon{Type: "keyStoreTypeInMemory"},
	}

	return &keyStoreTypeInMemoryTemp
}

// Config
type Config struct {
	tonCommon
	Config                 string `json:"config"`                    //
	BlockchainName         string `json:"blockchain_name"`           //
	UseCallbacksForNetwork bool   `json:"use_callbacks_for_network"` //
	IgnoreCache            bool   `json:"ignore_cache"`              //
}

// MessageType return the string telegram-type of Config
func (config *Config) MessageType() string {
	return "config"
}

// NewConfig creates a new Config
//
// @param config
// @param blockchainName
// @param useCallbacksForNetwork
// @param ignoreCache
func NewConfig(config string, blockchainName string, useCallbacksForNetwork bool, ignoreCache bool) *Config {
	configTemp := Config{
		tonCommon:              tonCommon{Type: "config"},
		Config:                 config,
		BlockchainName:         blockchainName,
		UseCallbacksForNetwork: useCallbacksForNetwork,
		IgnoreCache:            ignoreCache,
	}

	return &configTemp
}

// Options
type Options struct {
	tonCommon
	Config       *Config       `json:"config"`        //
	KeystoreType *KeyStoreType `json:"keystore_type"` //
}

// MessageType return the string telegram-type of Options
func (options *Options) MessageType() string {
	return "options"
}

// NewOptions creates a new Options
//
// @param config
// @param keystoreType
func NewOptions(config *Config, keystoreType *KeyStoreType) *Options {
	optionsTemp := Options{
		tonCommon:    tonCommon{Type: "options"},
		Config:       config,
		KeystoreType: keystoreType,
	}

	return &optionsTemp
}

// OptionsConfigInfo
type OptionsConfigInfo struct {
	tonCommon
	DefaultWalletId int64 `json:"default_wallet_id"` //
}

// MessageType return the string telegram-type of OptionsConfigInfo
func (optionsConfigInfo *OptionsConfigInfo) MessageType() string {
	return "options.configInfo"
}

// NewOptionsConfigInfo creates a new OptionsConfigInfo
//
// @param defaultWalletId
func NewOptionsConfigInfo(defaultWalletId int64) *OptionsConfigInfo {
	optionsConfigInfoTemp := OptionsConfigInfo{
		tonCommon:       tonCommon{Type: "options.configInfo"},
		DefaultWalletId: defaultWalletId,
	}

	return &optionsConfigInfoTemp
}

// OptionsInfo
type OptionsInfo struct {
	tonCommon
	ConfigInfo *OptionsConfigInfo `json:"config_info"` //
}

// MessageType return the string telegram-type of OptionsInfo
func (optionsInfo *OptionsInfo) MessageType() string {
	return "options.info"
}

// NewOptionsInfo creates a new OptionsInfo
//
// @param configInfo
func NewOptionsInfo(configInfo *OptionsConfigInfo) *OptionsInfo {
	optionsInfoTemp := OptionsInfo{
		tonCommon:  tonCommon{Type: "options.info"},
		ConfigInfo: configInfo,
	}

	return &optionsInfoTemp
}

// Key
type Key struct {
	tonCommon
	PublicKey string       `json:"public_key"` //
	Secret    *SecureBytes `json:"secret"`     //
}

// MessageType return the string telegram-type of Key
func (key *Key) MessageType() string {
	return "key"
}

// NewKey creates a new Key
//
// @param publicKey
// @param secret
func NewKey(publicKey string, secret *SecureBytes) *Key {
	keyTemp := Key{
		tonCommon: tonCommon{Type: "key"},
		PublicKey: publicKey,
		Secret:    secret,
	}

	return &keyTemp
}

// InputKeyRegular
type InputKeyRegular struct {
	tonCommon
	Key           *Key         `json:"key"`            //
	LocalPassword *SecureBytes `json:"local_password"` //
}

// MessageType return the string telegram-type of InputKeyRegular
func (inputKeyRegular *InputKeyRegular) MessageType() string {
	return "inputKeyRegular"
}

// NewInputKeyRegular creates a new InputKeyRegular
//
// @param key
// @param localPassword
func NewInputKeyRegular(key *Key, localPassword *SecureBytes) *InputKeyRegular {
	inputKeyRegularTemp := InputKeyRegular{
		tonCommon:     tonCommon{Type: "inputKeyRegular"},
		Key:           key,
		LocalPassword: localPassword,
	}

	return &inputKeyRegularTemp
}

// InputKeyFake
type InputKeyFake struct {
	tonCommon
}

// MessageType return the string telegram-type of InputKeyFake
func (inputKeyFake *InputKeyFake) MessageType() string {
	return "inputKeyFake"
}

// NewInputKeyFake creates a new InputKeyFake
//
func NewInputKeyFake() *InputKeyFake {
	inputKeyFakeTemp := InputKeyFake{
		tonCommon: tonCommon{Type: "inputKeyFake"},
	}

	return &inputKeyFakeTemp
}

// ExportedKey
type ExportedKey struct {
	tonCommon
	WordList []SecureString `json:"word_list"` //
}

// MessageType return the string telegram-type of ExportedKey
func (exportedKey *ExportedKey) MessageType() string {
	return "exportedKey"
}

// NewExportedKey creates a new ExportedKey
//
// @param wordList
func NewExportedKey(wordList []SecureString) *ExportedKey {
	exportedKeyTemp := ExportedKey{
		tonCommon: tonCommon{Type: "exportedKey"},
		WordList:  wordList,
	}

	return &exportedKeyTemp
}

// ExportedPemKey
type ExportedPemKey struct {
	tonCommon
	Pem *SecureString `json:"pem"` //
}

// MessageType return the string telegram-type of ExportedPemKey
func (exportedPemKey *ExportedPemKey) MessageType() string {
	return "exportedPemKey"
}

// NewExportedPemKey creates a new ExportedPemKey
//
// @param pem
func NewExportedPemKey(pem *SecureString) *ExportedPemKey {
	exportedPemKeyTemp := ExportedPemKey{
		tonCommon: tonCommon{Type: "exportedPemKey"},
		Pem:       pem,
	}

	return &exportedPemKeyTemp
}

// ExportedEncryptedKey
type ExportedEncryptedKey struct {
	tonCommon
	Data *SecureBytes `json:"data"` //
}

// MessageType return the string telegram-type of ExportedEncryptedKey
func (exportedEncryptedKey *ExportedEncryptedKey) MessageType() string {
	return "exportedEncryptedKey"
}

// NewExportedEncryptedKey creates a new ExportedEncryptedKey
//
// @param data
func NewExportedEncryptedKey(data *SecureBytes) *ExportedEncryptedKey {
	exportedEncryptedKeyTemp := ExportedEncryptedKey{
		tonCommon: tonCommon{Type: "exportedEncryptedKey"},
		Data:      data,
	}

	return &exportedEncryptedKeyTemp
}

// Bip39Hints
type Bip39Hints struct {
	tonCommon
	Words []string `json:"words"` //
}

// MessageType return the string telegram-type of Bip39Hints
func (bip39Hints *Bip39Hints) MessageType() string {
	return "bip39Hints"
}

// NewBip39Hints creates a new Bip39Hints
//
// @param words
func NewBip39Hints(words []string) *Bip39Hints {
	bip39HintsTemp := Bip39Hints{
		tonCommon: tonCommon{Type: "bip39Hints"},
		Words:     words,
	}

	return &bip39HintsTemp
}

// AccountAddress
type AccountAddress struct {
	tonCommon
	AccountAddress string `json:"account_address"` //
}

// MessageType return the string telegram-type of AccountAddress
func (accountAddress *AccountAddress) MessageType() string {
	return "accountAddress"
}

// NewAccountAddress creates a new AccountAddress
//
// @param accountAddress
func NewAccountAddress(accountAddress string) *AccountAddress {
	accountAddressTemp := AccountAddress{
		tonCommon:      tonCommon{Type: "accountAddress"},
		AccountAddress: accountAddress,
	}

	return &accountAddressTemp
}

// UnpackedAccountAddress
type UnpackedAccountAddress struct {
	tonCommon
	WorkchainId int32  `json:"workchain_id"` //
	Bounceable  bool   `json:"bounceable"`   //
	Testnet     bool   `json:"testnet"`      //
	Addr        []byte `json:"addr"`         //
}

// MessageType return the string telegram-type of UnpackedAccountAddress
func (unpackedAccountAddress *UnpackedAccountAddress) MessageType() string {
	return "unpackedAccountAddress"
}

// NewUnpackedAccountAddress creates a new UnpackedAccountAddress
//
// @param workchainId
// @param bounceable
// @param testnet
// @param addr
func NewUnpackedAccountAddress(workchainId int32, bounceable bool, testnet bool, addr []byte) *UnpackedAccountAddress {
	unpackedAccountAddressTemp := UnpackedAccountAddress{
		tonCommon:   tonCommon{Type: "unpackedAccountAddress"},
		WorkchainId: workchainId,
		Bounceable:  bounceable,
		Testnet:     testnet,
		Addr:        addr,
	}

	return &unpackedAccountAddressTemp
}

// InternalTransactionId
type InternalTransactionId struct {
	tonCommon
	Lt   JSONInt64 `json:"lt"`   //
	Hash []byte    `json:"hash"` //
}

// MessageType return the string telegram-type of InternalTransactionId
func (internalTransactionId *InternalTransactionId) MessageType() string {
	return "internal.transactionId"
}

// NewInternalTransactionId creates a new InternalTransactionId
//
// @param lt
// @param hash
func NewInternalTransactionId(lt JSONInt64, hash []byte) *InternalTransactionId {
	internalTransactionIdTemp := InternalTransactionId{
		tonCommon: tonCommon{Type: "internal.transactionId"},
		Lt:        lt,
		Hash:      hash,
	}

	return &internalTransactionIdTemp
}

// RawInitialAccountState
type RawInitialAccountState struct {
	tonCommon
	Code []byte `json:"code"` //
	Data []byte `json:"data"` //
}

// MessageType return the string telegram-type of RawInitialAccountState
func (rawInitialAccountState *RawInitialAccountState) MessageType() string {
	return "raw.initialAccountState"
}

// NewRawInitialAccountState creates a new RawInitialAccountState
//
// @param code
// @param data
func NewRawInitialAccountState(code []byte, data []byte) *RawInitialAccountState {
	rawInitialAccountStateTemp := RawInitialAccountState{
		tonCommon: tonCommon{Type: "raw.initialAccountState"},
		Code:      code,
		Data:      data,
	}

	return &rawInitialAccountStateTemp
}

// RawAccountState
type RawAccountState struct {
	tonCommon
	Balance           JSONInt64              `json:"balance"`             //
	Code              []byte                 `json:"code"`                //
	Data              []byte                 `json:"data"`                //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	FrozenHash        []byte                 `json:"frozen_hash"`         //
	SyncUtime         int64                  `json:"sync_utime"`          //
}

// MessageType return the string telegram-type of RawAccountState
func (rawAccountState *RawAccountState) MessageType() string {
	return "raw.accountState"
}

// NewRawAccountState creates a new RawAccountState
//
// @param balance
// @param code
// @param data
// @param lastTransactionId
// @param frozenHash
// @param syncUtime
func NewRawAccountState(balance JSONInt64, code []byte, data []byte, lastTransactionId *InternalTransactionId, frozenHash []byte, syncUtime int64) *RawAccountState {
	rawAccountStateTemp := RawAccountState{
		tonCommon:         tonCommon{Type: "raw.accountState"},
		Balance:           balance,
		Code:              code,
		Data:              data,
		LastTransactionId: lastTransactionId,
		FrozenHash:        frozenHash,
		SyncUtime:         syncUtime,
	}

	return &rawAccountStateTemp
}

// RawMessage
type RawMessage struct {
	tonCommon
	CreatedLt   JSONInt64 `json:"created_lt"`  //
	BodyHash    []byte    `json:"body_hash"`   //
	Message     []byte    `json:"message"`     //
	Source      string    `json:"source"`      //
	Destination string    `json:"destination"` //
	Value       JSONInt64 `json:"value"`       //
	FwdFee      JSONInt64 `json:"fwd_fee"`     //
	IhrFee      JSONInt64 `json:"ihr_fee"`     //
}

// MessageType return the string telegram-type of RawMessage
func (rawMessage *RawMessage) MessageType() string {
	return "raw.message"
}

// NewRawMessage creates a new RawMessage
//
// @param createdLt
// @param bodyHash
// @param message
// @param source
// @param destination
// @param value
// @param fwdFee
// @param ihrFee
func NewRawMessage(createdLt JSONInt64, bodyHash []byte, message []byte, source string, destination string, value JSONInt64, fwdFee JSONInt64, ihrFee JSONInt64) *RawMessage {
	rawMessageTemp := RawMessage{
		tonCommon:   tonCommon{Type: "raw.message"},
		CreatedLt:   createdLt,
		BodyHash:    bodyHash,
		Message:     message,
		Source:      source,
		Destination: destination,
		Value:       value,
		FwdFee:      fwdFee,
		IhrFee:      ihrFee,
	}

	return &rawMessageTemp
}

// RawTransaction
type RawTransaction struct {
	tonCommon
	Data          []byte                 `json:"data"`           //
	TransactionId *InternalTransactionId `json:"transaction_id"` //
	Fee           JSONInt64              `json:"fee"`            //
	StorageFee    JSONInt64              `json:"storage_fee"`    //
	OtherFee      JSONInt64              `json:"other_fee"`      //
	InMsg         *RawMessage            `json:"in_msg"`         //
	OutMsgs       []RawMessage           `json:"out_msgs"`       //
	Utime         int64                  `json:"utime"`          //
}

// MessageType return the string telegram-type of RawTransaction
func (rawTransaction *RawTransaction) MessageType() string {
	return "raw.transaction"
}

// NewRawTransaction creates a new RawTransaction
//
// @param data
// @param transactionId
// @param fee
// @param storageFee
// @param otherFee
// @param inMsg
// @param outMsgs
// @param utime
func NewRawTransaction(data []byte, transactionId *InternalTransactionId, fee JSONInt64, storageFee JSONInt64, otherFee JSONInt64, inMsg *RawMessage, outMsgs []RawMessage, utime int64) *RawTransaction {
	rawTransactionTemp := RawTransaction{
		tonCommon:     tonCommon{Type: "raw.transaction"},
		Data:          data,
		TransactionId: transactionId,
		Fee:           fee,
		StorageFee:    storageFee,
		OtherFee:      otherFee,
		InMsg:         inMsg,
		OutMsgs:       outMsgs,
		Utime:         utime,
	}

	return &rawTransactionTemp
}

// RawTransactions
type RawTransactions struct {
	tonCommon
	Transactions          []RawTransaction       `json:"transactions"`            //
	PreviousTransactionId *InternalTransactionId `json:"previous_transaction_id"` //
}

// MessageType return the string telegram-type of RawTransactions
func (rawTransactions *RawTransactions) MessageType() string {
	return "raw.transactions"
}

// NewRawTransactions creates a new RawTransactions
//
// @param transactions
// @param previousTransactionId
func NewRawTransactions(transactions []RawTransaction, previousTransactionId *InternalTransactionId) *RawTransactions {
	rawTransactionsTemp := RawTransactions{
		tonCommon:             tonCommon{Type: "raw.transactions"},
		Transactions:          transactions,
		PreviousTransactionId: previousTransactionId,
	}

	return &rawTransactionsTemp
}

// TestWalletInitialAccountState
type TestWalletInitialAccountState struct {
	tonCommon
	PublicKey string `json:"public_key"` //
}

// MessageType return the string telegram-type of TestWalletInitialAccountState
func (testWalletInitialAccountState *TestWalletInitialAccountState) MessageType() string {
	return "testWallet.initialAccountState"
}

// NewTestWalletInitialAccountState creates a new TestWalletInitialAccountState
//
// @param publicKey
func NewTestWalletInitialAccountState(publicKey string) *TestWalletInitialAccountState {
	testWalletInitialAccountStateTemp := TestWalletInitialAccountState{
		tonCommon: tonCommon{Type: "testWallet.initialAccountState"},
		PublicKey: publicKey,
	}

	return &testWalletInitialAccountStateTemp
}

// TestWalletAccountState
type TestWalletAccountState struct {
	tonCommon
	Balance           JSONInt64              `json:"balance"`             //
	Seqno             int32                  `json:"seqno"`               //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	SyncUtime         int64                  `json:"sync_utime"`          //
}

// MessageType return the string telegram-type of TestWalletAccountState
func (testWalletAccountState *TestWalletAccountState) MessageType() string {
	return "testWallet.accountState"
}

// NewTestWalletAccountState creates a new TestWalletAccountState
//
// @param balance
// @param seqno
// @param lastTransactionId
// @param syncUtime
func NewTestWalletAccountState(balance JSONInt64, seqno int32, lastTransactionId *InternalTransactionId, syncUtime int64) *TestWalletAccountState {
	testWalletAccountStateTemp := TestWalletAccountState{
		tonCommon:         tonCommon{Type: "testWallet.accountState"},
		Balance:           balance,
		Seqno:             seqno,
		LastTransactionId: lastTransactionId,
		SyncUtime:         syncUtime,
	}

	return &testWalletAccountStateTemp
}

// WalletInitialAccountState
type WalletInitialAccountState struct {
	tonCommon
	PublicKey string `json:"public_key"` //
}

// MessageType return the string telegram-type of WalletInitialAccountState
func (walletInitialAccountState *WalletInitialAccountState) MessageType() string {
	return "wallet.initialAccountState"
}

// NewWalletInitialAccountState creates a new WalletInitialAccountState
//
// @param publicKey
func NewWalletInitialAccountState(publicKey string) *WalletInitialAccountState {
	walletInitialAccountStateTemp := WalletInitialAccountState{
		tonCommon: tonCommon{Type: "wallet.initialAccountState"},
		PublicKey: publicKey,
	}

	return &walletInitialAccountStateTemp
}

// WalletAccountState
type WalletAccountState struct {
	tonCommon
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	SyncUtime         int64                  `json:"sync_utime"`          //
	Balance           JSONInt64              `json:"balance"`             //
	Seqno             int32                  `json:"seqno"`               //
}

// MessageType return the string telegram-type of WalletAccountState
func (walletAccountState *WalletAccountState) MessageType() string {
	return "wallet.accountState"
}

// NewWalletAccountState creates a new WalletAccountState
//
// @param lastTransactionId
// @param syncUtime
// @param balance
// @param seqno
func NewWalletAccountState(lastTransactionId *InternalTransactionId, syncUtime int64, balance JSONInt64, seqno int32) *WalletAccountState {
	walletAccountStateTemp := WalletAccountState{
		tonCommon:         tonCommon{Type: "wallet.accountState"},
		LastTransactionId: lastTransactionId,
		SyncUtime:         syncUtime,
		Balance:           balance,
		Seqno:             seqno,
	}

	return &walletAccountStateTemp
}

// WalletV3InitialAccountState
type WalletV3InitialAccountState struct {
	tonCommon
	WalletId  int64  `json:"wallet_id"`  //
	PublicKey string `json:"public_key"` //
}

// MessageType return the string telegram-type of WalletV3InitialAccountState
func (walletV3InitialAccountState *WalletV3InitialAccountState) MessageType() string {
	return "wallet.v3.initialAccountState"
}

// NewWalletV3InitialAccountState creates a new WalletV3InitialAccountState
//
// @param walletId
// @param publicKey
func NewWalletV3InitialAccountState(walletId int64, publicKey string) *WalletV3InitialAccountState {
	walletV3InitialAccountStateTemp := WalletV3InitialAccountState{
		tonCommon: tonCommon{Type: "wallet.v3.initialAccountState"},
		WalletId:  walletId,
		PublicKey: publicKey,
	}

	return &walletV3InitialAccountStateTemp
}

// WalletV3AccountState
type WalletV3AccountState struct {
	tonCommon
	Seqno             int32                  `json:"seqno"`               //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	SyncUtime         int64                  `json:"sync_utime"`          //
	Balance           JSONInt64              `json:"balance"`             //
	WalletId          int64                  `json:"wallet_id"`           //
}

// MessageType return the string telegram-type of WalletV3AccountState
func (walletV3AccountState *WalletV3AccountState) MessageType() string {
	return "wallet.v3.accountState"
}

// NewWalletV3AccountState creates a new WalletV3AccountState
//
// @param seqno
// @param lastTransactionId
// @param syncUtime
// @param balance
// @param walletId
func NewWalletV3AccountState(seqno int32, lastTransactionId *InternalTransactionId, syncUtime int64, balance JSONInt64, walletId int64) *WalletV3AccountState {
	walletV3AccountStateTemp := WalletV3AccountState{
		tonCommon:         tonCommon{Type: "wallet.v3.accountState"},
		Seqno:             seqno,
		LastTransactionId: lastTransactionId,
		SyncUtime:         syncUtime,
		Balance:           balance,
		WalletId:          walletId,
	}

	return &walletV3AccountStateTemp
}

// TestGiverAccountState
type TestGiverAccountState struct {
	tonCommon
	Balance           JSONInt64              `json:"balance"`             //
	Seqno             int32                  `json:"seqno"`               //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	SyncUtime         int64                  `json:"sync_utime"`          //
}

// MessageType return the string telegram-type of TestGiverAccountState
func (testGiverAccountState *TestGiverAccountState) MessageType() string {
	return "testGiver.accountState"
}

// NewTestGiverAccountState creates a new TestGiverAccountState
//
// @param balance
// @param seqno
// @param lastTransactionId
// @param syncUtime
func NewTestGiverAccountState(balance JSONInt64, seqno int32, lastTransactionId *InternalTransactionId, syncUtime int64) *TestGiverAccountState {
	testGiverAccountStateTemp := TestGiverAccountState{
		tonCommon:         tonCommon{Type: "testGiver.accountState"},
		Balance:           balance,
		Seqno:             seqno,
		LastTransactionId: lastTransactionId,
		SyncUtime:         syncUtime,
	}

	return &testGiverAccountStateTemp
}

// UninitedAccountState
type UninitedAccountState struct {
	tonCommon
	SyncUtime         int64                  `json:"sync_utime"`          //
	Balance           JSONInt64              `json:"balance"`             //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	FrozenHash        []byte                 `json:"frozen_hash"`         //
}

// MessageType return the string telegram-type of UninitedAccountState
func (uninitedAccountState *UninitedAccountState) MessageType() string {
	return "uninited.accountState"
}

// NewUninitedAccountState creates a new UninitedAccountState
//
// @param syncUtime
// @param balance
// @param lastTransactionId
// @param frozenHash
func NewUninitedAccountState(syncUtime int64, balance JSONInt64, lastTransactionId *InternalTransactionId, frozenHash []byte) *UninitedAccountState {
	uninitedAccountStateTemp := UninitedAccountState{
		tonCommon:         tonCommon{Type: "uninited.accountState"},
		SyncUtime:         syncUtime,
		Balance:           balance,
		LastTransactionId: lastTransactionId,
		FrozenHash:        frozenHash,
	}

	return &uninitedAccountStateTemp
}

// GenericAccountStateRaw
type GenericAccountStateRaw struct {
	tonCommon
	AccountState *RawAccountState `json:"account_state"` //
}

// MessageType return the string telegram-type of GenericAccountStateRaw
func (genericAccountStateRaw *GenericAccountStateRaw) MessageType() string {
	return "generic.accountStateRaw"
}

// NewGenericAccountStateRaw creates a new GenericAccountStateRaw
//
// @param accountState
func NewGenericAccountStateRaw(accountState *RawAccountState) *GenericAccountStateRaw {
	genericAccountStateRawTemp := GenericAccountStateRaw{
		tonCommon:    tonCommon{Type: "generic.accountStateRaw"},
		AccountState: accountState,
	}

	return &genericAccountStateRawTemp
}

// GenericAccountStateTestWallet
type GenericAccountStateTestWallet struct {
	tonCommon
	AccountState *TestWalletAccountState `json:"account_state"` //
}

// MessageType return the string telegram-type of GenericAccountStateTestWallet
func (genericAccountStateTestWallet *GenericAccountStateTestWallet) MessageType() string {
	return "generic.accountStateTestWallet"
}

// NewGenericAccountStateTestWallet creates a new GenericAccountStateTestWallet
//
// @param accountState
func NewGenericAccountStateTestWallet(accountState *TestWalletAccountState) *GenericAccountStateTestWallet {
	genericAccountStateTestWalletTemp := GenericAccountStateTestWallet{
		tonCommon:    tonCommon{Type: "generic.accountStateTestWallet"},
		AccountState: accountState,
	}

	return &genericAccountStateTestWalletTemp
}

// GenericAccountStateWallet
type GenericAccountStateWallet struct {
	tonCommon
	AccountState *WalletAccountState `json:"account_state"` //
}

// MessageType return the string telegram-type of GenericAccountStateWallet
func (genericAccountStateWallet *GenericAccountStateWallet) MessageType() string {
	return "generic.accountStateWallet"
}

// NewGenericAccountStateWallet creates a new GenericAccountStateWallet
//
// @param accountState
func NewGenericAccountStateWallet(accountState *WalletAccountState) *GenericAccountStateWallet {
	genericAccountStateWalletTemp := GenericAccountStateWallet{
		tonCommon:    tonCommon{Type: "generic.accountStateWallet"},
		AccountState: accountState,
	}

	return &genericAccountStateWalletTemp
}

// GenericAccountStateWalletV3
type GenericAccountStateWalletV3 struct {
	tonCommon
	AccountState *WalletV3AccountState `json:"account_state"` //
}

// MessageType return the string telegram-type of GenericAccountStateWalletV3
func (genericAccountStateWalletV3 *GenericAccountStateWalletV3) MessageType() string {
	return "generic.accountStateWalletV3"
}

// NewGenericAccountStateWalletV3 creates a new GenericAccountStateWalletV3
//
// @param accountState
func NewGenericAccountStateWalletV3(accountState *WalletV3AccountState) *GenericAccountStateWalletV3 {
	genericAccountStateWalletV3Temp := GenericAccountStateWalletV3{
		tonCommon:    tonCommon{Type: "generic.accountStateWalletV3"},
		AccountState: accountState,
	}

	return &genericAccountStateWalletV3Temp
}

// GenericAccountStateTestGiver
type GenericAccountStateTestGiver struct {
	tonCommon
	AccountState *TestGiverAccountState `json:"account_state"` //
}

// MessageType return the string telegram-type of GenericAccountStateTestGiver
func (genericAccountStateTestGiver *GenericAccountStateTestGiver) MessageType() string {
	return "generic.accountStateTestGiver"
}

// NewGenericAccountStateTestGiver creates a new GenericAccountStateTestGiver
//
// @param accountState
func NewGenericAccountStateTestGiver(accountState *TestGiverAccountState) *GenericAccountStateTestGiver {
	genericAccountStateTestGiverTemp := GenericAccountStateTestGiver{
		tonCommon:    tonCommon{Type: "generic.accountStateTestGiver"},
		AccountState: accountState,
	}

	return &genericAccountStateTestGiverTemp
}

// GenericAccountStateUninited
type GenericAccountStateUninited struct {
	tonCommon
	AccountState *UninitedAccountState `json:"account_state"` //
}

// MessageType return the string telegram-type of GenericAccountStateUninited
func (genericAccountStateUninited *GenericAccountStateUninited) MessageType() string {
	return "generic.accountStateUninited"
}

// NewGenericAccountStateUninited creates a new GenericAccountStateUninited
//
// @param accountState
func NewGenericAccountStateUninited(accountState *UninitedAccountState) *GenericAccountStateUninited {
	genericAccountStateUninitedTemp := GenericAccountStateUninited{
		tonCommon:    tonCommon{Type: "generic.accountStateUninited"},
		AccountState: accountState,
	}

	return &genericAccountStateUninitedTemp
}

// SendGramsResult
type SendGramsResult struct {
	tonCommon
	SentUntil int64  `json:"sent_until"` //
	BodyHash  []byte `json:"body_hash"`  //
}

// MessageType return the string telegram-type of SendGramsResult
func (sendGramsResult *SendGramsResult) MessageType() string {
	return "sendGramsResult"
}

// NewSendGramsResult creates a new SendGramsResult
//
// @param sentUntil
// @param bodyHash
func NewSendGramsResult(sentUntil int64, bodyHash []byte) *SendGramsResult {
	sendGramsResultTemp := SendGramsResult{
		tonCommon: tonCommon{Type: "sendGramsResult"},
		SentUntil: sentUntil,
		BodyHash:  bodyHash,
	}

	return &sendGramsResultTemp
}

// SyncStateDone
type SyncStateDone struct {
	tonCommon
}

// MessageType return the string telegram-type of SyncStateDone
func (syncStateDone *SyncStateDone) MessageType() string {
	return "syncStateDone"
}

// NewSyncStateDone creates a new SyncStateDone
//
func NewSyncStateDone() *SyncStateDone {
	syncStateDoneTemp := SyncStateDone{
		tonCommon: tonCommon{Type: "syncStateDone"},
	}

	return &syncStateDoneTemp
}

// SyncStateInProgress
type SyncStateInProgress struct {
	tonCommon
	ToSeqno      int32 `json:"to_seqno"`      //
	CurrentSeqno int32 `json:"current_seqno"` //
	FromSeqno    int32 `json:"from_seqno"`    //
}

// MessageType return the string telegram-type of SyncStateInProgress
func (syncStateInProgress *SyncStateInProgress) MessageType() string {
	return "syncStateInProgress"
}

// NewSyncStateInProgress creates a new SyncStateInProgress
//
// @param toSeqno
// @param currentSeqno
// @param fromSeqno
func NewSyncStateInProgress(toSeqno int32, currentSeqno int32, fromSeqno int32) *SyncStateInProgress {
	syncStateInProgressTemp := SyncStateInProgress{
		tonCommon:    tonCommon{Type: "syncStateInProgress"},
		ToSeqno:      toSeqno,
		CurrentSeqno: currentSeqno,
		FromSeqno:    fromSeqno,
	}

	return &syncStateInProgressTemp
}

// Fees
type Fees struct {
	tonCommon
	InFwdFee   int64 `json:"in_fwd_fee"`  //
	StorageFee int64 `json:"storage_fee"` //
	GasFee     int64 `json:"gas_fee"`     //
	FwdFee     int64 `json:"fwd_fee"`     //
}

// MessageType return the string telegram-type of Fees
func (fees *Fees) MessageType() string {
	return "fees"
}

// NewFees creates a new Fees
//
// @param inFwdFee
// @param storageFee
// @param gasFee
// @param fwdFee
func NewFees(inFwdFee int64, storageFee int64, gasFee int64, fwdFee int64) *Fees {
	feesTemp := Fees{
		tonCommon:  tonCommon{Type: "fees"},
		InFwdFee:   inFwdFee,
		StorageFee: storageFee,
		GasFee:     gasFee,
		FwdFee:     fwdFee,
	}

	return &feesTemp
}

// QueryFees
type QueryFees struct {
	tonCommon
	SourceFees      *Fees `json:"source_fees"`      //
	DestinationFees *Fees `json:"destination_fees"` //
}

// MessageType return the string telegram-type of QueryFees
func (queryFees *QueryFees) MessageType() string {
	return "query.fees"
}

// NewQueryFees creates a new QueryFees
//
// @param sourceFees
// @param destinationFees
func NewQueryFees(sourceFees *Fees, destinationFees *Fees) *QueryFees {
	queryFeesTemp := QueryFees{
		tonCommon:       tonCommon{Type: "query.fees"},
		SourceFees:      sourceFees,
		DestinationFees: destinationFees,
	}

	return &queryFeesTemp
}

// QueryInfo
type QueryInfo struct {
	tonCommon
	Id         int64  `json:"id"`          //
	ValidUntil int64  `json:"valid_until"` //
	BodyHash   []byte `json:"body_hash"`   //
}

// MessageType return the string telegram-type of QueryInfo
func (queryInfo *QueryInfo) MessageType() string {
	return "query.info"
}

// NewQueryInfo creates a new QueryInfo
//
// @param id
// @param validUntil
// @param bodyHash
func NewQueryInfo(id int64, validUntil int64, bodyHash []byte) *QueryInfo {
	queryInfoTemp := QueryInfo{
		tonCommon:  tonCommon{Type: "query.info"},
		Id:         id,
		ValidUntil: validUntil,
		BodyHash:   bodyHash,
	}

	return &queryInfoTemp
}

// TvmSlice
type TvmSlice struct {
	tonCommon
	Bytes []byte `json:"bytes"` //
}

// MessageType return the string telegram-type of TvmSlice
func (tvmSlice *TvmSlice) MessageType() string {
	return "tvm.slice"
}

// NewTvmSlice creates a new TvmSlice
//
// @param bytes
func NewTvmSlice(bytes []byte) *TvmSlice {
	tvmSliceTemp := TvmSlice{
		tonCommon: tonCommon{Type: "tvm.slice"},
		Bytes:     bytes,
	}

	return &tvmSliceTemp
}

// TvmCell
type TvmCell struct {
	tonCommon
	Bytes []byte `json:"bytes"` //
}

// MessageType return the string telegram-type of TvmCell
func (tvmCell *TvmCell) MessageType() string {
	return "tvm.cell"
}

// NewTvmCell creates a new TvmCell
//
// @param bytes
func NewTvmCell(bytes []byte) *TvmCell {
	tvmCellTemp := TvmCell{
		tonCommon: tonCommon{Type: "tvm.cell"},
		Bytes:     bytes,
	}

	return &tvmCellTemp
}

// TvmNumberDecimal
type TvmNumberDecimal struct {
	tonCommon
	Number string `json:"number"` //
}

// MessageType return the string telegram-type of TvmNumberDecimal
func (tvmNumberDecimal *TvmNumberDecimal) MessageType() string {
	return "tvm.numberDecimal"
}

// NewTvmNumberDecimal creates a new TvmNumberDecimal
//
// @param number
func NewTvmNumberDecimal(number string) *TvmNumberDecimal {
	tvmNumberDecimalTemp := TvmNumberDecimal{
		tonCommon: tonCommon{Type: "tvm.numberDecimal"},
		Number:    number,
	}

	return &tvmNumberDecimalTemp
}

// TvmTuple
type TvmTuple struct {
	tonCommon
	Elements []TvmStackEntry `json:"elements"` //
}

// MessageType return the string telegram-type of TvmTuple
func (tvmTuple *TvmTuple) MessageType() string {
	return "tvm.tuple"
}

// NewTvmTuple creates a new TvmTuple
//
// @param elements
func NewTvmTuple(elements []TvmStackEntry) *TvmTuple {
	tvmTupleTemp := TvmTuple{
		tonCommon: tonCommon{Type: "tvm.tuple"},
		Elements:  elements,
	}

	return &tvmTupleTemp
}

// TvmList
type TvmList struct {
	tonCommon
	Elements []TvmStackEntry `json:"elements"` //
}

// MessageType return the string telegram-type of TvmList
func (tvmList *TvmList) MessageType() string {
	return "tvm.list"
}

// NewTvmList creates a new TvmList
//
// @param elements
func NewTvmList(elements []TvmStackEntry) *TvmList {
	tvmListTemp := TvmList{
		tonCommon: tonCommon{Type: "tvm.list"},
		Elements:  elements,
	}

	return &tvmListTemp
}

// TvmStackEntrySlice
type TvmStackEntrySlice struct {
	tonCommon
	Slice *TvmSlice `json:"slice"` //
}

// MessageType return the string telegram-type of TvmStackEntrySlice
func (tvmStackEntrySlice *TvmStackEntrySlice) MessageType() string {
	return "tvm.stackEntrySlice"
}

// NewTvmStackEntrySlice creates a new TvmStackEntrySlice
//
// @param slice
func NewTvmStackEntrySlice(slice *TvmSlice) *TvmStackEntrySlice {
	tvmStackEntrySliceTemp := TvmStackEntrySlice{
		tonCommon: tonCommon{Type: "tvm.stackEntrySlice"},
		Slice:     slice,
	}

	return &tvmStackEntrySliceTemp
}

// TvmStackEntryCell
type TvmStackEntryCell struct {
	tonCommon
	Cell *TvmCell `json:"cell"` //
}

// MessageType return the string telegram-type of TvmStackEntryCell
func (tvmStackEntryCell *TvmStackEntryCell) MessageType() string {
	return "tvm.stackEntryCell"
}

// NewTvmStackEntryCell creates a new TvmStackEntryCell
//
// @param cell
func NewTvmStackEntryCell(cell *TvmCell) *TvmStackEntryCell {
	tvmStackEntryCellTemp := TvmStackEntryCell{
		tonCommon: tonCommon{Type: "tvm.stackEntryCell"},
		Cell:      cell,
	}

	return &tvmStackEntryCellTemp
}

// TvmStackEntryNumber
type TvmStackEntryNumber struct {
	tonCommon
	Number *TvmNumber `json:"number"` //
}

// MessageType return the string telegram-type of TvmStackEntryNumber
func (tvmStackEntryNumber *TvmStackEntryNumber) MessageType() string {
	return "tvm.stackEntryNumber"
}

// NewTvmStackEntryNumber creates a new TvmStackEntryNumber
//
// @param number
func NewTvmStackEntryNumber(number *TvmNumber) *TvmStackEntryNumber {
	tvmStackEntryNumberTemp := TvmStackEntryNumber{
		tonCommon: tonCommon{Type: "tvm.stackEntryNumber"},
		Number:    number,
	}

	return &tvmStackEntryNumberTemp
}

// TvmStackEntryTuple
type TvmStackEntryTuple struct {
	tonCommon
	Tuple *TvmTuple `json:"tuple"` //
}

// MessageType return the string telegram-type of TvmStackEntryTuple
func (tvmStackEntryTuple *TvmStackEntryTuple) MessageType() string {
	return "tvm.stackEntryTuple"
}

// NewTvmStackEntryTuple creates a new TvmStackEntryTuple
//
// @param tuple
func NewTvmStackEntryTuple(tuple *TvmTuple) *TvmStackEntryTuple {
	tvmStackEntryTupleTemp := TvmStackEntryTuple{
		tonCommon: tonCommon{Type: "tvm.stackEntryTuple"},
		Tuple:     tuple,
	}

	return &tvmStackEntryTupleTemp
}

// TvmStackEntryList
type TvmStackEntryList struct {
	tonCommon
	List *TvmList `json:"list"` //
}

// MessageType return the string telegram-type of TvmStackEntryList
func (tvmStackEntryList *TvmStackEntryList) MessageType() string {
	return "tvm.stackEntryList"
}

// NewTvmStackEntryList creates a new TvmStackEntryList
//
// @param list
func NewTvmStackEntryList(list *TvmList) *TvmStackEntryList {
	tvmStackEntryListTemp := TvmStackEntryList{
		tonCommon: tonCommon{Type: "tvm.stackEntryList"},
		List:      list,
	}

	return &tvmStackEntryListTemp
}

// TvmStackEntryUnsupported
type TvmStackEntryUnsupported struct {
	tonCommon
}

// MessageType return the string telegram-type of TvmStackEntryUnsupported
func (tvmStackEntryUnsupported *TvmStackEntryUnsupported) MessageType() string {
	return "tvm.stackEntryUnsupported"
}

// NewTvmStackEntryUnsupported creates a new TvmStackEntryUnsupported
//
func NewTvmStackEntryUnsupported() *TvmStackEntryUnsupported {
	tvmStackEntryUnsupportedTemp := TvmStackEntryUnsupported{
		tonCommon: tonCommon{Type: "tvm.stackEntryUnsupported"},
	}

	return &tvmStackEntryUnsupportedTemp
}

// SmcInfo
type SmcInfo struct {
	tonCommon
	Id int64 `json:"id"` //
}

// MessageType return the string telegram-type of SmcInfo
func (smcInfo *SmcInfo) MessageType() string {
	return "smc.info"
}

// NewSmcInfo creates a new SmcInfo
//
// @param id
func NewSmcInfo(id int64) *SmcInfo {
	smcInfoTemp := SmcInfo{
		tonCommon: tonCommon{Type: "smc.info"},
		Id:        id,
	}

	return &smcInfoTemp
}

// SmcMethodIdNumber
type SmcMethodIdNumber struct {
	tonCommon
	Number int32 `json:"number"` //
}

// MessageType return the string telegram-type of SmcMethodIdNumber
func (smcMethodIdNumber *SmcMethodIdNumber) MessageType() string {
	return "smc.methodIdNumber"
}

// NewSmcMethodIdNumber creates a new SmcMethodIdNumber
//
// @param number
func NewSmcMethodIdNumber(number int32) *SmcMethodIdNumber {
	smcMethodIdNumberTemp := SmcMethodIdNumber{
		tonCommon: tonCommon{Type: "smc.methodIdNumber"},
		Number:    number,
	}

	return &smcMethodIdNumberTemp
}

// SmcMethodIdName
type SmcMethodIdName struct {
	tonCommon
	Name string `json:"name"` //
}

// MessageType return the string telegram-type of SmcMethodIdName
func (smcMethodIdName *SmcMethodIdName) MessageType() string {
	return "smc.methodIdName"
}

// NewSmcMethodIdName creates a new SmcMethodIdName
//
// @param name
func NewSmcMethodIdName(name string) *SmcMethodIdName {
	smcMethodIdNameTemp := SmcMethodIdName{
		tonCommon: tonCommon{Type: "smc.methodIdName"},
		Name:      name,
	}

	return &smcMethodIdNameTemp
}

// SmcRunResult
type SmcRunResult struct {
	tonCommon
	ExitCode int32           `json:"exit_code"` //
	GasUsed  int64           `json:"gas_used"`  //
	Stack    []TvmStackEntry `json:"stack"`     //
}

// MessageType return the string telegram-type of SmcRunResult
func (smcRunResult *SmcRunResult) MessageType() string {
	return "smc.runResult"
}

// NewSmcRunResult creates a new SmcRunResult
//
// @param exitCode
// @param gasUsed
// @param stack
func NewSmcRunResult(exitCode int32, gasUsed int64, stack []TvmStackEntry) *SmcRunResult {
	smcRunResultTemp := SmcRunResult{
		tonCommon: tonCommon{Type: "smc.runResult"},
		ExitCode:  exitCode,
		GasUsed:   gasUsed,
		Stack:     stack,
	}

	return &smcRunResultTemp
}

// UpdateSendLiteServerQuery
type UpdateSendLiteServerQuery struct {
	tonCommon
	Id   JSONInt64 `json:"id"`   //
	Data []byte    `json:"data"` //
}

// MessageType return the string telegram-type of UpdateSendLiteServerQuery
func (updateSendLiteServerQuery *UpdateSendLiteServerQuery) MessageType() string {
	return "updateSendLiteServerQuery"
}

// NewUpdateSendLiteServerQuery creates a new UpdateSendLiteServerQuery
//
// @param id
// @param data
func NewUpdateSendLiteServerQuery(id JSONInt64, data []byte) *UpdateSendLiteServerQuery {
	updateSendLiteServerQueryTemp := UpdateSendLiteServerQuery{
		tonCommon: tonCommon{Type: "updateSendLiteServerQuery"},
		Id:        id,
		Data:      data,
	}

	return &updateSendLiteServerQueryTemp
}

// UpdateSyncState
type UpdateSyncState struct {
	tonCommon
	SyncState *SyncState `json:"sync_state"` //
}

// MessageType return the string telegram-type of UpdateSyncState
func (updateSyncState *UpdateSyncState) MessageType() string {
	return "updateSyncState"
}

// NewUpdateSyncState creates a new UpdateSyncState
//
// @param syncState
func NewUpdateSyncState(syncState *SyncState) *UpdateSyncState {
	updateSyncStateTemp := UpdateSyncState{
		tonCommon: tonCommon{Type: "updateSyncState"},
		SyncState: syncState,
	}

	return &updateSyncStateTemp
}

// LogStreamDefault The log is written to stderr or an OS specific log
type LogStreamDefault struct {
	tonCommon
}

// MessageType return the string telegram-type of LogStreamDefault
func (logStreamDefault *LogStreamDefault) MessageType() string {
	return "logStreamDefault"
}

// NewLogStreamDefault creates a new LogStreamDefault
//
func NewLogStreamDefault() *LogStreamDefault {
	logStreamDefaultTemp := LogStreamDefault{
		tonCommon: tonCommon{Type: "logStreamDefault"},
	}

	return &logStreamDefaultTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamDefault *LogStreamDefault) GetLogStreamEnum() LogStreamEnum {
	return LogStreamDefaultType
}

// LogStreamFile The log is written to a file
type LogStreamFile struct {
	tonCommon
	Path        string `json:"path"`          // Path to the file to where the internal tonlib log will be written
	MaxFileSize int64  `json:"max_file_size"` // Maximum size of the file to where the internal tonlib log is written before the file will be auto-rotated
}

// MessageType return the string telegram-type of LogStreamFile
func (logStreamFile *LogStreamFile) MessageType() string {
	return "logStreamFile"
}

// NewLogStreamFile creates a new LogStreamFile
//
// @param path Path to the file to where the internal tonlib log will be written
// @param maxFileSize Maximum size of the file to where the internal tonlib log is written before the file will be auto-rotated
func NewLogStreamFile(path string, maxFileSize int64) *LogStreamFile {
	logStreamFileTemp := LogStreamFile{
		tonCommon:   tonCommon{Type: "logStreamFile"},
		Path:        path,
		MaxFileSize: maxFileSize,
	}

	return &logStreamFileTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamFile *LogStreamFile) GetLogStreamEnum() LogStreamEnum {
	return LogStreamFileType
}

// LogStreamEmpty The log is written nowhere
type LogStreamEmpty struct {
	tonCommon
}

// MessageType return the string telegram-type of LogStreamEmpty
func (logStreamEmpty *LogStreamEmpty) MessageType() string {
	return "logStreamEmpty"
}

// NewLogStreamEmpty creates a new LogStreamEmpty
//
func NewLogStreamEmpty() *LogStreamEmpty {
	logStreamEmptyTemp := LogStreamEmpty{
		tonCommon: tonCommon{Type: "logStreamEmpty"},
	}

	return &logStreamEmptyTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamEmpty *LogStreamEmpty) GetLogStreamEnum() LogStreamEnum {
	return LogStreamEmptyType
}

// LogVerbosityLevel Contains a tonlib internal log verbosity level
type LogVerbosityLevel struct {
	tonCommon
	VerbosityLevel int32 `json:"verbosity_level"` // Log verbosity level
}

// MessageType return the string telegram-type of LogVerbosityLevel
func (logVerbosityLevel *LogVerbosityLevel) MessageType() string {
	return "logVerbosityLevel"
}

// NewLogVerbosityLevel creates a new LogVerbosityLevel
//
// @param verbosityLevel Log verbosity level
func NewLogVerbosityLevel(verbosityLevel int32) *LogVerbosityLevel {
	logVerbosityLevelTemp := LogVerbosityLevel{
		tonCommon:      tonCommon{Type: "logVerbosityLevel"},
		VerbosityLevel: verbosityLevel,
	}

	return &logVerbosityLevelTemp
}

// LogTags Contains a list of available tonlib internal log tags
type LogTags struct {
	tonCommon
	Tags []string `json:"tags"` // List of log tags
}

// MessageType return the string telegram-type of LogTags
func (logTags *LogTags) MessageType() string {
	return "logTags"
}

// NewLogTags creates a new LogTags
//
// @param tags List of log tags
func NewLogTags(tags []string) *LogTags {
	logTagsTemp := LogTags{
		tonCommon: tonCommon{Type: "logTags"},
		Tags:      tags,
	}

	return &logTagsTemp
}

// Data
type Data struct {
	tonCommon
	Bytes *SecureBytes `json:"bytes"` //
}

// MessageType return the string telegram-type of Data
func (data *Data) MessageType() string {
	return "data"
}

// NewData creates a new Data
//
// @param bytes
func NewData(bytes *SecureBytes) *Data {
	dataTemp := Data{
		tonCommon: tonCommon{Type: "data"},
		Bytes:     bytes,
	}

	return &dataTemp
}

// LiteServerInfo
type LiteServerInfo struct {
	tonCommon
	Now          int64     `json:"now"`          //
	Version      int32     `json:"version"`      //
	Capabilities JSONInt64 `json:"capabilities"` //
}

// MessageType return the string telegram-type of LiteServerInfo
func (liteServerInfo *LiteServerInfo) MessageType() string {
	return "liteServer.info"
}

// NewLiteServerInfo creates a new LiteServerInfo
//
// @param now
// @param version
// @param capabilities
func NewLiteServerInfo(now int64, version int32, capabilities JSONInt64) *LiteServerInfo {
	liteServerInfoTemp := LiteServerInfo{
		tonCommon:    tonCommon{Type: "liteServer.info"},
		Now:          now,
		Version:      version,
		Capabilities: capabilities,
	}

	return &liteServerInfoTemp
}

func unmarshalLogStream(rawMsg *json.RawMessage) (LogStream, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LogStreamEnum(objMap["@type"].(string)) {
	case LogStreamDefaultType:
		var logStreamDefault LogStreamDefault
		err := json.Unmarshal(*rawMsg, &logStreamDefault)
		return &logStreamDefault, err

	case LogStreamFileType:
		var logStreamFile LogStreamFile
		err := json.Unmarshal(*rawMsg, &logStreamFile)
		return &logStreamFile, err

	case LogStreamEmptyType:
		var logStreamEmpty LogStreamEmpty
		err := json.Unmarshal(*rawMsg, &logStreamEmpty)
		return &logStreamEmpty, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}
