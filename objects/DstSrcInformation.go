package objects

import "github.com/bhbosman/yaccidl"

type DstSrcInformation struct {
	SourceLexemData *yaccidl.LexemValue `json:"source_lexem_data"`
	Destination     Destination         `json:"destination"`
}

func (self *DstSrcInformation) DclResolveFolderUsage() (id string, folderName string) {
	id = self.SourceLexemData.SourceFolderId
	destinationFolderId := self.Destination.FolderId
	folderName = self.SourceLexemData.SourceFolderName
	destinationFolderName := self.Destination.Folder
	if destinationFolderId != "" {
		id = destinationFolderId
		folderName = destinationFolderName
	}
	return id, folderName
}

func (self *DstSrcInformation) GetDestinationFolderName() string {
	return self.Destination.Folder
}

func (self *DstSrcInformation) GetDestinationFolderId() string {
	return self.Destination.FolderId
}

func (self *DstSrcInformation) GetSourceFileName() string {
	return self.SourceLexemData.GetSourceFileName()
}

func (self *DstSrcInformation) GetSourceFolderName() string {
	return self.SourceLexemData.GetSourceFolderName()
}

func (self *DstSrcInformation) GetSourceFolderId() string {
	return self.SourceLexemData.GetSourceFolderId()
}

func (self *DstSrcInformation) GetSourceLexemData() *yaccidl.LexemValue {
	return self.SourceLexemData
}

func (self *DstSrcInformation) toIDstSrcInformation() IDstSrcInformation {
	return self
}

func (self *DstSrcInformation) IsIDstSrcInformation() bool {
	return true
}
