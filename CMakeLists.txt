cmake_minimum_required(VERSION 3.23)
project(CppVulkanStudy VERSION 1.0.0)

set(CMAKE_CXX_STANDARD 14)

add_executable(${PROJECT_NAME} src/main.cpp)


add_subdirectory(external/glfw-3.3.8)

target_link_libraries(${PROJECT_NAME} glfw)