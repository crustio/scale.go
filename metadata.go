package scalecodec

import (
	"errors"

	"github.com/crustio/scale.go/types"
	"github.com/crustio/scale.go/types/scaleBytes"
	"github.com/crustio/scale.go/utiles"
)

type MetadataDecoder struct {
	types.ScaleDecoder
	Version  string               `json:"version"`
	Metadata types.MetadataStruct `json:"metadata"`
}

func (m *MetadataDecoder) Init(data []byte) {
	sData := scaleBytes.ScaleBytes{Data: data}
	m.ScaleDecoder.Init(sData, nil)
}

func (m *MetadataDecoder) Process() error {
	magicBytes := m.NextBytes(4)
	if string(magicBytes) == "meta" {
		metadataVersion := utiles.U256(utiles.BytesToHex(m.Data.Data[m.Data.Offset : m.Data.Offset+1]))
		m.Version = m.ProcessAndUpdateData("MetadataVersion").(string)
		m.Metadata = m.ProcessAndUpdateData(m.Version).(types.MetadataStruct)
		m.Metadata.MetadataVersion = int(metadataVersion.Int64())
		return nil
	}
	return errors.New("not metadata")

}
