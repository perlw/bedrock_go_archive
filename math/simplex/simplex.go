package simplex

import "math"

var gradient = [12][3]float64{
	{1.0, 1.0, 0.0}, {-1.0, 1.0, 0.0}, {1.0, -1.0, 0.0}, {-1.0, -1.0, 0.0},
	{1.0, 0.0, 1.0}, {-1.0, 0.0, 1.0}, {1.0, 0.0, -1.0}, {-1.0, 0.0, -1.0},
	{0.0, 1.0, 1.0}, {0.0, -1.0, 1.0}, {0.0, 1.0, -1.0}, {0.0, -1.0, -1.0},
}

var permutation = [512]int{
	151, 160, 137, 91, 90,
	15, 131, 13, 201, 95,
	96, 53, 194, 233, 7,
	225, 140, 36, 103, 30,
	69, 142, 8, 99, 37,
	240, 21, 10, 23, 190,
	6, 148, 247, 120, 234,
	75, 0, 26, 197, 62,
	94, 252, 219, 203, 117,
	35, 11, 32, 57, 177,
	33, 88, 237, 149, 56,
	87, 174, 20, 125, 136,
	171, 168, 68, 175, 74,
	165, 71, 134, 139, 48,
	27, 166, 77, 146, 158,
	231, 83, 111, 229, 122,
	60, 211, 133, 230, 220,
	105, 92, 41, 55, 46,
	245, 40, 244, 102, 143,
	54, 65, 25, 63, 161,
	1, 216, 80, 73, 209,
	76, 132, 187, 208, 89,
	18, 169, 200, 196, 135,
	130, 116, 188, 159, 86,
	164, 100, 109, 198, 173,
	186, 3, 64, 52, 217, 226,
	250, 124, 123, 5, 202,
	38, 147, 118, 126, 255,
	82, 85, 212, 207, 206,
	59, 227, 47, 16, 58,
	17, 182, 189, 28, 42,
	223, 183, 170, 213, 119,
	248, 152, 2, 44, 154,
	163, 70, 221, 153, 101,
	155, 167, 43, 172, 9,
	129, 22, 39, 253, 19,
	98, 108, 110, 79, 113,
	224, 232, 178, 185, 112,
	104, 218, 246, 97, 228,
	251, 34, 242, 193, 238,
	210, 144, 12, 191, 179,
	162, 241, 81, 51, 145,
	235, 249, 14, 239, 107,
	49, 192, 214, 31, 181,
	199, 106, 157, 184, 84,
	204, 176, 115, 121, 50,
	45, 127, 4, 150, 254,
	138, 236, 205, 93, 222,
	114, 67, 29, 24, 72, 243,
	141, 128, 195, 78, 66,
	215, 61, 156, 180, 151,
	160, 137, 91, 90, 15,
	131, 13, 201, 95, 96,
	53, 194, 233, 7, 225,
	140, 36, 103, 30, 69,
	142, 8, 99, 37, 240,
	21, 10, 23, 190, 6,
	148, 247, 120, 234, 75,
	0, 26, 197, 62, 94,
	252, 219, 203, 117, 35,
	11, 32, 57, 177, 33,
	88, 237, 149, 56, 87,
	174, 20, 125, 136, 171,
	168, 68, 175, 74, 165,
	71, 134, 139, 48, 27,
	166, 77, 146, 158, 231,
	83, 111, 229, 122, 60,
	211, 133, 230, 220, 105,
	92, 41, 55, 46, 245,
	40, 244, 102, 143, 54,
	65, 25, 63, 161, 1,
	216, 80, 73, 209, 76,
	132, 187, 208, 89, 18,
	169, 200, 196, 135, 130,
	116, 188, 159, 86, 164,
	100, 109, 198, 173, 186,
	3, 64, 52, 217, 226,
	250, 124, 123, 5, 202,
	38, 147, 118, 126, 255,
	82, 85, 212, 207, 206,
	59, 227, 47, 16, 58,
	17, 182, 189, 28, 42,
	223, 183, 170, 213, 119,
	248, 152, 2, 44, 154,
	163, 70, 221, 153, 101,
	155, 167, 43, 172, 9,
	129, 22, 39, 253, 19,
	98, 108, 110, 79, 113,
	224, 232, 178, 185, 112,
	104, 218, 246, 97, 228,
	251, 34, 242, 193, 238,
	210, 144, 12, 191, 179,
	162, 241, 81, 51, 145,
	235, 249, 14, 239, 107,
	49, 192, 214, 31, 181,
	199, 106, 157, 184, 84,
	204, 176, 115, 121, 50,
	45, 127, 4, 150, 254,
	138, 236, 205, 93, 222,
	114, 67, 29, 24, 72,
	243, 141, 128, 195, 78,
	66, 215, 61, 156, 180,
}

