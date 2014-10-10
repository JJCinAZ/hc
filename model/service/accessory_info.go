package service

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/characteristic"
)

type AccessoryInfo struct {
    *Service
    
    Identify *characteristic.Identify
    Serial *characteristic.SerialNumber
    Model *characteristic.Model
    Manufacturer *characteristic.Manufacturer
    Name *characteristic.Name
}

func NewInfo(info model.Info) *AccessoryInfo {
    return NewAccessoryInfo(info.Name, info.Serial, info.Manufacturer, info.Model)
}
    
func NewAccessoryInfo(accessoryName, serialNumber, manufacturerName, modelName string) *AccessoryInfo {
    identify        := characteristic.NewIdentify(false)
    serial          := characteristic.NewSerialNumber(serialNumber)
    model           := characteristic.NewModel(modelName)
    manufacturer    := characteristic.NewManufacturer(manufacturerName)
    name            := characteristic.NewName(accessoryName)
    
    service := NewService()
    service.Type = TypeAccessoryInfo
    service.AddCharacteristic(identify)
    service.AddCharacteristic(serial)
    service.AddCharacteristic(model)
    service.AddCharacteristic(manufacturer)
    service.AddCharacteristic(name)
    
    return &AccessoryInfo{service, identify, serial, model, manufacturer, name}
}