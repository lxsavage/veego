package veego

type DeviceCapabilityType string
type DeviceType string

const (
	apiV1 = "https://openapi.api.govee.com/router/api/v1"
)

const (
	TypeOnOff               DeviceCapabilityType = "devices.capabilities.on_off"
	TypeRange               DeviceCapabilityType = "devices.capabilities.range"
	TypeColorSetting        DeviceCapabilityType = "devices.capabilities.color_setting"
	TypeWorkMode            DeviceCapabilityType = "devices.capabilities.work_mode"
	TypeSegmentColorSetting DeviceCapabilityType = "devices.capabilities.segment_color_setting"
	TypeDynamicScene        DeviceCapabilityType = "devices.capabilities.dynamic_scene"
	TypeTemperatureSetting  DeviceCapabilityType = "devices.capabilities.temperature_setting"
)

const (
	DeviceLight         DeviceType = "devices.types.light"
	DeviceAirPurifier   DeviceType = "devices.types.air_purifier"
	DeviceThermometer   DeviceType = "devices.types.thermometer"
	DeviceSocket        DeviceType = "devices.types.socket"
	DeviceSensor        DeviceType = "devices.types.sensor"
	DeviceHeater        DeviceType = "devices.types.heater"
	DeviceHumidifier    DeviceType = "devices.types.humidifier"
	DeviceDehumidifier  DeviceType = "devices.types.dehumidifier"
	DeviceIceMaker      DeviceType = "devices.types.ice_maker"
	DeviceArmoaDiffuser DeviceType = "devices.types.aroma_diffuser"
	DeviceBox           DeviceType = "devices.types.box"
)
