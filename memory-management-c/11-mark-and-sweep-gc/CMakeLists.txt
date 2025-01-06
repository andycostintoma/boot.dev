cmake_minimum_required(VERSION 3.30)
project(11_mark_and_sweep_gc C)

set(CMAKE_C_STANDARD 11)

# CMakeLists.txt
add_executable(11_mark_and_sweep_gc sneknew.c snekobject.c stack.c vm.c)