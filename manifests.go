package goBungieNet

import (
  "fmt"
  "errors"
  "io"
  "io/ioutil"
  "os"
  "net/http"
  "archive/zip"
  "bytes"

  "github.com/mitchellh/mapstructure"
)

type Manifest struct {
  Version string
  AssetContentPath string `mapstructure:"mobileAssetContentPath"`
  ClanBannerDatabasePath string `mapstructure:"mobileClanBannerDatabasePath"`
  GearAssetDatabases []GearAssetDatabase `mapstructure:"mobileGearAssetDataBases"`
  GearCDN map[string]string `mapstructure:"mobileGearCDN"`
  WorldContentPaths map[string]string `mapstructure:"mobileWorldContentPaths"`
}

type GearAssetDatabase struct {
  Version int32
  Path string
}

var (
  currentManifest Manifest
  ManifestPath string
)

func init() {
  currentManifest.Version = "No Manifest"
  ManifestPath = ""
}

func ManifestUpdate() error {
  response, err := get("/Destiny2/Manifest/")
  if (err != nil) { return err }

  if response.ErrorCode.isError() {
    return errors.New(fmt.Sprintf("Error: %s", response.ErrorCode))
  }

  newManifest := Manifest{}
  err = mapstructure.Decode(response.Response, &newManifest)
  if err != nil { return err }

  if newManifest.Version == currentManifest.Version { return nil }
  err = manifestLocationValid()
  if err != nil { return err }

  err = manifestExtractToFile(newManifest.AssetContentPath, "/Assets.sqlite")
  if err != nil { return err }
  err = manifestExtractToFile(newManifest.ClanBannerDatabasePath, "/ClanBanner.sqlite")
  if err != nil { return err }

  for index, gearDatabase := range (newManifest.GearAssetDatabases) {
    outputPath := fmt.Sprintf("/GearDatabase-%d-%d.sqlite", index, gearDatabase.Version)
    err = manifestExtractToFile(gearDatabase.Path, outputPath)
    if err != nil { return err }
  }

  // Hmm, there's a bug here, only do english for now...
  /*
  for language, path := range (newManifest.WorldContentPaths) {
    outputPath := fmt.Sprintf("/WorldContent-%s.sqlite", language)
    err = manifestExtractToFile(path, outputPath)
    if err != nil { return err }
  }*/

  err = manifestExtractToFile(newManifest.WorldContentPaths["en"], "/WorldContent-en.sqlite")
  if err != nil { return err }

  fmt.Printf("Updating to Manifest: %s\n", newManifest.Version)

  return nil
}

func manifestLocationValid() error {
  if ManifestPath == "" { return errors.New("ManifestPath not set.") }
  var fileInfo os.FileInfo

  fileInfo, err := os.Stat(ManifestPath)
  if err != nil {
    errorString := fmt.Sprintf("ManifestPath '%s' does not exist.", ManifestPath)
    return errors.New(errorString)
  }
  if !fileInfo.IsDir() {
    errorString := fmt.Sprintf("ManifestPath '%s' is not a directory.", ManifestPath)
    return errors.New(errorString)
  }
  return nil
}

func test() {
  File, err := os.Open("world_sql_content_c423afdd29d2f6acb8ac2c859aeb4958.content")
  if err != nil { fmt.Printf("Couldn't open file! %s", err.Error()); return; }
  buf := make([]byte, 1358990)
  num, err1 := File.Read(buf)
  if err1 != nil { fmt.Printf("Couldn't read input file! %s", err1.Error()); return; }
  fmt.Printf("Read %d from file\n", num)

  bufReader := bytes.NewReader(buf)

  //var reader *zip.Reader
  _, err = zip.NewReader(bufReader, 1358990)
  if err != nil { fmt.Printf("Couldn't unzip file! %s", err.Error()); return; }

  fmt.Printf("Read zip file!")
}

func manifestExtractToFile(urlPart string, path string) error {
  response, err := http.DefaultClient.Get(serverBase + urlPart)
  if err != nil { return err }

  var buf []byte
  buf, err = ioutil.ReadAll(response.Body)
  if err != nil { return err }

  bufReader := bytes.NewReader(buf)

  var reader *zip.Reader
  reader, err = zip.NewReader(bufReader, response.ContentLength)
  if err != nil { return err }

  if len(reader.File) < 1 { return errors.New("Not enough files in zip!")}

  var zipOut io.ReadCloser
  zipOut, err = reader.File[0].Open()
  if err != nil { return err }

  var writeFile *os.File
  writeFile, err = os.Create(ManifestPath + path)
  if err != nil { return err }

  _, err = io.Copy(writeFile, zipOut)

  return nil
}
