package citrix

import (
	"strconv"
	"strings"
)

func (machineCatalogs *MachineCatalogs) ListImageVersions() map[string]MachineCatalogCurrentImage {
	mcats := make(map[string]MachineCatalogCurrentImage)

	for _, mc := range *machineCatalogs {
		var currImg MachineCatalogCurrentImage
		// mcats[mc.Name] = mc.ProvisioningScheme.CurrentDiskImage.Image.Name
		if mc.ProvisioningScheme.CurrentDiskImage == nil {
			preparedImageVersion := strconv.Itoa(mc.ProvisioningScheme.CurrentImageVersion.ImageVersion.Number)
			currImg.IsPreparedImage = true
			currImg.PreparedImageName = mc.ProvisioningScheme.CurrentImageVersion.ImageVersion.ImageDefinition.Name
			currImg.PreparedImageVersion = preparedImageVersion
		}

		XdPath := strings.Split(mc.ProvisioningScheme.MasterImage.XdPath, "\\")

		for _, str := range XdPath {
			if strings.Contains(str, ".resourcegroup") {
				currImg.ResourceGroup = strings.Split(str, ".resourcegroup")[0]
			} else if strings.Contains(str, ".gallery") {
				currImg.ImageGallery = strings.Split(str, ".gallery")[0]
			} else if strings.Contains(str, ".imagedefinition") {
				currImg.ImageDefinitionName = strings.Split(str, ".imagedefinition")[0]
			} else if strings.Contains(str, ".imageversion") {
				currImg.Version = strings.Split(str, ".imageversion")[0]
			}
		}

		currImg.LastCitrixSync = mc.LastCitrixSync

		mcats[mc.Name] = currImg
	}

	return mcats
}
