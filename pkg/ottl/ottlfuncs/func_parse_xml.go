package ottlfuncs

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ParseXMLArguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewParseXMLFactory[K any]() ottl.Factory[K] {
	return ottl.NewFactory("ParseXML", &ParseXMLArguments[K]{}, createParseXMLFunction[K])
}

func createParseXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	args, ok := oArgs.(*ParseXMLArguments[K])

	if !ok {
		return nil, fmt.Errorf("ParseXMLFactory args must be of type *ParseXMLArguments[K]")
	}

	return parseXML(args.Target), nil
}

// parseXML returns a `pcommon.Map` struct that is a result of parsing the target string as XML
func parseXML[K any](target ottl.StringGetter[K]) ottl.ExprFunc[K] {
	return func(ctx context.Context, tCtx K) (any, error) {
		targetVal, err := target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}

		parsedXML := anyXML{}

		decoder := xml.NewDecoder(strings.NewReader(targetVal))
		err = decoder.Decode(&parsedXML)
		if err != nil {
			return nil, fmt.Errorf("unmarshal xml: %w", err)
		}

		parsedMap := pcommon.NewMap()
		parsedXML.intoMap(parsedMap)

		return parsedMap, nil
	}
}

type anyXML struct {
	tag        string
	attributes []xml.Attr
	text       string
	children   []anyXML
}

// UnmarshalXML implements xml.Unmarshaler for anyXML
func (a *anyXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.tag = start.Name.Local
	a.attributes = start.Attr

	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("decode next token: %w", err)
		}

		switch t := tok.(type) {
		case xml.StartElement:
			child := anyXML{}
			err := d.DecodeElement(&child, &t)
			if err != nil {
				return fmt.Errorf("decode start element: %w", err)
			}

			a.children = append(a.children, child)
		case xml.EndElement:
			// End element means we've reached the end of parsing
			return nil
		case xml.CharData:
			// Strip leading/trailing spaces to ignore newlines and
			// indentation in formatted XML
			a.text += string(bytes.TrimSpace([]byte(t)))
		case xml.Comment: // ignore comments
		case xml.ProcInst: // ignore processing instructions
		case xml.Directive: // ignore directives
		default:
			return fmt.Errorf("unexpected token type %T", t)
		}
	}
}

// intoMap converts and adds the anyXML into the provided pcommon.Map.
func (a anyXML) intoMap(m pcommon.Map) {
	m.EnsureCapacity(4)

	m.PutStr("tag", a.tag)

	if a.text != "" {
		m.PutStr("content", a.text)
	}

	if len(a.attributes) > 0 {
		attrs := m.PutEmptyMap("attributes")
		attrs.EnsureCapacity(len(a.attributes))

		for _, attr := range a.attributes {
			attrs.PutStr(attr.Name.Local, attr.Value)
		}
	}

	if len(a.children) > 0 {
		children := m.PutEmptySlice("children")
		children.EnsureCapacity(len(a.children))

		for _, child := range a.children {
			child.intoMap(children.AppendEmpty().SetEmptyMap())
		}
	}
}
