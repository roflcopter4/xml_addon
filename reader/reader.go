/*
 * Theoretically this package ought to implement the xmlTextReader interface
 * found in "xmlreader.h". It allows one to walk through an XML document line by
 * line, with somewhat more information on the actual structure of the document.
 * Iteration will include both start and end tags in the order they appear, in
 * addition to all whitespace in the document.
 */
package reader

import (
	// "fmt"
	// "unsafe"
	// "bufio"
	// "os"
	// "path/filepath"
	// XML "github.com/lestrrat-go/libxml2"
	// XMLclib "github.com/lestrrat-go/libxml2/clib"
	// "github.com/pkg/errors"
	XMLdom "github.com/lestrrat-go/libxml2/dom"
	XMLtypes "github.com/lestrrat-go/libxml2/types"
	XMLxsd "github.com/lestrrat-go/libxml2/xsd"

	myXMLclib "github.com/roflcopter4/xml_addon/clib"
)

/****************************************************************************************/
// Functions

// Creates a new text reader from an XML document.
func NewTextReaderFromDoc(document XMLtypes.Document) (TextReader, error) {
	ptr, err := myXMLclib.XMLReaderWalker(document)
	if err != nil {
		return nil, err
	}
	return &XMLTextReader{
		ptr:    ptr,
		mortal: true,
	}, nil
}

/****************************************************************************************/
// Methods

// Pointer returns the pointer to the underlying C struct
func (r *XMLTextReader) Pointer() uintptr {
	return r.ptr
}

// Free releases the underlying C struct
func (r *XMLTextReader) Free() {
	myXMLclib.XMLTextReaderFree(r)
	r.ptr = 0
}

// The main iterator method. Named "TextRead" rather than just "Read" to avoid
// obvious problems.
func (r *XMLTextReader) TextRead() int {
	ret, err := myXMLclib.XMLTextReaderRead(r)
	if err != nil {
		panic(err)
	}
	return ret
}

/****************************************************************************************/
// Attributes of the current node

func (r *XMLTextReader) AttributeCount() int {
	val, err := myXMLclib.XMLTextReaderAttributeCount(r)
	if err != nil {
		return 0
	}
	return val
}

func (r *XMLTextReader) Depth() int {
	val, err := myXMLclib.XMLTextReaderDepth(r)
	if err != nil {
		return (-1)
	}
	return val
}

func (r *XMLTextReader) NodeType() myXMLclib.XMLReaderType {
	val, err := myXMLclib.XMLTextReaderNodeType(r)
	if err != nil {
		return myXMLclib.XMLReaderType(0)
	}
	return myXMLclib.XMLReaderType(val)
}

func (r *XMLTextReader) Name() string {
	str, err := myXMLclib.XMLTextReaderName(r)
	if err != nil {
		return ""
	}
	return str
}

func (r *XMLTextReader) Value() string {
	str, err := myXMLclib.XMLTextReaderValue(r)
	if err != nil {
		return ""
	}
	return str
}

/****************************************************************************************/
// "Methods"

func (r *XMLTextReader) MoveToFirstAttribute() error {
	return myXMLclib.XMLTextReaderMoveToFirstAttribute(r)
}

func (r *XMLTextReader) MoveToNextAttribute() error {
	return myXMLclib.XMLTextReaderMoveToNextAttribute(r)
}

func (r *XMLTextReader) MoveToAttributeNo(no int) error {
	return myXMLclib.XMLTextReaderMoveToAttributeNo(r, no)
}

func (r *XMLTextReader) MoveToElement() error {
	return myXMLclib.XMLTextReaderMoveToElement(r)
}

/****************************************************************************************/
// "Extensions"

func (r *XMLTextReader) SetParserProp(prop myXMLclib.XMLParserProperties, value int) error {
	return myXMLclib.XMLTextReaderSetParserProp(r, int(prop), value)
}

func (r *XMLTextReader) CurrentNode() (XMLtypes.Node, error) {
	ptr, err := myXMLclib.XMLTextReaderCurrentNode(r)
	if err != nil {
		return nil, err
	}
	return XMLdom.WrapNode(ptr)
}

func (r *XMLTextReader) SetSchema(schema *XMLxsd.Schema) error {
	return myXMLclib.XMLTextReaderSetSchema(r, schema)
}
