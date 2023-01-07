#include <iostream>
#include <vector>

#include "../some_funcs.hpp"

using namespace std;

class Solution {
 public:
  int sum(vector<int>& nums, int i) {
    int sum = 0;
    if ((i + 1) > nums.size() - 1) return 0;
    for (int j = i + 1; j < nums.size(); j++) {
      sum += nums[j];
    }

    return sum;
  }

  int pivotIndex(vector<int>& nums) {
    int index = 0;
    int left = 0;
    int right = sum(nums, index);

    for (; index < nums.size();) {
      if ((left - right) == 0) return index;
      left += nums[index];
      if (++index > nums.size() - 1) continue;
      right -= nums[index];
    }

    return -1;
  }
};

int main() {
  Solution sol;

  vector<int> nums = {-1,-1,-1,1,1,1};

  auto ans = sol.pivotIndex(nums);
  cout << ans;

  return 0;
}