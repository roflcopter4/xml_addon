package reader

import (
	XMLtypes "github.com/lestrrat-go/libxml2/types"
	XMLxsd "github.com/lestrrat-go/libxml2/xsd"

	myXMLclib "github.com/roflcopter4/xml_addon/clib"
)

// FIXME These names are bad
const (
	Reader_None                  = myXMLclib.Reader_None
	Reader_Element               = myXMLclib.Reader_Element
	Reader_Attribute             = myXMLclib.Reader_Attribute
	Reader_Text                  = myXMLclib.Reader_Text
	Reader_Cdata                 = myXMLclib.Reader_Cdata
	Reader_EntityReference       = myXMLclib.Reader_EntityReference
	Reader_Entity                = myXMLclib.Reader_Entity
	Reader_ProcessingInstruction = myXMLclib.Reader_ProcessingInstruction
	Reader_Comment               = myXMLclib.Reader_Comment
	Reader_Document              = myXMLclib.Reader_Document
	Reader_DocumentType          = myXMLclib.Reader_DocumentType
	Reader_DocumentFragment      = myXMLclib.Reader_DocumentFragment
	Reader_Notation              = myXMLclib.Reader_Notation
	Reader_Whitespace            = myXMLclib.Reader_Whitespace
	Reader_SignificantWhitespace = myXMLclib.Reader_SignificantWhitespace
	Reader_EndElement            = myXMLclib.Reader_EndElement
	Reader_EndEntity             = myXMLclib.Reader_EndEntity
	Reader_XmlDeclaration        = myXMLclib.Reader_XmlDeclaration
)

const (
	ParserLoadDTD       = myXMLclib.ParserLoadDTD
	ParserDefaultAttrs  = myXMLclib.ParserDefaultAttrs
	ParserValidate      = myXMLclib.ParserValidate
	ParserSubstEntities = myXMLclib.ParserSubstEntities
)

type XMLTextReader struct {
	ptr    uintptr // *C.xmlTextReader
	mortal bool    // I have no clue what purpose this serves but c'est la vie
}

type TextReader interface {
	XMLtypes.PtrSource
	TextRead() int
	NodeType() myXMLclib.XMLReaderType
	Name() string
	Value() string
	AttributeCount() int
	Depth() int
	MoveToAttributeNo(int) error
	MoveToElement() error
	CurrentNode() (XMLtypes.Node, error)
	SetParserProp(prop myXMLclib.XMLParserProperties, value int) error
	SetSchema(*XMLxsd.Schema) error
}
