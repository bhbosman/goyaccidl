package processIdlFile

import (
	"context"
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/goyaccidl/process"
	"github.com/bhbosman/goyaccidl/typeHelpers"
	"github.com/bhbosman/goyaccidl/writers/gowriter"
	"github.com/bhbosman/goyaccidl/yacc"
	"github.com/bhbosman/goyaccidl/yaccSpecProcessing"
	"github.com/bhbosman/lexidl"
	"github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"log"
	"os"
	"testing"
)

type fff struct {
	fx.In
	YaccIdlLexer                yacc.IYaccIdlLexer
	YaccTokensToYaccSpecProcess process.IProcess `name:"YaccTokensToYaccSpec"`
	YaccSpecToDeclTypesProcess  process.IProcess `name:"YaccSpecToDeclTypes"`
	FlattenDeclaredTypesProcess process.IProcess `name:"FlattenDeclaredTypes"`
}

func TestName(t *testing.T) {
	writer := &log2.TestWriter{
		TestLogger: t,
	}
	logger := log.New(writer, "", log.LstdFlags)
	resolveFileName := &ctx.ResolveFileName{}
	defaultProcessor := "IDLProcessor"
	OutputFolder, _ := os.Getwd()
	flags := &ctx.GoYaccAppCtx{

		OutputFolder: OutputFolder,
		ForceOutput:  false,
		Verbose:      false,

		InputFileName: "",
		InputReader:   nil,

		IdlDefinitions:   ctx.IdlDefinitions{},
		Stdin:            nil,
		Stdout:           nil,
		Stderr:           nil,
		GoSrcPath:        "",
		DefaultProcessor: defaultProcessor,
	}
	t.Run("type prefix declared ahead", func(t *testing.T) {
		ss := `
			typeprefix CORBA "omg.org";
			module CORBA 
			{
				struct AA 
				{
				};
			};
		`
		f := fff{}
		app := fxtest.New(
			t,
			log2.ProvideLogFactory(logger, nil),
			ctx.ProvideGoYaccAppCtx(flags),
			ctx.ProvideResolveFileName(resolveFileName),

			implementations.ProvideUniqueSessionNumber(),
			typeHelpers.ProvideTypeHelpers(),
			ProvideIdlProcess(),
			Service.Provide(),
			yaccidl.ProvideFileInformation(),
			ProvideCreateKnownTypeService(),
			yaccSpecProcessing.Provide(),
			gowriter.ProvideAllWriters(),

			gowriter.ProvideGoWriterFactory(),

			fx.Provide(func(idlDefaultTypes *Service.IdlDefaultTypes, logFactory *log2.LogFactory) (yacc.IYaccIdlLexer, error) {
				lexHandler, _ := lexidl.NewLexIdlHandlerFromData(t.Name(), ss, nil)
				return yacc.NewIdlLexer(idlDefaultTypes, logFactory.Create("(ABC)"), lexHandler)
			}),
			fx.Populate(&f))

		_ = app.Start(context.TODO())
		defer app.Stop(context.TODO())

		yaccSpec, err := f.YaccTokensToYaccSpecProcess.Run(f.YaccIdlLexer)
		if !assert.NoError(t, err) {
			return
		}
		declTypes, err := f.YaccSpecToDeclTypesProcess.Run(yaccSpec)
		if !assert.NoError(t, err) {
			return
		}

		stm, err := f.FlattenDeclaredTypesProcess.Run(declTypes)
		if !assert.NoError(t, err) {
			return
		}
		dcl, ok := stm.(*objects.FlattenDeclaredTypesOut)
		if !assert.True(t, ok) {
			return
		}
		assert.Len(t, dcl.InputStreamDcl.Members, 1)
		structDcl, ok := dcl.InputStreamDcl.Members[0].(objects.IStructDcl)

		if !assert.True(t, ok) {
			return
		}
		if !assert.Equal(t, "omg.org", structDcl.GetTypePrefix()) {
			return
		}
		assert.Equal(t, "CORBA::AA", string(dcl.InputStreamDcl.Members[0].GetName()))
	})

	t.Run("type prefix declared after", func(t *testing.T) {
		ss := `
			module CORBA 
			{
				typeprefix CORBA "omg.org";
				struct AA 
				{
				};
			};
		`
		f := fff{}
		app := fxtest.New(
			t,
			log2.ProvideLogFactory(logger, nil),
			ctx.ProvideGoYaccAppCtx(flags),
			ctx.ProvideResolveFileName(resolveFileName),

			implementations.ProvideUniqueSessionNumber(),
			typeHelpers.ProvideTypeHelpers(),
			ProvideIdlProcess(),
			Service.Provide(),
			yaccidl.ProvideFileInformation(),
			ProvideCreateKnownTypeService(),

			yaccSpecProcessing.Provide(),
			gowriter.ProvideAllWriters(),

			gowriter.ProvideGoWriterFactory(),

			fx.Provide(func(idlDefaultTypes *Service.IdlDefaultTypes, logFactory *log2.LogFactory) (yacc.IYaccIdlLexer, error) {
				lexHandler, _ := lexidl.NewLexIdlHandlerFromData(t.Name(), ss, nil)
				return yacc.NewIdlLexer(idlDefaultTypes, logFactory.Create("(ABC)"), lexHandler)
			}),
			fx.Populate(&f))

		_ = app.Start(context.TODO())
		defer app.Stop(context.TODO())

		yaccSpec, err := f.YaccTokensToYaccSpecProcess.Run(f.YaccIdlLexer)
		if !assert.NoError(t, err) {
			return
		}
		declTypes, err := f.YaccSpecToDeclTypesProcess.Run(yaccSpec)
		if !assert.NoError(t, err) {
			return
		}

		stm, err := f.FlattenDeclaredTypesProcess.Run(declTypes)
		if !assert.NoError(t, err) {
			return
		}
		dcl, ok := stm.(*objects.FlattenDeclaredTypesOut)
		if !assert.True(t, ok) {
			return
		}
		assert.Len(t, dcl.InputStreamDcl.Members, 1)
		structDcl, ok := dcl.InputStreamDcl.Members[0].(*objects.StructDcl)

		if !assert.True(t, ok) {
			return
		}
		if !assert.Equal(t, "omg.org", structDcl.GetTypePrefix()) {
			return
		}
		assert.Equal(t, "CORBA::AA", string(dcl.InputStreamDcl.Members[0].GetName()))
	})

	t.Run("type prefix for something else", func(t *testing.T) {
		ss := `
			module CORBA 
			{
				typeprefix SomethingElse "omg.org";
				struct AA 
				{
				};
			};
		`
		f := fff{}
		app := fxtest.New(
			t,
			log2.ProvideLogFactory(logger, nil),
			ctx.ProvideGoYaccAppCtx(flags),
			ctx.ProvideResolveFileName(resolveFileName),

			implementations.ProvideUniqueSessionNumber(),
			typeHelpers.ProvideTypeHelpers(),
			ProvideIdlProcess(),
			Service.Provide(),
			yaccidl.ProvideFileInformation(),
			ProvideCreateKnownTypeService(),
			yaccSpecProcessing.Provide(),

			gowriter.ProvideAllWriters(),

			gowriter.ProvideGoWriterFactory(),

			fx.Provide(func(idlDefaultTypes *Service.IdlDefaultTypes, logFactory *log2.LogFactory) (yacc.IYaccIdlLexer, error) {
				lexHandler, _ := lexidl.NewLexIdlHandlerFromData(t.Name(), ss, nil)
				return yacc.NewIdlLexer(idlDefaultTypes, logFactory.Create("(ABC)"), lexHandler)
			}),
			fx.Populate(&f))

		_ = app.Start(context.TODO())
		defer app.Stop(context.TODO())

		yaccSpec, err := f.YaccTokensToYaccSpecProcess.Run(f.YaccIdlLexer)
		if !assert.NoError(t, err) {
			return
		}
		declTypes, err := f.YaccSpecToDeclTypesProcess.Run(yaccSpec)
		if !assert.NoError(t, err) {
			return
		}

		stm, err := f.FlattenDeclaredTypesProcess.Run(declTypes)
		if !assert.NoError(t, err) {
			return
		}
		dcl, ok := stm.(*objects.FlattenDeclaredTypesOut)
		if !assert.True(t, ok) {
			return
		}
		assert.Len(t, dcl.InputStreamDcl.Members, 1)
		structDcl, ok := dcl.InputStreamDcl.Members[0].(*objects.StructDcl)

		if !assert.True(t, ok) {
			return
		}
		if !assert.Equal(t, "", structDcl.GetTypePrefix()) {
			return
		}
		assert.Equal(t, "CORBA::AA", string(dcl.InputStreamDcl.Members[0].GetName()))
	})

	t.Run("type prefix for struct", func(t *testing.T) {
		ss := `
			module CORBA 
			{
				typeprefix AA "aa.omg.org";
				struct AA 
				{
				};
			};
		`
		f := fff{}
		app := fxtest.New(
			t,
			log2.ProvideLogFactory(logger, nil),
			ctx.ProvideGoYaccAppCtx(flags),
			ctx.ProvideResolveFileName(resolveFileName),

			implementations.ProvideUniqueSessionNumber(),
			typeHelpers.ProvideTypeHelpers(),
			gowriter.ProvideAllWriters(),

			gowriter.ProvideGoWriterFactory(),
			ProvideIdlProcess(),
			Service.Provide(),
			yaccidl.ProvideFileInformation(),
			ProvideCreateKnownTypeService(),
			yaccSpecProcessing.Provide(),

			fx.Provide(func(idlDefaultTypes *Service.IdlDefaultTypes, logFactory *log2.LogFactory) (yacc.IYaccIdlLexer, error) {
				lexHandler, _ := lexidl.NewLexIdlHandlerFromData(t.Name(), ss, nil)
				return yacc.NewIdlLexer(idlDefaultTypes, logFactory.Create("(ABC)"), lexHandler)
			}),
			fx.Populate(&f))

		_ = app.Start(context.TODO())
		defer app.Stop(context.TODO())

		yaccSpec, err := f.YaccTokensToYaccSpecProcess.Run(f.YaccIdlLexer)
		if !assert.NoError(t, err) {
			return
		}
		declTypes, err := f.YaccSpecToDeclTypesProcess.Run(yaccSpec)
		if !assert.NoError(t, err) {
			return
		}

		stm, err := f.FlattenDeclaredTypesProcess.Run(declTypes)
		if !assert.NoError(t, err) {
			return
		}
		dcl, ok := stm.(*objects.FlattenDeclaredTypesOut)
		if !assert.True(t, ok) {
			return
		}
		assert.Len(t, dcl.InputStreamDcl.Members, 1)
		structDcl, ok := dcl.InputStreamDcl.Members[0].(*objects.StructDcl)

		if !assert.True(t, ok) {
			return
		}
		if !assert.Equal(t, "aa.omg.org", structDcl.GetTypePrefix()) {
			return
		}
		assert.Equal(t, "CORBA::AA", string(dcl.InputStreamDcl.Members[0].GetName()))
	})
}
