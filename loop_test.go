package loop

import "testing"

func TestSumOfFirst(t *testing.T) {
	given := 3
	want := 6

	result := sumOfFirst(given)
	if result != want {
		t.Errorf("sumOfFirst(%d) = %d, want %d", given, result, want)
	}
}

func TestIsPrime(t *testing.T) {
	testcases := []struct {
		given int
		want  bool
	}{
		{given: 1, want: false},
		{given: 2, want: true},
		{given: 3, want: true},
		{given: 4, want: false},
	}

	for _, tc := range testcases {
		t.Run("prime", func(t *testing.T) {

			given := tc.given
			want := tc.want

			result := isPrime(given)
			if result != want {
				t.Errorf("isPrime(%d) = %t, want %t", given, result, want)
			}
		})
	}

}
