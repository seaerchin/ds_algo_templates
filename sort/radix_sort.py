from math import inf
import unittest

# radix sort is written in python instead of golang due to python's ease of passing around anonymous functions 
# another reason for doing so is due to python's sorted function being a higher order function, accepting a key parameter 
# k represents the number of digits in the largest number
def radix_sort(arr, k):
    # NOTE: this version of radix sort doesn't run in o(n) time - this is because the intermediate sort used by python is not a o(n) sort
    # however, if this were to use counting sort for the intermediate sort, this would run in o(n) time
    for i in range(-1, -k, -1):
        arr = sorted(arr, key = lambda x: int(str(x)[i]) if abs(i) < len(str(x)) else -inf)
    return arr


class TestRadixSort(unittest.TestCase):
    def test_radix_sort(self):
        arr = [5, 33, 4431, 123, 20]
        self.assertEqual(radix_sort(arr, 4), sorted(arr))

if __name__ == '__main__':
    unittest.main()