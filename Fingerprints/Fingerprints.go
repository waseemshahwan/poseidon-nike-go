package Fingerprints

import (
	tls "github.com/refraction-networking/utls"
)

// 8466c4390d4bc355
// uTLS does not support 0xAAAA as max version
func HASH_TOP_1() *tls.ClientHelloSpec {
	return &tls.ClientHelloSpec {
	CipherSuites: []uint16{
		tls.GREASE_PLACEHOLDER,
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		CompressionMethods: []byte{
			0x00, // compressionNone
		},
		Extensions: []tls.TLSExtension{
			&tls.UtlsGREASEExtension{},
			&tls.SNIExtension{},
			&tls.UtlsExtendedMasterSecretExtension{},
			&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient},
			&tls.SupportedCurvesExtension{[]tls.CurveID{
				tls.CurveID(tls.GREASE_PLACEHOLDER),
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
			}},
			&tls.SupportedPointsExtension{SupportedPoints: []byte{
				0x00, // pointFormatUncompressed
			}},
			&tls.SessionTicketExtension{},
			&tls.ALPNExtension{AlpnProtocols: []string{"h2","http/1.1"}},
			&tls.StatusRequestExtension{},
			&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256,
				tls.PSSWithSHA256,
				tls.PKCS1WithSHA256,
				tls.ECDSAWithP384AndSHA384,
				tls.PSSWithSHA384,
				tls.PKCS1WithSHA384,
				tls.PSSWithSHA512,
				tls.PKCS1WithSHA512,
			},},
			&tls.SCTExtension{},
			&tls.KeyShareExtension{[]tls.KeyShare{
				{Group: tls.CurveID(tls.GREASE_PLACEHOLDER), Data: []byte{0}},
				{Group: tls.X25519},
			}},
			&tls.PSKKeyExchangeModesExtension{[]uint8{
				tls.PskModeDHE,
			}},
			&tls.SupportedVersionsExtension{[]uint16{
				tls.GREASE_PLACEHOLDER,
				tls.VersionTLS13,
				tls.VersionTLS12,
				tls.VersionTLS11,
				tls.VersionTLS10,
			}},
			&tls.FakeCertCompressionAlgsExtension{[]tls.CertCompressionAlgo{
				tls.CertCompressionBrotli,
			}},
			&tls.GenericExtension {Id: 0x4469}, // WARNING: UNKNOWN EXTENSION, USE AT YOUR OWN RISK
			&tls.UtlsGREASEExtension{},
			&tls.UtlsPaddingExtension{GetPaddingLen: tls.BoringPaddingStyle},
		},
	}
}

// 133e933dd1dfea90
// uTLS does not support 0x2A2A as max version
// uTLS does not support 0xFAFA as max version
func HASH_TOP_2() *tls.ClientHelloSpec {
	return &tls.ClientHelloSpec {
		CipherSuites: []uint16{
			tls.GREASE_PLACEHOLDER,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.DISABLED_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.DISABLED_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.DISABLED_TLS_RSA_WITH_AES_256_CBC_SHA256,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			0xc008,
			tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
		CompressionMethods: []byte{
			0x00, // compressionNone
		},
		Extensions: []tls.TLSExtension{
			&tls.UtlsGREASEExtension{},
			&tls.SNIExtension{},
			&tls.UtlsExtendedMasterSecretExtension{},
			&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient},
			&tls.SupportedCurvesExtension{[]tls.CurveID{
				tls.CurveID(tls.GREASE_PLACEHOLDER),
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
				tls.CurveP521,
			}},
			&tls.SupportedPointsExtension{SupportedPoints: []byte{
				0x00, // pointFormatUncompressed
			}},
			&tls.ALPNExtension{AlpnProtocols: []string{"h2","http/1.1"}},
			&tls.StatusRequestExtension{},
			&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256,
				tls.PSSWithSHA256,
				tls.PKCS1WithSHA256,
				tls.ECDSAWithP384AndSHA384,
				tls.ECDSAWithSHA1,
				tls.PSSWithSHA384,
				tls.PSSWithSHA384,
				tls.PKCS1WithSHA384,
				tls.PSSWithSHA512,
				tls.PKCS1WithSHA512,
				tls.PKCS1WithSHA1,
			},},
			&tls.SCTExtension{},
			&tls.KeyShareExtension{[]tls.KeyShare{
				{Group: tls.CurveID(tls.GREASE_PLACEHOLDER), Data: []byte{0}},
				{Group: tls.X25519},
			}},
			&tls.PSKKeyExchangeModesExtension{[]uint8{
				tls.PskModeDHE,
			}},
			&tls.SupportedVersionsExtension{[]uint16{
				tls.GREASE_PLACEHOLDER,
				tls.VersionTLS13,
				tls.VersionTLS12,
				tls.VersionTLS11,
				tls.VersionTLS10,
			}},
			&tls.UtlsGREASEExtension{},
			&tls.UtlsPaddingExtension{GetPaddingLen: tls.BoringPaddingStyle},
		},
	}
}

