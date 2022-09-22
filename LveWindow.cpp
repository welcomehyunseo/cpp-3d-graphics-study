//
// Created by welco on 9/23/2022.
//

#include "LveWindow.hpp"

#include <utility>

namespace lve {
    void LveWindow::initWindow() {
        glfwInit();
        glfwWindowHint(GLFW_CLIENT_API, GLFW_NO_API);
        glfwWindowHint(GLFW_RESIZABLE, GLFW_FALSE);

        window = glfwCreateWindow(width, height, windowName.c_str(), nullptr, nullptr);
    }

    LveWindow::LveWindow(int w, int h, std::string name): width{w}, height{h}, windowName{std::move(name)} {
        initWindow();
    }

    LveWindow::~LveWindow() {
        glfwDestroyWindow(window);
        glfwTerminate();
    }
} // lve