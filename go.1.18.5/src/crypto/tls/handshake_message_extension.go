package tls

type IBuilder interface {
	AddUint16(v uint16)
}

type IHelloExtension interface {
	GetExtensionType() uint16
	GetExtensionData() []byte
}

type FetchNewExtensionsFunc func(sessionID []byte) []IHelloExtension
type OnDecryptConnRecordFunc func(record []byte, recordType uint8)

var (
	fetchServerNewExtensions FetchNewExtensionsFunc
	fetchClientNewExtensions FetchNewExtensionsFunc
	onDecryptConnRecord      OnDecryptConnRecordFunc
)

func SetFetchServerNewExtensionsFunc(f FetchNewExtensionsFunc) {
	fetchServerNewExtensions = f
}

func SetFetchClientNewExtensionsFunc(f FetchNewExtensionsFunc) {
	fetchClientNewExtensions = f
}

func SetOnDecryptConnRecordFunc(f OnDecryptConnRecordFunc) {
	onDecryptConnRecord = f
}
