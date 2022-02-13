#pragma once

#include <vector>
#include <iostream>

template <typename T>
void cout_vec(std::vector<T> vec) {
  for (auto& el : vec) {
    std::cout << el << " ";
  }
  std::cout << std::endl;
};

template <typename T>
void cout_vec(std::vector<std::vector<T>> vecs) {
  for (auto& vec : vecs) {
    for (auto& el : vec) {
      std::cout << el << " ";
    }
    std::cout << std::endl;
  }
};

