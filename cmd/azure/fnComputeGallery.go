package azure

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetGalleryImageVersions(subscriptionId string, resourceGroup string, galleryName string, galleryImageName string, mat lib.AzureMultiAuthToken) (imageVersions lib.GalleryImageVersionList, imageVersionsDetailed []lib.GalleryImageVersionDetailed) {

	var (
		listResponse lib.ListGalleryImageVersionsResponse
		nextLink     string
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroup +
		"/providers/Microsoft.Compute/galleries/" +
		galleryName +
		"/images/" +
		galleryImageName +
		"/versions?api-version=2023-07-03"

	res, err := HttpGet(urlString, mat)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &listResponse)

	for _, val := range listResponse.Value {
		str, _ := json.Marshal(val)

		var imgVer lib.GalleryImageVersion
		json.Unmarshal(str, &imgVer)

		imageVersions.Versions = append(imageVersions.Versions, imgVer)
		imageVersionsDetailed = append(imageVersionsDetailed, val)
	}

	nextLink = listResponse.NextLink

	for nextLink != "" {
		var currentSet lib.ListGalleryImageVersionsResponse

		res, _ := HttpGet(nextLink, mat)
		json.Unmarshal(res, &currentSet)
		nextLink = currentSet.NextLink

		for _, val := range currentSet.Value {
			str, _ := json.Marshal(val)

			var imgVer lib.GalleryImageVersion

			json.Unmarshal(str, &imgVer)
			imageVersions.Versions = append(imageVersions.Versions, imgVer)
			imageVersionsDetailed = append(imageVersionsDetailed, val)
		}
	}
	return imageVersions, imageVersionsDetailed
}

func GetGalleryImage(subscriptionId string, resourceGroup string, galleryName string, galleryImageName string, mat *lib.AzureMultiAuthToken) lib.GalleryImage {

	var (
		imageDefinition lib.GalleryImage
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroup +
		"/providers/Microsoft.Compute/galleries/" +
		galleryName +
		"/images/" +
		galleryImageName +
		"?api-version=2023-07-03"

	res, err := HttpGet(urlString, *mat)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &imageDefinition)

	// if len(listResponse.Value) == 0 {
	// 	log.Fatalln("No versions exist")
	// }

	// fmt.Println(string(res))
	return imageDefinition
}

func GetAllGalleryImages(subscriptionId string, resourceGroup string, galleryName string, mat *lib.AzureMultiAuthToken) (images []lib.GalleryImage) {
	var (
		resData lib.GetAllGalleryImagesResponse
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroup +
		"/providers/Microsoft.Compute/galleries/" +
		galleryName +
		"/images" +
		"?api-version=2023-07-03"

	res, err := HttpGet(urlString, *mat)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &resData)

	for _, img := range resData.Value {
		currImage := img
		currImage.SubscriptionId = subscriptionId
		currImage.ResourceGroup = resourceGroup
		currImage.GalleryName = galleryName
		currImage.TenantName = mat.TenantName
		images = append(images, currImage)
	}

	return
}

func GetAllImagesAndVersionsForAllGalleries(tokens lib.AllTenantTokens) (imagesWithVersions []lib.GalleryImage) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		// wg2    sync.WaitGroup
		// mutex2 sync.Mutex
	)

	for _, token := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tenantGalleries := GetAllImageGalleriesForTenant(&token)
			for _, ig := range tenantGalleries {
				// GetAllGalleryImages()
				images := GetAllGalleryImages(ig.SubscriptionId, ig.ResourceGroup, ig.Name, &token)
				for _, img := range images {
					wg.Add(1)
					go func() {
						defer wg.Done()
						currImage := img
						currImage.LastAzureSync = time.Now()
						currImage.ImageVersions = make(map[string]lib.GalleryImageVersionDetailed)
						_, versionsDetailed := GetGalleryImageVersions(ig.SubscriptionId, ig.ResourceGroup, ig.Name, img.Name, token)
						// currImage.ImageVersions = append(currImage.ImageVersions, versionsDetailed...)
						for _, version := range versionsDetailed {
							currImage.ImageVersions[version.Name] = version
						}
						mutex.Lock()
						imagesWithVersions = append(imagesWithVersions, currImage)
						mutex.Unlock()
					}()
				}
			}
		}()
	}
	wg.Wait()

	return
}

func GetAllImageGalleriesForTenant(mat *lib.AzureMultiAuthToken) (imageGalleries []ImageGallery) {
	tenantSubs, err := ListSubscriptions(*mat)
	lib.CheckFatalError(err)

	for _, sub := range tenantSubs {
		subGalleries := GetAllImageGalleriesForSubscription(sub.SubscriptionID, mat)
		for _, ig := range subGalleries {
			currGal := ig
			currGal.TenantName = mat.TenantName
			imageGalleries = append(imageGalleries, currGal)
		}
	}

	return
}

func GetAllImageGalleriesForSubscription(subscriptionId string, mat *lib.AzureMultiAuthToken) (imageGalleries []ImageGallery) {
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Compute/galleries?api-version=2023-07-03"

	res, err := HttpGet(urlString, *mat)
	lib.CheckFatalError(err)

	var resData GetAllImageGalleriesForSubscriptionResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	for _, ig := range resData.Value {
		currGal := ig
		idsplit := strings.Split(ig.ID, "/")
		resGrp := strings.ToLower(idsplit[4])

		// fmt.Println(resGrp)
		// os.Exit(0)
		currGal.ResourceGroup = resGrp
		currGal.SubscriptionId = subscriptionId
		imageGalleries = append(imageGalleries, currGal)
	}

	return
}

func GetAllImagesAndVersionsForGallery(subscriptionId string, resourceGroup string, galleryName string, mat *lib.AzureMultiAuthToken) (imagesWithVersions []lib.GalleryImage) {
	images := GetAllGalleryImages(subscriptionId, resourceGroup, galleryName, mat)

	for _, img := range images {
		currImage := img
		currImage.SubscriptionId = subscriptionId
		currImage.ResourceGroup = resourceGroup
		currImage.GalleryName = galleryName
		currImage.ImageVersions = make(map[string]lib.GalleryImageVersionDetailed)
		_, versionsDetailed := GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, img.Name, *mat)
		// currImage.ImageVersions = append(currImage.ImageVersions, versionsDetailed...)
		for _, version := range versionsDetailed {
			currImage.ImageVersions[version.Name] = version
		}
		imagesWithVersions = append(imagesWithVersions, currImage)
	}

	return
}

func GetAllImageGalleriesForAllConfiguredTenants(tokens lib.AllTenantTokens) (imageGalleries []ImageGallery) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	for _, token := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tenantGalleries := GetAllImageGalleriesForTenant(&token)

			mutex.Lock()
			imageGalleries = append(imageGalleries, tenantGalleries...)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	return
}
