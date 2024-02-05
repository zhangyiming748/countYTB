package opt

func Count(i float64) float64 {
	toBit := i * 1000    // bps
	toByte := toBit / 8  // Bps
	b2k := toByte / 1024 // KBps
	k2m := b2k / 1024    // MBps
	return k2m
}
