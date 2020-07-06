SplitProject: install clean ClearExports \
	CorbaProject \
  	IopProject \
  	GiopProject\
	MiopProject\
	IiopProject\
	TimeBaseProject \
	CosNotificationProject \
	CosEventCommProject \
	MathService

CorbaProject: install
	rm -f ../goIdlCorba/*.idl.go
	rm -f ../goIdlCorba/xdl_*.go
	rm -f ../goIdlCorba/*.idl.*.go

	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlCorba   ../orb/idl/omg/orb.idl
	go build ../goIdlCorba

IopProject: install CorbaProject
	rm -f ../goIdlIop/*.idl.go
	rm -f ../goIdlIop/xdl_*.go
	rm -f ../goIdlIop/*.idl.*.go

	goyaccidl -v     -ff  -of ../goIdlIop   ../orb/idl/omg/IOP.idl
	go build ../goIdlIop

GiopProject: install IopProject
	rm -f ../goIdlGiop/*.idl.go
	rm -f ../goIdlGiop/xdl_*.go
	rm -f ../goIdlGiop/*.idl.*.go
	goyaccidl -v -idldef "_COMPONENT_REPOSITORY_,GIOP_1_1"    -ff  -of ../goIdlGiop   ../orb/idl/omg/GIOP.idl
	go build ../goIdlGiop

MiopProject: install IopProject GiopProject
	rm -f ../goIdlMiop/*.idl.go
	rm -f ../goIdlMiop/xdl_*.go
	rm -f ../goIdlMiop/*.idl.*.go
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlMiop   ../orb/idl/omg/MIOP.idl
	go build ../goIdlMiop

MathService:install
	rm -f ../goIdlMathService/*.idl.go
	rm -f ../goIdlMathService/xdl_*.go
	rm -f ../goIdlMathService/*.idl.*.go
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlMathService   ../orb/idl/MathService/MathService.idl
	go build ../goIdlMiop


IiopProject: IopProject
	rm -f ../goIdlIiop/*.idl.go
	rm -f ../goIdlIiop/xdl_*.go
	rm -f ../goIdlIiop/*.idl.*.go
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlIiop   ../orb/idl/omg/IIOP.idl
	go build ../goIdlIiop



TimeBaseProject: install
	rm -f ../goIdlTimeBase/*.idl.go
	rm -f ../goIdlTimeBase/xdl_*.go
	rm -f ../goIdlTimeBase/*.idl.*.go
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlTimeBase   ../orb/idl/omg/TimeBase.idl
	go build ../goIdlTimeBase


CosNotificationProject:
	rm -f ../goIdlCosNotification/*.idl.go
	rm -f ../goIdlCosNotification/xdl_*.go
	rm -f ../goIdlCosNotification/*.idl.*.go
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlCosNotification   ../orb/idl/omg/CosNotification.idl
	go build ../goIdlCosNotification

CosEventCommProject:
	rm -f ../goIdlCosEventComm/*.idl.go
	rm -f ../goIdlCosEventComm/xdl_*.go
	rm -f ../goIdlCosEventComm/*.idl.*.go
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../goIdlCosEventComm   ../orb/idl/omg/CosEventComm.idl
	go build ../goIdlCosEventComm



all: clean build test install run_files buidOmg  buildTests buildGoCorba


clean:
	rm -rf goyaccidl
	rm -f ./output/*.orb.json
	rm -f ../orb/src/omg/*.idl.go
	rm -f ../orb/src/omg/*.idl.*.go
build:
	go build
buidOmg:
	go build ./../orb/src/omg

test:
	go test -v
install: build test
	go install
run_files: \
	ClearExports \
	PrimitiveProcessor \
	ORB \
	IOP  \
	IIOP \
	GIOP \
	MIOP \
	TimeBase \
	CosNotification \
	CosEventComm \
	CosNotifyComm \
	CosNotifyFilter \
	CosEventChannelAdmin \
	CosNotifyChannelAdmin \
	CosBridgeAdmin \
	CosCollection \
	CosTransactions \
	CosConcurrencyControl \
	CosNaming \
	CosTime \
	CosTrading \
	CosTypedEventComm \
	CosTypedEventChannelAdmin \
	CosTypedNotifyComm \
	CosTypedNotifyChannelAdmin \











ClearExports:
	goyaccidl -v -processor ClearExports
PrimitiveProcessor:
	goyaccidl -v -processor PrimitiveProcessor -ff  -of ./idl/omg/orb.idl
ORB:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/orb.idl
IOP: ORB
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/IOP.idl
IIOP: IOP
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/IIOP.idl
GIOP: IIOP
	goyaccidl -v -idldef "_COMPONENT_REPOSITORY_,GIOP_1_1"    -ff  -of ../orb/src/omg/   ../orb/idl/omg/GIOP.idl
MIOP: install GIOP
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/MIOP.idl
TimeBase:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/TimeBase.idl
CosNotification:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosNotification.idl
CosEventComm:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosEventComm.idl
CosNotifyComm: CosNotification CosEventComm
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosNotifyComm.idl
CosNotifyFilter: ORB CosNotifyComm
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosNotifyFilter.idl
CosEventChannelAdmin: CosEventComm
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosEventChannelAdmin.idl
CosNotifyChannelAdmin: CosNotification CosNotifyFilter CosNotifyComm CosEventChannelAdmin
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosNotifyChannelAdmin.idl
CosBridgeAdmin: CosNotifyChannelAdmin
#	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosBridgeAdmin.idl
CosCollection:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosCollection.idl
CosTransactions:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTransactions.idl
CosConcurrencyControl: CosTransactions
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosConcurrencyControl.idl
CosNaming:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosNaming.idl
CosTime: TimeBase
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTime.idl
CosTrading:
#	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTrading.idl
CosTypedEventComm: CosEventComm;
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTypedEventComm.idl
CosTypedEventChannelAdmin: CosEventChannelAdmin CosTypedEventComm
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTypedEventChannelAdmin.idl
CosTypedNotifyComm: CosNotifyChannelAdmin CosTypedEventComm
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTypedNotifyComm.idl
CosTypedNotifyChannelAdmin: CosNotifyChannelAdmin CosTypedNotifyComm CosTypedEventChannelAdmin
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/   ../orb/idl/omg/CosTypedNotifyChannelAdmin.idl










########################################################################################################################
########################################################################################################################
########################################################################################################################
ModuleOne:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/Abc   ../orb/idl/omg/testAbc.idl
ModuleTwo:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/Def   ../orb/idl/omg/testDef.idl
ModuleThree:
	goyaccidl -v -idldef _COMPONENT_REPOSITORY_    -ff  -of ../orb/src/omg/ModuleThree   ../orb/idl/omg/testThree.idl
buildModuleOne:
	go build ../orb/src/omg/Abc
buildModuleTwo:
	go build ../orb/src/omg/Def
buildModuleThree:
	go build ../orb/src/omg/ModuleThree

buildTests: install ModuleOne ModuleTwo ModuleThree buildModuleOne buildModuleTwo buildModuleThree


########################################################################################################################
########################################################################################################################
########################################################################################################################


buildGoCorba: install buildTest buildTest2 buildGoCdr buildGoCdrTest

buildTest:
	goyaccidl -v  -ff  -of ../orb/src/TestData/golang   ../orb/idl/TestData/test.idl


buildTest2:
	goyaccidl -v   -ff  -of ../orb/src/TestData/golang   ../orb/idl/TestData/test2.idl

buildGoCdr:
	go build ./../gocorba/cdr

buildGoCdrTest:
	go test ./../gocorba/cdr
	#go build ./../gocorba/giop
