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
  "strings"

  _ "github.com/mattn/go-sqlite3"
  "github.com/mitchellh/mapstructure"
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

type manifestInfo struct {
  Version string
  Path string
  DBs map[string]*sql.DB
}

type GearAssetDatabase struct {
  Version int32
  Path string
}

var (
  manifest manifestInfo
)

func init() {
  manifest = manifestInfo{
    Version: "None",
    Path: "",
    DBs: make(map[string]*sql.DB),
  }
}

func SetManifestPath(path string) error {
  if path == manifest.Path { return nil }
  manifest.Path = path

  err := loadManifestVersion()
  return err
}

func ManifestVersion() string {
  return manifest.Version
}

func ManifestUpdate() error {
  response, err := get("/Destiny2/Manifest/")
  if (err != nil) { return err }

  if response.ErrorCode.isError() {
    return errors.New(fmt.Sprintf("Error: %s", response.ErrorCode))
  }

  newManifest := manifestData{}
  err = mapstructure.Decode(response.Response, &newManifest)
  if err != nil { return err }

  if newManifest.Version == manifest.Version { return nil }
  err = manifestLocationValid()
  if err != nil { return err }

  err = manifestExtractToFile(newManifest.AssetContentPath, manifestAssets)
  if err != nil { return err }
  err = manifestExtractToFile(newManifest.ClanBannerDatabasePath, manifestClanBanner)
  if err != nil { return err }

  for _, gearDatabase := range (newManifest.GearAssetDatabases) {
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

  outputPath := fmt.Sprintf("%s-%s", manifestWorld, "en")
  err = manifestExtractToFile(newManifest.WorldContentPaths["en"], outputPath)
  if err != nil { return err }

  manifest.Version = newManifest.Version

  fmt.Printf("Manifest Version is now: %s", manifest.Version)

  return saveManifestVersion()
}

func manifestLocationValid() error {
  if manifest.Path == "" { return errors.New("ManifestPath not set.") }
  var fileInfo os.FileInfo

  fileInfo, err := os.Stat(manifest.Path)
  if err != nil {
    errorString := fmt.Sprintf("ManifestPath '%s' does not exist.", manifest.Path)
    return errors.New(errorString)
  }
  if !fileInfo.IsDir() {
    errorString := fmt.Sprintf("ManifestPath '%s' is not a directory.", manifest.Path)
    return errors.New(errorString)
  }
  return nil
}

func loadManifestVersion() error {
  err := manifestLocationValid()
  if err != nil { return err }

  manifestVersionPath := manifest.Path + "/Version"

  var fileInfo os.FileInfo
  fileInfo, err = os.Stat(manifestVersionPath)
  if err != nil {
    manifest.Version = "None"
    return nil
  }

  if fileInfo.IsDir() {
    errorString := fmt.Sprintf("'%s' - is not a file", manifestVersionPath)
    return errors.New(errorString)
  }

  var data []byte
  data, err = ioutil.ReadFile(manifestVersionPath)
  if err != nil {
    errorString := fmt.Sprintf("Could not read manifest version file '%s'", manifestVersionPath)
    return errors.New(errorString)
  }

  manifest.Version = strings.Trim(string(data), " \n")
  return nil
}

func saveManifestVersion() error {
  err := manifestLocationValid()
  if err != nil { return err }

  manifestVersionPath := manifest.Path + "/Version"

  var fileInfo os.FileInfo
  fileInfo, err = os.Stat(manifestVersionPath)
  if err == nil {
    if fileInfo.IsDir() {
      errorString := fmt.Sprintf("'%s' - is not a file", manifestVersionPath)
      return errors.New(errorString)
    }
  }

  return ioutil.WriteFile(manifestVersionPath, []byte(manifest.Version), os.FileMode(0600))
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
  writeFilePath := fmt.Sprintf("%s/%s.sqlite", manifest.Path, path)
  writeFile, err = os.Create(writeFilePath)
  if err != nil { return err }

  _, err = io.Copy(writeFile, zipOut)

  return nil
}

func manifestOpenData(dataFile string) (*sql.DB,error) {
  if db, ok := manifest.DBs[dataFile]; ok {
    return db, nil
  }

  dbPath := fmt.Sprintf("%s/%s.sqlite", manifest.Path,dataFile)
  db, err := sql.Open("sqlite3", dbPath)
  if err != nil { return nil, err }

  manifest.DBs[dataFile] = db
  return db, nil
}