// 256c61ee90fe805f
// uTLS does not support 0xEAEA as max version
// unsupported Curve in KeyShareExtension: 43690.To mimic it, fill the Data(key) field manually.
func HASH_TOP_3() *tls.ClientHelloSpec {
	return &tls.ClientHelloSpec {
		CipherSuites: []uint16{
			tls.GREASE_PLACEHOLDER,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		CompressionMethods: []byte{
			0x00, // compressionNone
		},
		Extensions: []tls.TLSExtension{
			&tls.UtlsGREASEExtension{},
			&tls.SNIExtension{},
			&tls.UtlsExtendedMasterSecretExtension{},
			&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient},
			&tls.SupportedCurvesExtension{[]tls.CurveID{
				tls.CurveID(tls.GREASE_PLACEHOLDER),
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
			}},
			&tls.SupportedPointsExtension{SupportedPoints: []byte{
				0x00, // pointFormatUncompressed
			}},
			&tls.SessionTicketExtension{},
			&tls.ALPNExtension{AlpnProtocols: []string{"h2","http/1.1"}},
			&tls.StatusRequestExtension{},
			&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256,
				tls.PSSWithSHA256,
				tls.PKCS1WithSHA256,
				tls.ECDSAWithP384AndSHA384,
				tls.PSSWithSHA384,
				tls.PKCS1WithSHA384,
				tls.PSSWithSHA512,
				tls.PKCS1WithSHA512,
			},},
			&tls.SCTExtension{},
			&tls.KeyShareExtension{[]tls.KeyShare{
				{Group: tls.CurveID(tls.GREASE_PLACEHOLDER), Data: []byte{0}},
				{Group: tls.X25519},
			}},
			&tls.PSKKeyExchangeModesExtension{[]uint8{
				tls.PskModeDHE,
			}},
			&tls.SupportedVersionsExtension{[]uint16{
				tls.GREASE_PLACEHOLDER,
				tls.VersionTLS13,
				tls.VersionTLS12,
				tls.VersionTLS11,
				tls.VersionTLS10,
			}},
			&tls.FakeCertCompressionAlgsExtension{[]tls.CertCompressionAlgo{
				tls.CertCompressionBrotli,
			}},
			&tls.GenericExtension {Id: 0x4469}, // WARNING: UNKNOWN EXTENSION, USE AT YOUR OWN RISK
			&tls.UtlsGREASEExtension{},
			&tls.GenericExtension {Id: 0x0029}, // WARNING: UNKNOWN EXTENSION, USE AT YOUR OWN RISK
		},
	}
}

