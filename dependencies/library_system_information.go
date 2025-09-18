package dependencies

import "strings"

type LibrarySystemInformation struct {
	Name     string
	Version  string
	Platform string
}

func CreateLibrarySystemInformation(colonSeparatedAttributeList string) LibrarySystemInformation {
	attributes := strings.Split(colonSeparatedAttributeList, ":")
	return LibrarySystemInformation{
		Name:     attributes[0],
		Version:  attributes[1],
		Platform: attributes[2],
	}
}

func ParseLibraryReplacements(unparsedReplacements string) []LibrarySystemInformation {
	if unparsedReplacements == "" {
		return []LibrarySystemInformation{}
	}

	var libraryInformationList []LibrarySystemInformation

	libraryInformationStringList := strings.Split(unparsedReplacements, ",")
	for _, libraryInformation := range libraryInformationStringList {
		libraryInformationList = append(libraryInformationList, CreateLibrarySystemInformation(libraryInformation))
	}
	return libraryInformationList
}
