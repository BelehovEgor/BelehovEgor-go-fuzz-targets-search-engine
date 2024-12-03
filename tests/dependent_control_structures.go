package tests

func if_true() bool {
	if true || false {
		return true
	}

	return false
}

func if_func(n int) bool {
	if n > 0 {
		return true
	}

	return false
}

func for_func(n int) bool {
	for i := 0; i < n; i++ {
		return true
	}

	return false
}

func switch_func(n int) bool {
	switch n {
	case 1, 3:
		return true
	}

	return false
}

func if_depend_vars(a, b int) bool {
	c := a + 1
	d := b + a
	if c+d > 0 {
		return true
	}

	return false
}

func mixed(a, b int) bool {
	c := a + 1
	d := b + a
	if c > 0 {
		if true {
		}
		if false {
		}

		switch d {
		}

		return true
	}

	for i := 0; i < a; i++ {
	}

	return false
}
