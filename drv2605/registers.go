package drv2605

const (
	// Register Addresses
	DRV2605_REG_STATUS      = 0x00 // Status register
	DRV2605_REG_MODE        = 0x01 // Mode register
	DRV2605_REG_RTPIN       = 0x02 // Real-time playback input register
	DRV2605_REG_LIBRARY     = 0x03 // Waveform library selection register
	DRV2605_REG_WAVESEQ1    = 0x04 // Waveform sequence register 1
	DRV2605_REG_WAVESEQ2    = 0x05 // Waveform sequence register 2
	DRV2605_REG_WAVESEQ3    = 0x06 // Waveform sequence register 3
	DRV2605_REG_WAVESEQ4    = 0x07 // Waveform sequence register 4
	DRV2605_REG_WAVESEQ5    = 0x08 // Waveform sequence register 5
	DRV2605_REG_WAVESEQ6    = 0x09 // Waveform sequence register 6
	DRV2605_REG_WAVESEQ7    = 0x0A // Waveform sequence register 7
	DRV2605_REG_WAVESEQ8    = 0x0B // Waveform sequence register 8
	DRV2605_REG_GO          = 0x0C // Go register
	DRV2605_REG_OVERDRIVE   = 0x0D // Overdrive time offset register
	DRV2605_REG_SUSTAINPOS  = 0x0E // Sustain time offset, positive register
	DRV2605_REG_SUSTAINNEG  = 0x0F // Sustain time offset, negative register
	DRV2605_REG_BREAK       = 0x10 // Brake time offset register
	DRV2605_REG_AUDIOCTRL   = 0x11 // Audio-to-vibe control register
	DRV2605_REG_AUDIOLVL    = 0x12 // Audio-to-vibe minimum input level register
	DRV2605_REG_AUDIOMAX    = 0x13 // Audio-to-vibe maximum input level register
	DRV2605_REG_AUDIOOUTMIN = 0x14 // Audio-to-vibe minimum output drive register
	DRV2605_REG_AUDIOOUTMAX = 0x15 // Audio-to-vibe maximum output drive register
	DRV2605_REG_RATEDV      = 0x16 // Rated voltage register
	DRV2605_REG_CLAMPV      = 0x17 // Overdrive clamp voltage register
	DRV2605_REG_AUTOCALCOMP = 0x18 // Auto-calibration compensation result register
	DRV2605_REG_AUTOCALEMP  = 0x19 // Auto-calibration back-EMF result register
	DRV2605_REG_FEEDBACK    = 0x1A // Feedback control register
	DRV2605_REG_CONTROL1    = 0x1B // Control1 Register
	DRV2605_REG_CONTROL2    = 0x1C // Control2 Register
	DRV2605_REG_CONTROL3    = 0x1D // Control3 Register
	DRV2605_REG_CONTROL4    = 0x1E // Control4 Register
	DRV2605_REG_VBAT        = 0x21 // Vbat voltage-monitor register
	DRV2605_REG_LRARESON    = 0x22 // LRA resonance-period register

	// Mode Values
	DRV2605_MODE_INTTRIG     = 0x00 // Internal trigger mode
	DRV2605_MODE_EXTTRIGEDGE = 0x01 // External edge trigger mode
	DRV2605_MODE_EXTTRIGLVL  = 0x02 // External level trigger mode
	DRV2605_MODE_PWMANALOG   = 0x03 // PWM/Analog input mode
	DRV2605_MODE_AUDIOVIBE   = 0x04 // Audio-to-vibe mode
	DRV2605_MODE_REALTIME    = 0x05 // Real-time playback (RTP) mode
	DRV2605_MODE_DIAGNOS     = 0x06 // Diagnostics mode
	DRV2605_MODE_AUTOCAL     = 0x07 // Auto calibration mode
)
