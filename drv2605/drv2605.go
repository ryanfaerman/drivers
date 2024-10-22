package drv2605

import (
	"tinygo.org/x/drivers"
)

// The default address for the DRV2605 according to its datasheet.
var Address uint16 = 0x5A

type Device struct {
	bus  drivers.I2C
	addr uint16
}

func New(bus drivers.I2C, addr uint16) *Device {
	d := &Device{bus: bus, addr: addr}
	return d
}

// Init initializes the device. This must be called before any other methods
// and cannot be called until after 250ms from power up to allow the
// driver to boot.
func (d *Device) Init() error {
	err := d.WriteRegister(DRV2605_REG_MODE, []byte{0x00}) // exit standby mode
	if err != nil {
		return err
	}

	d.WriteRegister(DRV2605_MODE_INTTRIG, []byte{0x00})   // internal trigger mode
	d.WriteRegister(DRV2605_REG_RTPIN, []byte{0x00})      // no real-time playback
	d.WriteRegister(DRV2605_REG_OVERDRIVE, []byte{0x00})  // no overdrive
	d.WriteRegister(DRV2605_REG_SUSTAINPOS, []byte{0x00}) // no positive sustain
	d.WriteRegister(DRV2605_REG_SUSTAINNEG, []byte{0x00}) // no negative sustain
	d.WriteRegister(DRV2605_REG_BREAK, []byte{0x00})      // no brake
	d.WriteRegister(DRV2605_REG_AUDIOMAX, []byte{0x64})

	d.UseERM()

	control, err := d.ReadRegister(DRV2605_REG_CONTROL3)
	if err != nil {
		return err
	}
	d.WriteRegister(DRV2605_REG_CONTROL3, []byte{control[0] & 0x20})

	return nil
}

// WriteRegister writes a register to the device.
func (d *Device) WriteRegister(reg byte, data []byte) error {
	buf := make([]uint8, len(data)+1)
	buf[0] = reg
	copy(buf[1:], data)
	return d.bus.Tx(uint16(d.addr), buf, nil)
}

// ReadRegister reads a register from the device.
func (d *Device) ReadRegister(reg byte) ([]byte, error) {
	data := []byte{0}
	err := d.bus.Tx(uint16(d.addr), []byte{reg}, data)
	return data, err
}

// Set the device mode, where mode is defined in section 7.6.2 of
// the datasheet. See: http://www.adafruit.com/datasheets/DRV2605.pdf
//
// 0: Internal trigger, call go() to start playback
// 1: External trigger, rising edge on IN pin starts playback
// 2: External trigger, playback follows the state of IN pin
// 3: PWM/analog input
// 4: Audio
// 5: Real-time playback
// 6: Diagnostics
// 7: Auto calibration
func (d *Device) SetMode(mode byte) error {
	return d.WriteRegister(DRV2605_REG_MODE, []byte{mode})
}

// SetWaveform sets the waveform for the given slot, which can be 0-7.
// The waveform can be 1-123, these are defined in the datasheet and
// linked in the README for this package.
func (d *Device) SetWaveform(slot byte, waveform byte) error {
	return d.WriteRegister(0x04+slot, []byte{waveform})
}

// Go starts the effect playing via the internal trigger.
func (d *Device) Go() error {
	return d.WriteRegister(0x0C, []byte{0x01}) // set 'GO' bit to start
}

// Stop the currently playing effect.
func (d *Device) Stop() error {
	return d.WriteRegister(0x0C, []byte{0x00}) // clear 'GO' bit to stop
}

// Set the realtime value when in RTP mode, used to directly drive
// the haptic motor.
func (d *Device) SetRealtimeValue(value byte) error {
	return d.WriteRegister(DRV2605_REG_RTPIN, []byte{value})
}

// Use an ERM (eccentric rotating mass) motor. This is the default.
func (d *Device) UseERM() error {
	feedback, err := d.ReadRegister(DRV2605_REG_FEEDBACK)
	if err != nil {
		return err
	}
	return d.WriteRegister(DRV2605_REG_FEEDBACK, []byte{feedback[0] & 0x7F})
}

// Use an LRA (linear resonant actuator) motor.
func (d *Device) UseLRA() error {
	feedback, err := d.ReadRegister(DRV2605_REG_FEEDBACK)
	if err != nil {
		return err
	}
	return d.WriteRegister(DRV2605_REG_FEEDBACK, []byte{feedback[0] & 0x80})
}
