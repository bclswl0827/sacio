package sacio

import (
	"fmt"
	"time"
)

func (s *SACData) SetTime(t time.Time) error {
	timeField, err := getParsedTimeField(t)
	if err != nil {
		return err
	}

	s.NVHDR = N(6) // Must be 6
	s.NZYEAR = N(timeField.Year)
	s.NZJDAY = N(timeField.Days)
	s.NZHOUR = N(timeField.Hour)
	s.NZMIN = N(timeField.Min)
	s.NZSEC = N(timeField.Sec)
	s.NZMSEC = N(timeField.Msec)
	return nil
}

func (s *SACData) SetInfo(network, station, location, channel string) error {
	var err error

	if len(network) > 8 {
		err = fmt.Errorf("network %s is too long", network)
		return err
	}

	if len(station) > 8 {
		err = fmt.Errorf("station %s is too long", station)
		return err
	}

	if len(location) > 8 {
		err = fmt.Errorf("location %s is too long", location)
		return err
	}

	if len(channel) > 8 {
		err = fmt.Errorf("channel %s is too long", channel)
		return err
	}

	s.KNETWK = K(network)
	s.KSTNM = K(station)
	s.KHOLE = K(location)
	s.KCMPNM = K(channel)
	return nil
}
