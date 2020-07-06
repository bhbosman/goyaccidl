package objects

import "github.com/bhbosman/yaccidl"

type IDstSrcInformation interface {
	IsIDstSrcInformation() bool
	GetSourceLexemData() *yaccidl.LexemValue
	GetSourceFolderId() string
	GetDestinationFolderId() string
	GetSourceFolderName() string
	GetSourceFileName() string
	GetDestinationFolderName() string
	DclResolveFolderUsage() (id string, folderName string)
}
