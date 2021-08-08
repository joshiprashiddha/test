package main

import (
	"fmt"
)

func calcAngle(h int64, m int64) (err error) {
	if h < 0 && m < 0 {
		return fmt.Errorf("Only positive int can be processed")
	}

	if m >= 60 {
	    m = m / 60
		h = h + m%60
	}

	if h >= 12 {
		h = h / 12
	}

	total_hour := float64(h + m/60)
	h_angle := (360 / 12) * total_hour
	m_angle := (360 / 60) * m
	var angle_diff float64
	if h_angle > float64(m_angle) {
		angle_diff = h_angle - float64(m_angle)
	} else {
		angle_diff = float64(m_angle) - h_angle
	}

	if angle_diff >= 180 {
		angle_diff = 360 - angle_diff
	}

	fmt.Println(angle_diff)
	return nil
}

func main() {
	err := calcAngle(2, 24)
	if err != nil {
		fmt.Println("Error Occured", err)
	}
}
