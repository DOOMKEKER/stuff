#include <iostream>
#include <string>
#include <vector>

#include "some_funcs.hpp"

using namespace std;

vector<string> std_read(void) {
  int strs, nms;
  char el;
  cin >> strs;
  vector<string> ret(strs);
  for (size_t i = 0; i < strs; i++) {
    cin >> ret[i];
  }
  return ret;
}

int main(void) {
  vector<string> vec = std_read();
  cout_vec(vec);

  return 0;
}
