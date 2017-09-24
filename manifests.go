package goBungieNet

import (
  "fmt"
  "errors"

  "github.com/mitchellh/mapstructure"
)

type Manifest struct {
  Version string
  AssetContentPath string
  //MobileGearAssetDatabases
  //MobileWorldContentPaths
  ClanBannerDatabasePath string
}

var (
  currentManifest Manifest
)

func init() {
  currentManifest.Version = "No Manifest"
}

func ManifestUpdate() error {
  response, err := get("/Destiny2/Manifest/")
  if (err != nil) { return err }

  if response.ErrorCode.isError() {
    return errors.New(fmt.Sprintf("Error: %s", response.ErrorCode))
  }

  err = mapstructure.Decode(response.Response, &currentManifest)

  if (err != nil) { return err }

  fmt.Printf("Manifest Version is now: %s", currentManifest.Version)
  return nil
}
