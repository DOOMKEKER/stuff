#include <iostream>
#include <tuple>
#include <vector>

#include "../some_funcs.hpp"

using namespace std;

class Solution {
 public:
  vector<int> findDiagonalOrder(vector<vector<int>>& nums) {
    vector<vector<int>> diagonals;
    vector<int> ret;

    for (size_t i = 0; i < nums.size(); i++) {
      for (size_t j = 0; j < nums[i].size(); j++) {
        if (diagonals.size() == i + j) diagonals.emplace_back(vector<int>());
        diagonals[i + j].emplace_back(nums[i][j]);
      }
    }

    for (auto& el : diagonals) {
      ret.insert(ret.end(), el.rbegin(), el.rend());
    }

    return ret;
  }
};

int main(void) {
  int a = 0;
  Solution sol;

  vector<vector<int>> v{{1, 2, 3}, {2, 3, 4}, {1, 3, 5}};

  vector<int> ans = sol.findDiagonalOrder(v);

  cout_vec(ans);
  return 0;
};
