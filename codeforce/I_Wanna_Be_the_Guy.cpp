#include <iostream>
#include <vector>

using namespace std;

vector<int> get_levels() {
    int n;
    cin >> n;

    vector<int> levels(n);
    for (int i = 0; i < n; i++) {
        cin >> levels[i];
    }

    return levels;
}

int main(void) {
  int all_lvls;
  char el = ' ';
  cin >> all_lvls;

  int x = 0,y = 0;

  vector<int> x_levels = get_levels();
  vector<int> y_levels = get_levels();

  for (int i = 1; i <= all_lvls; i++) {

    for (x = 0; x < x_levels.size(); x++) {
      if (i == x_levels[x]) {
        el = 'x';
        break;
      }
    }

    for (y = 0; y < y_levels.size(); y++) {
      if (i == y_levels[y]) {
        el = 'y';
        break;
      }
    }

    if (el != 'x' && el != 'y') {
      cout << "Oh, my keyboard!" << endl;
      return 0;
    }

    el = ' ';
  }

  cout << "I become the guy." << endl;

  return 0;
}