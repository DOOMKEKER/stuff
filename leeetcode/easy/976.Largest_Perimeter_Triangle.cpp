#include <algorithm>
#include <iostream>
#include <vector>

#include "/home/amon/projects/stuff/leeetcode/some_funcs.hpp"

using namespace std;

class Solution {
 public:
  int largestPerimeter(vector<int>& nums) {
    sort(nums.begin(), nums.end());
    int size = nums.size();
    int s = nums[size];
    for (int i = size - 1; i > 1; --i) {
      if (nums[i - 2] + nums[i - 1] > nums[i]) {
        return nums[i - 2] + nums[i - 1] + nums[i];
      }
    }

    return 0;
  }
};

int main(void) {
  Solution sol;
  vector<int> num = {2, 1, 2};
  cout << sol.largestPerimeter(num);
  return 0;
}