// cd9f86e32edc16fb
// x509: certificate is valid for *.newrelic.com, not nike.com
func HASH_TOP_4() *tls.ClientHelloSpec {
	return &tls.ClientHelloSpec {
		CipherSuites: []uint16{
			tls.GREASE_PLACEHOLDER,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		CompressionMethods: []byte{
			0x00,
		},
		Extensions: []tls.TLSExtension{
			&tls.UtlsGREASEExtension{},
			&tls.SNIExtension{},
			&tls.UtlsExtendedMasterSecretExtension{},
			&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient},
			&tls.SupportedCurvesExtension{[]tls.CurveID{
				tls.CurveID(tls.GREASE_PLACEHOLDER),
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
			}},
			&tls.SupportedPointsExtension{SupportedPoints: []byte{
				0x00,
			}},
			&tls.SessionTicketExtension{},
			&tls.ALPNExtension{AlpnProtocols: []string{"h2", "http/1.1"}},
			&tls.StatusRequestExtension{},
			&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256,
				tls.PSSWithSHA256,
				tls.PKCS1WithSHA256,
				tls.ECDSAWithP384AndSHA384,
				tls.PSSWithSHA384,
				tls.PKCS1WithSHA384,
				tls.PSSWithSHA512,
				tls.PKCS1WithSHA512,
			}},
			&tls.SCTExtension{},
			&tls.KeyShareExtension{[]tls.KeyShare{
				{Group: tls.CurveID(tls.GREASE_PLACEHOLDER), Data: []byte{0}},
				{Group: tls.X25519},
			}},
			&tls.PSKKeyExchangeModesExtension{[]uint8{
				tls.PskModeDHE,
			}},
			&tls.SupportedVersionsExtension{[]uint16{
				tls.GREASE_PLACEHOLDER,
				tls.VersionTLS13,
				tls.VersionTLS12,
				tls.VersionTLS11,
				tls.VersionTLS10,
			}},
			&tls.FakeCertCompressionAlgsExtension{[]tls.CertCompressionAlgo{
				tls.CertCompressionBrotli,
			}},
			&tls.GenericExtension {Id: 0x4469},
			&tls.UtlsGREASEExtension{},
			&tls.UtlsPaddingExtension{GetPaddingLen: tls.BoringPaddingStyle},
		},
	}
}

func FIREFOX_PLAYWRIGHT() *tls.ClientHelloSpec {
	return &tls.ClientHelloSpec {
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
		CompressionMethods: []byte{
			0x00, // compressionNone
		},
		Extensions: []tls.TLSExtension{
			&tls.SNIExtension{},
			&tls.UtlsExtendedMasterSecretExtension{},
			&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient},
			&tls.SupportedCurvesExtension{[]tls.CurveID{
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
				tls.CurveP521,
				0x0100,
				0x0101,
			}},
			&tls.SupportedPointsExtension{SupportedPoints: []byte{
				0x00, // pointFormatUncompressed
			}},
			&tls.SessionTicketExtension{},
			&tls.ALPNExtension{AlpnProtocols: []string{"h2","http/1.1"}},
			&tls.StatusRequestExtension{},
			&tls.KeyShareExtension{[]tls.KeyShare{
				{Group: tls.X25519},
				{Group: tls.CurveP256},
			}},
			&tls.SupportedVersionsExtension{[]uint16{
				tls.VersionTLS13,
				tls.VersionTLS12,
			}},
			&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256,
				tls.ECDSAWithP384AndSHA384,
				tls.ECDSAWithP521AndSHA512,
				tls.PSSWithSHA256,
				tls.PSSWithSHA384,
				tls.PSSWithSHA512,
				tls.PKCS1WithSHA256,
				tls.PKCS1WithSHA384,
				tls.PKCS1WithSHA512,
				tls.ECDSAWithSHA1,
				tls.PKCS1WithSHA1,
			},},
			&tls.PSKKeyExchangeModesExtension{[]uint8{
				tls.PskModeDHE,
			}},
			&tls.FakeRecordSizeLimitExtension{16385},
			&tls.UtlsPaddingExtension{GetPaddingLen: tls.BoringPaddingStyle},
		},
	}
}