package clib

// The C library uses the type "xmlReaderTypes" for this enumeration
type XMLReaderType int

// FIXME These names are bad
const (
	Reader_None = iota
	Reader_Element
	Reader_Attribute
	Reader_Text
	Reader_Cdata
	Reader_EntityReference
	Reader_Entity
	Reader_ProcessingInstruction
	Reader_Comment
	Reader_Document
	Reader_DocumentType
	Reader_DocumentFragment
	Reader_Notation
	Reader_Whitespace
	Reader_SignificantWhitespace
	Reader_EndElement
	Reader_EndEntity
	Reader_XmlDeclaration
)

type XMLParserProperties int

const (
	ParserLoadDTD       = 1
	ParserDefaultAttrs  = 2
	ParserValidate      = 3
	ParserSubstEntities = 4
)
