#include <stdio.h>
#include <cmath>

#include "gtest/gtest.h"

#include "calc.h"

using namespace std;

	// Returns n! (the factorial of n), n is a non-negative integer.
	int Factorial(int n) {

		if (n < 0) {
			throw std::invalid_argument("n should be a non-negative integer!");
		}

		int result = 1;
		for (int i = 1; i <= n; i++) {
			result *= i;
		}

		return result;
	}

	// Returns true if n is a prime number, n is a non-negative integer.
	bool IsPrime(int n) {
		// Trivial case 1: small numbers
		if (n <= 1) return false;

		// Trivial case 2: even numbers
		if (n % 2 == 0) return n == 2;

		// Now, we have that n is odd and n >= 3.
		// Try to divide n by every odd number i, starting from 3
		for (int i = 3; ; i += 2) {
			// We only have to try i up to the squre root of n
			if (i > n / i) break;

			// Now, we have i <= n/i < n.
			// If n is divisible by i, n is not prime.
			if (n % i == 0) return false;
		}

		// n has no integer factor in the range (1, n), and thus is prime.
		return true;
	}

	double SquareRoot(const double d) {
		if (d < 0) {
			throw std::invalid_argument("the square root of negatives are irrational");
		}

		return pow(d, 0.5);
	}

	int Divide(int dividend, int divisor) {
		if (divisor==0) {
			throw std::invalid_argument("divident can't be divided by 0.");
		}
		return dividend / divisor;
	}


////Begin to test
	TEST(Suite1, TestFactorial) {
		//  0<= n < INT_MAX
		INT_MAX;// maximum 32-bit integer

		EXPECT_EQ(1, Factorial(0));
		EXPECT_EQ(1, Factorial(1));
		EXPECT_EQ(6, Factorial(3));
		EXPECT_EQ(120, Factorial(5));

		EXPECT_THROW(Factorial(-1), std::invalid_argument);
//		ASSERT_NO_THROW(Factorial(5));
	}

//  How to test countless primes? 
	TEST(Suite1, TestIsPrime) {
		//  0<= n < INT_MAX
		EXPECT_EQ(false, IsPrime(1));
		EXPECT_EQ(true, IsPrime(2));
		EXPECT_EQ(true, IsPrime(3));
		EXPECT_EQ(false, IsPrime(4));
		EXPECT_EQ(true, IsPrime(5));
		EXPECT_EQ(true, IsPrime(17));
		EXPECT_EQ(true, IsPrime(29));
		EXPECT_EQ(true, IsPrime(443));
		EXPECT_EQ(true, IsPrime(2957));
		EXPECT_EQ(true, IsPrime(2939));
	}

	TEST(Suite1, DivideBy0) {

		EXPECT_EQ(Divide(10, 2), 5);
		EXPECT_EQ(Divide(10, -2), -5);
		EXPECT_EQ(Divide(-10, 2), -5);
		EXPECT_EQ(Divide(-10, -2), 5);

		int a = 0;
		int b = 0;
		//   1/0 = INFINITE and an exception occurs
		//EXPECT_NE(Divide(1, a - b),0);
		EXPECT_THROW(Divide(1, a-b), std::invalid_argument);
	}

	TEST(Suite2, Test_SquareRoot) {
		EXPECT_EQ(0.0, SquareRoot(0.0));
		EXPECT_EQ(18.0, SquareRoot(324.0));

		EXPECT_THROW(SquareRoot(-1), std::invalid_argument);
	}

TEST(Suite2, TestFloat) {
	float f0 = 99.9f;
	printf("f0 = %f\n", f0);

	float a = 10.550001f;
	float b = 10.550000f;
	float c =  0.000001f;
	float d = a - b;

	//EXPECT_EQ(a, b);
	EXPECT_FLOAT_EQ(a, b);
}

// 0.3 != 0.3 ?
TEST(Suite2, Oops) {
	double e = 0.15 + 0.15;
	double f = 0.1 + 0.2;

	//EXPECT_EQ(e, f);
	EXPECT_DOUBLE_EQ(e, f);
}

//Test whether a^0.5 * a^0.5 == a 
TEST(Suite2, TestDouble) {
	double d = SquareRoot(10);
	//EXPECT_EQ( d*d, 10) << "Surprise! d*d!=10.";
	EXPECT_DOUBLE_EQ( d*d, 10);
}

int main(int argc, char **argv) {
	::testing::InitGoogleTest(&argc, argv);

	return RUN_ALL_TESTS();
}