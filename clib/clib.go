package clib

/*
 * Implements the actual cgo wrapper code.
 */

/*

#cgo pkg-config: libxml-2.0
#include <string.h>
#include <stdbool.h>
#include <stdio.h>
#include <libxml/HTMLparser.h>
#include <libxml/HTMLtree.h>
#include <libxml/c14n.h>
#include <libxml/globals.h>
#include <libxml/parser.h>
#include <libxml/parserInternals.h>
#include <libxml/tree.h>
#include <libxml/xmlerror.h>
#include <libxml/xmlreader.h>
#include <libxml/xmlschemas.h>
#include <libxml/xpath.h>
#include <libxml/xpathInternals.h>

// Macro wrapper function. cgo cannot detect function-like macros,
// so this is how we avoid it
static inline xmlError* MY_xmlLastError() {
	return xmlGetLastError();
}

static int my_get_node_type(xmlNodePtr node)
{
	if (!node)
		return (-1);
	return node->type;
}

*/
import "C"

import (
	// XML "github.com/lestrrat-go/libxml2"
	XMLclib "github.com/lestrrat-go/libxml2/clib"
	XMLtypes "github.com/lestrrat-go/libxml2/types"

	// XMLxsd "github.com/lestrrat-go/libxml2/xsd"
	"unsafe"

	"github.com/pkg/errors"
)

var (
	ErrInvalidTextReader = errors.New("invalid text reader")
)

type ptrSource = XMLclib.PtrSource

/****************************************************************************************/
// Copied verbatim from libxml2/clib

func validDocumentPtr(doc ptrSource) (*C.xmlDoc, error) {
	if doc == nil {
		return nil, XMLclib.ErrInvalidDocument
	}

	if dptr := doc.Pointer(); dptr != 0 {
		return (*C.xmlDoc)(unsafe.Pointer(dptr)), nil
	}
	return nil, XMLclib.ErrInvalidDocument
}

func validNodePtr(n ptrSource) (*C.xmlNode, error) {
	if n == nil {
		return nil, XMLclib.ErrInvalidNode
	}

	nptr := n.Pointer()
	if nptr == 0 {
		return nil, XMLclib.ErrInvalidNode
	}

	return (*C.xmlNode)(unsafe.Pointer(nptr)), nil
}

func validAttributePtr(n ptrSource) (*C.xmlAttr, error) {
	if n == nil {
		return nil, XMLclib.ErrInvalidAttribute
	}

	if nptr := n.Pointer(); nptr != 0 {
		return (*C.xmlAttr)(unsafe.Pointer(nptr)), nil
	}

	return nil, XMLclib.ErrInvalidAttribute
}

func validSchemaPtr(schema ptrSource) (*C.xmlSchema, error) {
	if schema == nil {
		return nil, XMLclib.ErrInvalidSchema
	}
	sptr := schema.Pointer()
	if sptr == 0 {
		return nil, XMLclib.ErrInvalidSchema
	}

	return (*C.xmlSchema)(unsafe.Pointer(sptr)), nil
}

func validTextReaderPtr(reader XMLclib.PtrSource) (*C.xmlTextReader, error) {
	if reader == nil {
		return nil, ErrInvalidTextReader
	}
	rdptr := reader.Pointer()
	if rdptr == 0 {
		return nil, ErrInvalidTextReader
	}

	return (*C.xmlTextReader)(unsafe.Pointer(rdptr)), nil
}

func xmlCharToString(s *C.xmlChar) string {
	return C.GoString((*C.char)(unsafe.Pointer(s)))
}

/****************************************************************************************/

// Creates a new reader
func XMLReaderWalker(document XMLtypes.Document) (uintptr, error) {
	docptr, err := validDocumentPtr(document)
	if err != nil {
		return 0, err
	}

	reader := C.xmlReaderWalker(docptr)
	if reader == nil {
		e := C.MY_xmlLastError()
		if e != nil {
			return 0, errors.New("failed to create textReader: " + C.GoString(e.message))
		} else {
			return 0, errors.New("failed to create textReader")
		}
	}

	return uintptr(unsafe.Pointer(reader)), nil
}

func XMLTextReaderFree(reader ptrSource) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}
	C.xmlFreeTextReader(rdptr)
	return nil
}

func XMLTextReaderRead(reader ptrSource) (int, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	return int(C.xmlTextReaderRead(rdptr)), nil
}

/****************************************************************************************/
// Attributes of the current node

func XMLTextReaderAttributeCount(reader ptrSource) (int, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	ret := C.xmlTextReaderAttributeCount(rdptr)
	return int(ret), nil
}

func XMLTextReaderDepth(reader ptrSource) (int, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	ret := C.xmlTextReaderDepth(rdptr)
	return int(ret), nil
}

func XMLTextReaderHasAttributes(reader ptrSource) (bool, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return false, err
	}
	ret := C.xmlTextReaderHasAttributes(rdptr)
	return (ret != 0), nil
}

func XMLTextReaderHasValue(reader ptrSource) (bool, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return false, err
	}
	ret := C.xmlTextReaderHasValue(rdptr)
	return (ret != 0), nil
}

