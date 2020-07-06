package yacc

import (
	"context"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"log"
	"testing"
)

func dddd(t *testing.T, data []lexem) (*fxtest.App, IYaccIdlLexer) {
	writer := &log2.TestWriter{
		TestLogger: t,
	}
	var yaccIdlLexer IYaccIdlLexer = nil
	logger := log.New(writer, "", log.LstdFlags)
	app := fxtest.New(
		t,
		log2.ProvideLogFactory(logger, nil),
		Service.Provide(),

		fx.Provide(
			func() IYaccIdlLexerHandler {
				return newIdlLexerHandlerForTesting(0, data)
			}),
		fx.Provide(
			func(idlDefaultTypes *Service.IdlDefaultTypes, logFactory *log2.LogFactory, yaccIdlLexerHandler IYaccIdlLexerHandler) (IYaccIdlLexer, error) {
				return NewIdlLexer(idlDefaultTypes, logFactory.Create("IYaccIdlLexer"), yaccIdlLexerHandler)
			}),
		fx.Populate(&yaccIdlLexer))

	return app, yaccIdlLexer

}

func TestLexerPragma(t *testing.T) {

	t.Run("", func(t *testing.T) {
		app, yaccIdlLexer := dddd(t, []lexem{
			{token: yaccidl.RWstruct, data: nil},
			{token: yaccidl.Identifier, data: "Address"},
			{token: '{', data: nil},
			{token: '}', data: nil},
			{token: ';', data: nil},
		})
		app.Start(context.TODO())
		defer app.Stop(context.TODO())
		parser := yaccidl.YaccIdlNewParser()
		parseValue := parser.Parse(yaccIdlLexer)
		assert.Equal(t, 0, parseValue)
		spec := yaccIdlLexer.GetSpecification()
		assert.NotNil(t, spec)
		structValue, ok := spec.(yaccidl.IStructDcl)
		if !assert.True(t, ok) {
			return
		}
		assert.Equal(t, "Address", structValue.Identifier())
		spec = spec.GetNextNode()
		if !assert.Nil(t, spec) {
			return
		}
	})
}