func dot(x, y, z float64, g [3]float64) float64 {
	return x*g[0] + y*g[1] + z*g[2]
}

func Noise(xin, yin, zin float64) float64 {
	var F3, G3, t, X0, Y0, Z0, x0, y0, z0, s, x1, y1, z1, x2, y2, z2, x3, y3, z3, t0, t1, t2, t3, n0, n1, n2, n3 float64
	var i, j, k, ii, jj, kk, i1, j1, k1, i2, j2, k2, gi0, gi1, gi2, gi3 int

	F3 = 1.0 / 3.0
	s = (xin + yin + zin) * F3
	i = int(xin + s)
	j = int(yin + s)
	k = int(zin + s)
	G3 = 1.0 / 6.0
	t = float64(i+j+k) * G3
	X0 = float64(i) - t
	Y0 = float64(j) - t
	Z0 = float64(k) - t
	x0 = xin - X0
	y0 = yin - Y0
	z0 = zin - Z0

	if x0 >= y0 {
		if y0 >= z0 {
			i1 = 1
			j1 = 0
			k1 = 0
			i2 = 1
			j2 = 1
			k2 = 0
		} else if x0 >= z0 {
			i1 = 1
			j1 = 0
			k1 = 0
			i2 = 1
			j2 = 0
			k2 = 1
		} else {
			i1 = 0
			j1 = 0
			k1 = 1
			i2 = 1
			j2 = 0
			k2 = 1
		}
	} else {
		if y0 < z0 {
			i1 = 0
			j1 = 0
			k1 = 1
			i2 = 0
			j2 = 1
			k2 = 1
		} else if x0 < z0 {
			i1 = 0
			j1 = 1
			k1 = 0
			i2 = 0
			j2 = 1
			k2 = 1
		} else {
			i1 = 0
			j1 = 1
			k1 = 0
			i2 = 1
			j2 = 1
			k2 = 0
		}
	}

	x1 = x0 - float64(i1) + G3
	y1 = y0 - float64(j1) + G3
	z1 = z0 - float64(k1) + G3
	x2 = x0 - float64(i2) + 2.0*G3
	y2 = y0 - float64(j2) + 2.0*G3
	z2 = z0 - float64(k2) + 2.0*G3
	x3 = x0 - 1.0 + 3.0*G3
	y3 = y0 - 1.0 + 3.0*G3
	z3 = z0 - 1.0 + 3.0*G3

	ii = i & 255
	jj = j & 255
	kk = k & 255

	gi0 = permutation[ii+permutation[jj+permutation[kk]]] % 12
	gi1 = permutation[ii+i1+permutation[jj+j1+permutation[kk+k1]]] % 12
	gi2 = permutation[ii+i2+permutation[jj+j2+permutation[kk+k2]]] % 12
	gi3 = permutation[ii+1+permutation[jj+1+permutation[kk+1]]] % 12

	t0 = 0.6 - x0*x0 - y0*y0 - z0*z0
	if t0 < 0 {
		n0 = 0.0
	} else {
		t0 *= t0
		n0 = t0 * t0 * dot(x0, y0, z0, gradient[gi0])
	}

	t1 = 0.6 - x1*x1 - y1*y1 - z1*z1
	if t1 < 0 {
		n1 = 0.0
	} else {
		t1 *= t1
		n1 = t1 * t1 * dot(x1, y1, z1, gradient[gi1])
	}

	t2 = 0.6 - x2*x2 - y2*y2 - z2*z2
	if t2 < 0 {
		n2 = 0.0
	} else {
		t2 *= t2
		n2 = t2 * t2 * dot(x2, y2, z2, gradient[gi2])
	}

	t3 = 0.6 - x3*x3 - y3*y3 - z3*z3
	if t3 < 0 {
		n3 = 0.0
	} else {
		t3 *= t3
		n3 = t3 * t3 * dot(x3, y3, z3, gradient[gi3])
	}

	return 16.0*(n0+n1+n2+n3) + 1.0
}

func NoiseOctave(octaves int, x, y, z float64) float64 {
	value := 0.0

	for t := 0; t < octaves; t++ {
		pow := math.Pow(2, float64(t))
		value += Noise(x*pow, y*pow, z*pow)
	}

	return value
}
