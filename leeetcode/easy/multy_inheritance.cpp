#include <iostream>
#include <vector>

using namespace std;

int climbStairs(int n) {
  int n_1 = 0;
  int n_2 = 1;
  int ret = 0;
  for (int i = 0; i < n; ++i) {
    ret = n_1 + n_2;
    n_1 = n_2;
    n_2 = ret;
  }
  return ret;
}

int main(int argc, char const *argv[]) {
  cout << climbStairs(atoi(argv[1]));
  return 0;
}