func XMLTextReaderIsDefault(reader ptrSource) (bool, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return false, err
	}
	ret := C.xmlTextReaderIsDefault(rdptr)
	return (ret != 0), nil
}

func XMLTextReaderIsEmptyElement(reader ptrSource) (bool, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return false, err
	}
	ret := C.xmlTextReaderIsEmptyElement(rdptr)
	return (ret != 0), nil
}

func XMLTextReaderNodeType(reader ptrSource) (int, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	ret := C.xmlTextReaderNodeType(rdptr)
	return int(ret), nil
}

func XMLTextReaderQuoteChar(reader ptrSource) (int, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	ret := C.xmlTextReaderQuoteChar(rdptr)
	return int(ret), nil
}

func XMLTextReaderReadState(reader ptrSource) (int, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	ret := C.xmlTextReaderReadState(rdptr)
	return int(ret), nil
}

func XMLTextReaderIsNamespaceDecl(reader ptrSource) (bool, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return false, err
	}
	ret := C.xmlTextReaderQuoteChar(rdptr)
	return (ret != 0), nil
}

func XMLTextReaderBaseUri(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstBaseUri(rdptr)
	return xmlCharToString(ret), nil
}

func XMLTextReaderLocalName(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstLocalName(rdptr)
	return xmlCharToString(ret), nil
}

func XMLTextReaderName(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstName(rdptr)
	return xmlCharToString(ret), nil
}

func XMLTextReaderNamespaceUri(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstNamespaceUri(rdptr)
	return xmlCharToString(ret), nil
}

func XMLTextReaderPrefix(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstPrefix(rdptr)
	return xmlCharToString(ret), nil
}

func XMLTextReaderXmlLang(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstXmlLang(rdptr)
	return xmlCharToString(ret), nil
}

func XMLTextReaderValue(reader ptrSource) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderConstValue(rdptr)
	return xmlCharToString(ret), nil
}

/****************************************************************************************/
// "Methods"

func XMLTextReaderGetAttributeNo(reader ptrSource, no int) (string, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return "", err
	}
	ret := C.xmlTextReaderGetAttributeNo(rdptr, C.int(no))
	return xmlCharToString(ret), nil
}

func XMLTextReaderMoveToAttributeNo(reader ptrSource, no int) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}
	ret := C.xmlTextReaderMoveToAttributeNo(rdptr, C.int(no))
	if ret != 0 {
		e := C.MY_xmlLastError()
		if e != nil {
			return errors.New("cannot move to element: " + C.GoString(e.message))
		} else {
			return errors.New("cannot move to element")
		}
	}
	return nil
}

func XMLTextReaderMoveToFirstAttribute(reader ptrSource) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}
	ret := C.xmlTextReaderMoveToFirstAttribute(rdptr)
	if ret != 0 {
		e := C.MY_xmlLastError()
		if e != nil {
			return errors.New("cannot move to attribute: " + C.GoString(e.message))
		} else {
			return errors.New("cannot move to attribute")
		}
	}
	return nil
}

func XMLTextReaderMoveToNextAttribute(reader ptrSource) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}
	ret := C.xmlTextReaderMoveToNextAttribute(rdptr)
	if ret != 0 {
		e := C.MY_xmlLastError()
		if e != nil {
			return errors.New("cannot move to attribute: " + C.GoString(e.message))
		} else {
			return errors.New("cannot move to attribute")
		}
	}
	return nil
}

func XMLTextReaderMoveToElement(reader ptrSource) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}
	ret := C.xmlTextReaderMoveToElement(rdptr)
	if ret != 0 {
		e := C.MY_xmlLastError()
		if e != nil {
			return errors.New("cannot move to element: " + C.GoString(e.message))
		} else {
			return errors.New("cannot move to element")
		}
	}
	return nil
}

/****************************************************************************************/
// "Extensions"

func XMLTextReaderCurrentNode(reader ptrSource) (uintptr, error) {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return 0, err
	}
	node := C.xmlTextReaderCurrentNode(rdptr)
	return uintptr(unsafe.Pointer(node)), nil
}

func XMLTextReaderSetSchema(reader, schema ptrSource) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}
	sptr, err := validSchemaPtr(schema)
	if err != nil {
		return err
	}

	ret := C.xmlTextReaderSetSchema(rdptr, sptr)
	if ret != 0 {
		e := C.MY_xmlLastError()
		if e != nil {
			return errors.New("Error setting schema: " + C.GoString(e.message))
		} else {
			return errors.New("Error setting schema")
		}
	}
	return nil
}

func XMLTextReaderSetParserProp(reader ptrSource, prop, value int) error {
	rdptr, err := validTextReaderPtr(reader)
	if err != nil {
		return err
	}

	ret := C.xmlTextReaderSetParserProp(rdptr, C.int(prop), C.int(value))
	if ret != 0 {
		e := C.MY_xmlLastError()
		if e != nil {
			return errors.New("Error setting parser property: " + C.GoString(e.message))
		} else {
			return errors.Errorf("Error setting parser property (%d)", ret)
		}
	}
	return nil
}
