package sacio

func (s *SACData) SetBody(series []float32, sampleRate int) error {
	s.B = F(0.0)
	s.E = F(float32(len(series)) / float32(sampleRate))

	s.SCALE = 1.0
	s.NPTS = N(len(series))
	s.DELTA = F(1.0 / float32(sampleRate))
	s.DEPMEN = F(getFloat32MeanValue(series))
	s.DEPMAX = F(getFloat32MaxValue(series))
	s.DEPMIN = F(getFloat32MinValue(series))

	s.IFTYPE = "itime"
	s.IZTYPE = "ib"
	s.LEVEN = true
	s.LPSPOL = true
	s.LOVROK = true
	s.LCALDA = false

	for _, v := range series {
		s.Body = append(s.Body, F(v))
	}
	return nil
}
