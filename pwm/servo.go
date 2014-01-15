package pwm

import "time"

func servoPositionToDuty(position float32) time.Duration {
	return time.Duration(700e3 + uint32(position*1600e3+0.5))
}

type Servo struct {
	pwm      *PWM
	position float32
}

func NewServo(key string, position float32) (*Servo, error) {
	pwm, err := NewPWM(key, 2e7, servoPositionToDuty(position), POLARITY_LOW)
	if err != nil {
		return nil, err
	}
	servo := &Servo{
		pwm:      pwm,
		position: position,
	}
	return servo, nil
}

// Position returns the servo position in the range from 0.0 to 1.0
func (servo *Servo) Position() float32 {
	return servo.position
}

// SetPosition sets the servo position in the range from 0.0 to 1.0.
// position will be clamped if outside 0.0 to 1.0
func (servo *Servo) SetPosition(position float32) error {
	if position < 0 {
		position = 0
	} else if position > 1 {
		position = 1
	}
	err := servo.pwm.SetDuty(servoPositionToDuty(position))
	if err != nil {
		return err
	}
	servo.position = position
	return nil
}

// Closes the current servo instance
func (servo *Servo) Close() error {
	return servo.pwm.Close()
}