#include <string>
#include <iostream>

class Solution {
 public:
  std::string longestPalindrome(std::string s) {
    int size = s.size();
    if (size == 0) return ""; 
    int left = 0, right = 0;
    int m_left = 0, m_right = 0;
    for (int i = 0; i < size; i++) {
      left = i, right = i;
      findpol(left, right, size, s);
      // std::cout<<left << " " << right << std::endl;
      if ((right - left) > (m_right - m_left)) {
        m_left = left;
        m_right = right;
      }
      left = i, right = i+1;
      findpol(left, right, size, s);
      if ((right - left) > (m_right - m_left)) {
        m_left = left;
        m_right = right-1;
      }
    }
    return s.substr(m_left+1, m_right - m_left);
  }

  void findpol(int& left, int& right, int const& size, std::string s) {
    while (left >= 0 && right < size && s[left] == s[right]) {
      --left;
      ++right;
    }
    std::cout<< std::endl;
    return;
  }
};

int main(void) {
  Solution sol;
  std::string s = "ksaoidweloololllsaodqpwloooooolkfjd";

  std::string ans = sol.longestPalindrome(s);
  std::cout << ans << std::endl;
}