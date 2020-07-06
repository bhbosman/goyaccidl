package objects

type FindFoldersInUse struct {
	FolderId   string
	FolderName string
	What       map[string]string
}

func (self *FindFoldersInUse) Add(name string) {
	if _, ok := self.What[name]; !ok {
		self.What[name] = name
	}

}

func NewFindFoldersInUse(folderId, folderName string, what string) *FindFoldersInUse {
	result := &FindFoldersInUse{
		FolderId:   folderId,
		FolderName: folderName,
		What:       make(map[string]string),
	}
	if what == "" {
		result.What[folderName] = folderName
	} else {
		result.What[what] = what
	}

	return result
}

type FileUsage map[string]*FindFoldersInUse

type FileUsages map[string]*FileData

func (self FileUsages) ToDclArray() IDclArray {
	var result IDclArray
	for _, v := range self {
		result = append(result, v.Members...)
	}
	return result
}

type FileData struct {
	SourceFolderId      string    `json:"source_folder_id"`
	DestinationFolderId string    `json:"destination_folder_id"`
	FolderUsage         FileUsage `json:"folder_usage"`
	Members             IDclArray `json:"members"`
	TargetFileName      string    `json:"target_file_name"`
}

func (self *FileData) Add(dcl IDcl) {
	self.Members = append(self.Members, dcl)
}

func NewFileData(SourceFolderId string, DestinationFolderId string, targetFileName string) *FileData {
	return &FileData{
		SourceFolderId:      SourceFolderId,
		DestinationFolderId: DestinationFolderId,
		FolderUsage:         make(FileUsage),
		Members:             nil,
		TargetFileName:      targetFileName,
	}
}
