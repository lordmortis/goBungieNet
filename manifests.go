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
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
//  "github.com/mitchellh/mapstructure"
)

const manifestAssets = "Assets"
const manifestClanBanner = "ClanBanner"
const manifestGear = "Gear"
const manifestWorld = "WorldContent"

type manifestData struct {
  Version string
  AssetContentPath string `mapstructure:"mobileAssetContentPath"`
  ClanBannerDatabasePath string `mapstructure:"mobileClanBannerDatabasePath"`
  GearAssetDatabases []GearAssetDatabase `mapstructure:"mobileGearAssetDataBases"`
  GearCDN map[string]string `mapstructure:"mobileGearCDN"`
  WorldContentPaths map[string]string `mapstructure:"mobileWorldContentPaths"`
}

type DestinyDisplayProperties struct {
  Description string
  Name string
  icon string
  HasIcon bool
}

type DestinyItemQuantity struct {
  ItemHash uint32
  ItemInstanceID int64
  Quantity int32
}

type manifest struct {
  Version string
}

type GearAssetDatabase struct {
  Version int32
  Path string
}

var (
  currentManifest manifest
  ManifestPath string
  manifestDBs map[string]*sql.DB
)

func init() {
  currentManifest.Version = "No Manifest"
  ManifestPath = ""
  manifestDBs = make(map[string]*sql.DB)
}

func ManifestUpdate() error {
/*  response, err := get("/Destiny2/Manifest/")
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

  err = manifestExtractToFile(newManifest.AssetContentPath, manifestAssets)
  if err != nil { return err }
  err = manifestExtractToFile(newManifest.ClanBannerDatabasePath, manifestClanBanner)
  if err != nil { return err }

  for index, gearDatabase := range (newManifest.GearAssetDatabases) {
    outputPath := fmt.Sprintf("%s-%d", manifestGear, gearDatabase.Version)
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

  /*
  outputPath := fmt.Sprintf("%s-%s", manifestWorld, "en")
  err = manifestExtractToFile(newManifest.WorldContentPaths["en"], outputPath)
  if err != nil { return err }*/

  //fmt.Printf("Updating to Manifest: %s\n", newManifest.Version)

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
  writeFilePath := fmt.Sprintf("%s/%s.sqlite", ManifestPath, path)
  writeFile, err = os.Create(writeFilePath)
  if err != nil { return err }

  _, err = io.Copy(writeFile, zipOut)

  return nil
}

func manifestOpenData(dataFile string) (*sql.DB,error) {
  if db, ok := manifestDBs[dataFile]; ok {
    return db, nil
  }

  dbPath := fmt.Sprintf("%s/%s.sqlite", ManifestPath,dataFile)
  db, err := sql.Open("sqlite3", dbPath)
  if err != nil { return nil, err }

  manifestDBs[dataFile] = db
  return db, nil
}
