package goBungieNet

import (
  "reflect"
  "time"
  "strconv"

  "github.com/mitchellh/mapstructure"
)

func decodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
  if t == reflect.TypeOf(time.Time{}) && f.Kind() == reflect.String {
    return time.Parse(time.RFC3339, data.(string))
  }

  if t.Kind() == reflect.Int64 && f.Kind() == reflect.String {
    return strconv.ParseInt(data.(string), 10, 64)
  }

  if t.Kind() == reflect.Uint32 && f.Kind() == reflect.String {
    return strconv.ParseUint(data.(string), 10, 32)
  }

  return data, nil
}

func decode(source interface{}, target interface{}) error {
  config := mapstructure.DecoderConfig{
    DecodeHook: decodeHook,
    Result:     target,
  }

  decoder, err := mapstructure.NewDecoder(&config)
  if err != nil { return err }

  return decoder.Decode(source)
}
