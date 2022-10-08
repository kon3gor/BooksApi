package mapper

import (
	"fmt"
)

func MapSmallImage(coverId int) string {
	return mapImageWithSuffix(coverId, "S")
}

func MapMediumImage(coverId int) string {
	return mapImageWithSuffix(coverId, "M")
}

func MapLargeImage(coverId int) string {
	return mapImageWithSuffix(coverId, "L")
}

const OPEN_LIB_COVER_URL = "https://covers.openlibrary.org/b/id/" 

func mapImageWithSuffix(coverId int, suffix string) string {
	return fmt.Sprintf("%s%d-%s.jpg", OPEN_LIB_COVER_URL, coverId, suffix)
}
